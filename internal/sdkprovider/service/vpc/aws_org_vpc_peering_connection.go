package vpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	avngen "github.com/aiven/go-client-codegen"
	"github.com/aiven/go-client-codegen/handler/organizationvpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/plugin/util"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

var aivenAWSOrgVPCPeeringConnectionSchema = map[string]*schema.Schema{
	"vpc_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
		ValidateFunc: func(i interface{}, s string) ([]string, []error) {
			var (
				errs  []error
				warns []string
			)

			v, ok := i.(string)
			if !ok {
				errs = append(errs, fmt.Errorf("expected type of %s to be string", s))

				return warns, errs
			}

			parts := strings.Split(v, "/")
			if len(parts) != 2 {
				errs = append(errs, fmt.Errorf("VPC ID must be in format 'organization_id/organization_vpc_id'"))
				return warns, errs
			}

			if parts[0] == "" || parts[1] == "" {
				errs = append(errs, fmt.Errorf("both organization_id and organization_vpc_id must be non-empty"))
				return warns, errs
			}

			return warns, errs
		},
		Description: userconfig.Desc("The ID of the Aiven Organization VPC.").ForceNew().Build(),
	},
	"aws_account_id": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("AWS account ID.").ForceNew().Build(),
	},
	"aws_vpc_id": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("AWS VPC ID.").ForceNew().Build(),
	},
	"aws_vpc_region": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("The AWS region of the peered VPC. For example, `eu-central-1`.").Build(),
	},
	"organization_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Identifier of the organization.",
	},
	"organization_vpc_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Identifier of the organization VPC.",
	},
	"peering_connection_id": {
		Computed:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("The ID of the peering connection.").Build(),
	},
	"aws_vpc_peering_connection_id": {
		Computed:    true,
		Type:        schema.TypeString,
		Description: "The ID of the AWS VPC peering connection.",
	},
	"state": {
		Computed:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("State of the peering connection.").PossibleValuesString(organizationvpc.VpcPeeringConnectionStateTypeChoices()...).Build(),
	},
}

func ResourceAWSOrgVPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages an AWS VPC peering connection with an Aiven Organization VPC.",
		CreateContext: common.WithGenClientDiag(resourceAWSOrgVPCPeeringConnectionCreate),
		ReadContext:   common.WithGenClientDiag(resourceAWSOrgVPCPeeringConnectionRead),
		DeleteContext: common.WithGenClientDiag(resourceAWSOrgVPCPeeringConnectionDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema: aivenAWSOrgVPCPeeringConnectionSchema,
	}
}

func resourceAWSOrgVPCPeeringConnectionCreate(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	var (
		vpcID        = d.Get("vpc_id").(string)
		awsAccountID = d.Get("aws_account_id").(string)
		awsVPCId     = d.Get("aws_vpc_id").(string)
		awsRegion    = d.Get("aws_vpc_region").(string)
	)

	orgID, orgVpcID, err := schemautil.SplitResourceID2(vpcID)
	if err != nil {
		return diag.FromErr(err)
	}

	var req = organizationvpc.OrganizationVpcPeeringConnectionCreateIn{
		PeerRegion:       util.ToPtr(awsRegion),
		PeerVpc:          awsVPCId,
		PeerCloudAccount: awsAccountID,
	}

	pc, err := client.OrganizationVpcPeeringConnectionCreate(
		ctx,
		orgID,
		orgVpcID,
		&req,
	)
	if err != nil || pc.PeeringConnectionId == nil {
		return diag.Errorf("Error creating VPC peering connection: %s", err)
	}

	// wait for the VPC peering connection to be approved
	stateChangeConf := &retry.StateChangeConf{
		Target:  []string{""}, // empty target means we don't care about the target state
		Pending: []string{string(organizationvpc.VpcPeeringConnectionStateTypeApproved)},
		Refresh: func() (any, string, error) {
			resp, err := client.OrganizationVpcGet(ctx, orgID, orgVpcID)
			if err != nil {
				return nil, "", err
			}

			for _, pCon := range resp.PeeringConnections {
				if pCon.PeeringConnectionId != nil && *pCon.PeeringConnectionId == *pc.PeeringConnectionId {
					if pCon.State == organizationvpc.VpcPeeringConnectionStateTypeApproved {

						return pCon, string(pCon.State), nil
					}

					return pCon, "", nil
				}
			}

			return nil, "", fmt.Errorf("VPC peering connection not found")
		},
		Delay:      1 * time.Second,
		Timeout:    d.Timeout(schema.TimeoutCreate),
		MinTimeout: common.DefaultStateChangeMinTimeout,
	}

	res, err := stateChangeConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("Error creating VPC peering connection: %s", err)
	}

	pCon, ok := res.(organizationvpc.OrganizationVpcGetPeeringConnectionOut)
	if !ok {
		// this should never happen
		return diag.Errorf("Error creating VPC peering connection: %s", err)
	}

	diags := getDiagnosticsFromState(newOrganizationVPCPeeringState(&pCon))

	d.SetId(schemautil.BuildResourceID(orgID, orgVpcID, awsAccountID, awsVPCId, awsRegion))

	// in case of an error delete VPC peering connection
	if diags.HasError() {
		return append(diags, resourceAWSOrgVPCPeeringConnectionDelete(ctx, d, client)...)
	}

	return append(diags, resourceAWSOrgVPCPeeringConnectionRead(ctx, d, client)...)
}

