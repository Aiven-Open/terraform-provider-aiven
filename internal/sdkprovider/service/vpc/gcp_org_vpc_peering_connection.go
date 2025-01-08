package vpc

import (
	"context"
	"errors"

	avngen "github.com/aiven/go-client-codegen"
	"github.com/aiven/go-client-codegen/handler/organizationvpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

var aivenGCPOrgVPCPeeringConnectionSchema = map[string]*schema.Schema{
	"vpc_id": {
		ForceNew:     true,
		Required:     true,
		Type:         schema.TypeString,
		Description:  userconfig.Desc("The VPC the peering connection belongs to.").ForceNew().Build(),
		ValidateFunc: validateVPCID,
	},
	"gcp_project_id": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("Google Cloud project ID.").ForceNew().Build(),
	},
	"peer_vpc": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: userconfig.Desc("Google Cloud VPC network name.").ForceNew().Build(),
	},
	"state": {
		Computed:    true,
		Type:        schema.TypeString,
		Description: "State of the peering connection.",
	},
}

func ResourceGCPOrgVPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages a Google Cloud VPC peering connection.",
		CreateContext: common.WithGenClientDiag(resourceGCPOrgVPCPeeringConnectionCreate),
		ReadContext:   common.WithGenClientDiag(resourceGCPOrgVPCPeeringConnectionRead),
		DeleteContext: common.WithGenClientDiag(resourceGCPOrgVPCPeeringConnectionDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema: aivenGCPOrgVPCPeeringConnectionSchema,
	}
}

func resourceGCPOrgVPCPeeringConnectionCreate(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	orgID, vpcID, err := schemautil.SplitResourceID2(d.Get("vpc_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	var (
		gcpProjectID = d.Get("gcp_project_id").(string)
		peerVPC      = d.Get("peer_vpc").(string)

		req = organizationvpc.OrganizationVpcPeeringConnectionCreateIn{
			PeerCloudAccount: gcpProjectID,
			PeerVpc:          peerVPC,
		}
	)

	pCon, err := createPeeringConnection(ctx, orgID, vpcID, client, d, req)
	if err != nil {
		return diag.Errorf("Error creating VPC peering connection: %s", err)
	}

	diags := getDiagnosticsFromState(newOrganizationVPCPeeringState(pCon))

	d.SetId(schemautil.BuildResourceID(orgID, vpcID, pCon.PeerCloudAccount, pCon.PeerVpc))

	// in case of an error delete VPC peering connection
	if diags.HasError() {
		deleteDiags := resourceAzureOrgVPCPeeringConnectionDelete(ctx, d, client)
		d.SetId("") // Clear the ID after delete

		return append(diags, deleteDiags...)
	}

	return append(diags, resourceGCPOrgVPCPeeringConnectionRead(ctx, d, client)...)
}

func resourceGCPOrgVPCPeeringConnectionRead(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	orgID, vpcID, cloudAcc, peerVPC, err := schemautil.SplitResourceID4(d.Id())
	if err != nil {
		return diag.Errorf("error parsing GCP peering VPC ID: %s", err)
	}

	vpc, err := client.OrganizationVpcGet(ctx, orgID, vpcID)
	if err != nil {
		if avngen.IsNotFound(err) {
			return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
		}

		return diag.Errorf("Error finding VPC: %s", err)
	}

	pCon := lookupGCPPeeringConnection(vpc, cloudAcc, peerVPC)
	if pCon == nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(errors.New("VPC peering connection not found"), d))
	}

	if err = d.Set("vpc_id", schemautil.BuildResourceID(orgID, vpcID)); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("gcp_project_id", pCon.PeerCloudAccount); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("peer_vpc", pCon.PeerVpc); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("state", string(pCon.State)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGCPOrgVPCPeeringConnectionDelete(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	orgID, vpcID, cloudAcc, peerVPC, err := schemautil.SplitResourceID4(d.Id())
	if err != nil {
		return diag.Errorf("error parsing GCP peering VPC ID: %s", err)
	}

	vpc, err := client.OrganizationVpcGet(ctx, orgID, vpcID)
	if err != nil {
		if avngen.IsNotFound(err) {
			return nil // consider already deleted
		}

		return diag.Errorf("Error finding VPC: %s", err)
	}

	if err = deletePeeringConnection(
		ctx,
		orgID,
		vpcID,
		client,
		d,
		lookupGCPPeeringConnection(vpc, cloudAcc, peerVPC),
	); err != nil {
		return diag.Errorf("Error deleting GCP Aiven VPC Peering Connection: %s", err)
	}

	return nil
}

func lookupGCPPeeringConnection(
	vpc *organizationvpc.OrganizationVpcGetOut,
	cloudAcc, peerVPC string,
) *organizationvpc.OrganizationVpcGetPeeringConnectionOut {
	var pCon *organizationvpc.OrganizationVpcGetPeeringConnectionOut
	for _, p := range vpc.PeeringConnections {
		if p.PeerCloudAccount == cloudAcc && p.PeerVpc == peerVPC && p.PeeringConnectionId != nil {
			pCon = &p
			break
		}
	}

	return pCon
}