func resourceAWSOrgVPCPeeringConnectionRead(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	orgID, orgVpcID, awsAccountID, awsVpcID, awsRegion, err := schemautil.SplitResourceID5(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	vpc, err := client.OrganizationVpcGet(ctx, orgID, orgVpcID)
	if err != nil {
		if avngen.IsNotFound(err) {
			return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
		}

		return diag.Errorf("Error finding VPC: %s", err)
	}

	pc := lookupAWSPeeringConnection(vpc, awsAccountID, awsVpcID, awsRegion)
	if pc == nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
	}

	if err = d.Set("organization_id", vpc.OrganizationId); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("organization_vpc_id", vpc.OrganizationVpcId); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("peering_connection_id", *pc.PeeringConnectionId); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("aws_account_id", pc.PeerCloudAccount); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("aws_vpc_id", pc.PeerVpc); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("aws_vpc_region", *pc.PeerRegion); err != nil {
		return diag.FromErr(err)
	}
	//TODO: fix the API response schema, at this moment it's not complete
	if err = d.Set("aws_vpc_peering_connection_id", pc.StateInfo.AwsVpcPeeringConnectionId); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("state", string(pc.State)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceAWSOrgVPCPeeringConnectionDelete(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	orgID, orgVpcID, awsAccountID, awsVpcID, awsRegion, err := schemautil.SplitResourceID5(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	vpc, err := client.OrganizationVpcGet(ctx, orgID, orgVpcID)
	if common.IsCritical(err) {
		// if the VPC is not found, consider all peering connections as deleted
		return diag.Errorf("Error finding VPC: %s", err)
	}

	pc := lookupAWSPeeringConnection(vpc, awsAccountID, awsVpcID, awsRegion)
	if pc == nil {
		// if the peering connection is not found, consider it deleted
		return nil
	}

	_, err = client.OrganizationVpcPeeringConnectionDeleteById(
		ctx,
		vpc.OrganizationId,
		vpc.OrganizationVpcId,
		*pc.PeeringConnectionId,
	)
	if common.IsCritical(err) {
		return diag.Errorf("Error deleting VPC peering connection: %s", err)
	}

	stateChangeConf := &retry.StateChangeConf{
		Target: []string{string(organizationvpc.VpcPeeringConnectionStateTypeDeleted)},
		Refresh: func() (interface{}, string, error) {
			resp, err := client.OrganizationVpcGet(ctx, orgID, orgVpcID)
			if err != nil {
				if avngen.IsNotFound(err) {
					// If the resource is not found, consider it deleted
					// Return empty struct to signal to the state change function that the resource is deleted
					return struct{}{}, string(organizationvpc.VpcPeeringConnectionStateTypeDeleted), nil
				}

				return nil, "", err
			}

			pc = lookupAWSPeeringConnection(resp, awsAccountID, awsVpcID, awsRegion)
			if pc == nil {
				// if the peering connection is not found, consider it deleted
				return struct{}{}, string(organizationvpc.VpcPeeringConnectionStateTypeDeleted), nil
			}

			return pc, string(pc.State), nil
		},
		Delay:      1 * time.Second,
		Timeout:    d.Timeout(schema.TimeoutDelete),
		MinTimeout: common.DefaultStateChangeMinTimeout,
	}

	if _, err = stateChangeConf.WaitForStateContext(ctx); err != nil && !avngen.IsNotFound(err) {
		return diag.Errorf("Error waiting for AWS Aiven VPC Peering Connection to be DELETED: %s", err)
	}

	return nil
}

func lookupAWSPeeringConnection(
	vpc *organizationvpc.OrganizationVpcGetOut,
	awsAccountID, awsVpcID, awsRegion string,
) *organizationvpc.OrganizationVpcGetPeeringConnectionOut {
	for _, pCon := range vpc.PeeringConnections {
		if pCon.PeerCloudAccount == awsAccountID &&
			pCon.PeerVpc == awsVpcID &&
			pCon.PeerRegion != nil &&
			*pCon.PeerRegion == awsRegion &&
			pCon.PeeringConnectionId != nil {

			return &pCon
		}
	}

	return nil
}
