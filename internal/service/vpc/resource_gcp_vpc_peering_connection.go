package vpc

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aiven/terraform-provider-aiven/internal/meta"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var aivenGCPVPCPeeringConnectionSchema = map[string]*schema.Schema{
	"vpc_id": {
		ForceNew:     true,
		Required:     true,
		Type:         schema.TypeString,
		Description:  schemautil.Complex("The VPC the peering connection belongs to.").ForceNew().Build(),
		ValidateFunc: validateVPCID,
	},
	"gcp_project_id": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: schemautil.Complex("GCP project ID.").ForceNew().Build(),
	},
	"peer_vpc": {
		ForceNew:    true,
		Required:    true,
		Type:        schema.TypeString,
		Description: schemautil.Complex("GCP VPC network name.").ForceNew().Build(),
	},
	"state": {
		Computed:    true,
		Type:        schema.TypeString,
		Description: "State of the peering connection",
	},
	"state_info": {
		Computed:    true,
		Type:        schema.TypeMap,
		Description: "State-specific help or error information",
	},
}

func ResourceGCPVPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		Description:   "The GCP VPC Peering Connection resource allows the creation and management of Aiven GCP VPC Peering Connections.",
		CreateContext: resourceGCPVPCPeeringConnectionCreate,
		ReadContext:   resourceGCPVPCPeeringConnectionRead,
		DeleteContext: resourceGCPVPCPeeringConnectionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGCPVPCPeeringConnectionImport,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},

		Schema: aivenGCPVPCPeeringConnectionSchema,
	}
}

func resourceGCPVPCPeeringConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		pc  *aiven.VPCPeeringConnection
		err error
	)

	client := m.(*meta.Meta).Client
	projectName, vpcID := schemautil.SplitResourceID2(d.Get("vpc_id").(string))
	gcpProjectId := d.Get("gcp_project_id").(string)
	peerVPC := d.Get("peer_vpc").(string)

	if _, err = client.VPCPeeringConnections.Create(
		projectName,
		vpcID,
		aiven.CreateVPCPeeringConnectionRequest{
			PeerCloudAccount: gcpProjectId,
			PeerVPC:          peerVPC,
		},
	); err != nil {
		return diag.Errorf("Error waiting for VPC peering connection creation: %s", err)
	}

	stateChangeConf := &resource.StateChangeConf{
		Pending: []string{"APPROVED"},
		Target: []string{
			"ACTIVE",
			"REJECTED_BY_PEER",
			"PENDING_PEER",
			"INVALID_SPECIFICATION",
			"DELETING",
			"DELETED",
			"DELETED_BY_PEER",
		},
		Refresh: func() (interface{}, string, error) {
			pc, err := client.VPCPeeringConnections.GetVPCPeering(
				projectName,
				vpcID,
				gcpProjectId,
				peerVPC,
				nil,
			)
			if err != nil {
				return nil, "", err
			}
			return pc, pc.State, nil
		},
		Delay:      10 * time.Second,
		Timeout:    d.Timeout(schema.TimeoutCreate),
		MinTimeout: 2 * time.Second,
	}

	res, err := stateChangeConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.Errorf("Error creating VPC peering connection: %s", err)
	}

	pc = res.(*aiven.VPCPeeringConnection)
	diags := getDiagnosticsFromState(pc)

	d.SetId(schemautil.BuildResourceID(projectName, vpcID, pc.PeerCloudAccount, pc.PeerVPC))

	// in case of an error delete VPC peering connection
	if diags.HasError() {
		return append(diags, resourceGCPVPCPeeringConnectionDelete(ctx, d, m)...)
	}

	return append(diags, resourceGCPVPCPeeringConnectionRead(ctx, d, m)...)
}

func resourceGCPVPCPeeringConnectionRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var pc *aiven.VPCPeeringConnection
	client := m.(*meta.Meta).Client

	projectName, vpcID, peerCloudAccount, peerVPC, peerRegion := parsePeeringVPCId(d.Id())
	pc, err := client.VPCPeeringConnections.GetVPCPeering(
		projectName, vpcID, peerCloudAccount, peerVPC, peerRegion)
	if err != nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d, m))
	}

	return copyGCPVPCPeeringConnectionPropertiesFromAPIResponseToTerraform(d, pc, projectName, vpcID)
}

func resourceGCPVPCPeeringConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*meta.Meta).Client

	projectName, vpcID, peerCloudAccount, peerVPC, peerRegion := parsePeeringVPCId(d.Id())

	err := client.VPCPeeringConnections.DeleteVPCPeering(
		projectName,
		vpcID,
		peerCloudAccount,
		peerVPC,
		peerRegion,
	)
	if err != nil && !aiven.IsNotFound(err) {
		return diag.Errorf("Error deleting GCP VPC peering connection: %s", err)
	}

	stateChangeConf := &resource.StateChangeConf{
		Pending: []string{
			"ACTIVE",
			"APPROVED",
			"APPROVED_PEER_REQUESTED",
			"DELETING",
			"INVALID_SPECIFICATION",
			"PENDING_PEER",
			"REJECTED_BY_PEER",
			"DELETED_BY_PEER",
		},
		Target: []string{
			"DELETED",
		},
		Refresh: func() (interface{}, string, error) {
			pc, err := client.VPCPeeringConnections.GetVPCPeering(
				projectName,
				vpcID,
				peerCloudAccount,
				peerVPC,
				peerRegion,
			)
			if err != nil {
				return nil, "", err
			}
			return pc, pc.State, nil
		},
		Delay:      10 * time.Second,
		Timeout:    d.Timeout(schema.TimeoutDelete),
		MinTimeout: 2 * time.Second,
	}
	if _, err := stateChangeConf.WaitForStateContext(ctx); err != nil && !aiven.IsNotFound(err) {
		return diag.Errorf("Error waiting for GCP Aiven VPC Peering Connection to be DELETED: %s", err)
	}
	return nil
}

func resourceGCPVPCPeeringConnectionImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	m.(*meta.Meta).Import = true

	if len(strings.Split(d.Id(), "/")) != 4 {
		return nil, fmt.Errorf("invalid identifier %v, expected <project_name>/<vpc_id>/<gcp_project_id>/<peer_vpc>", d.Id())
	}

	dig := resourceGCPVPCPeeringConnectionRead(ctx, d, m)
	if dig.HasError() {
		return nil, errors.New("cannot get GCP VPC peering connection")
	}

	return []*schema.ResourceData{d}, nil
}

func copyGCPVPCPeeringConnectionPropertiesFromAPIResponseToTerraform(
	d *schema.ResourceData,
	peeringConnection *aiven.VPCPeeringConnection,
	project string,
	vpcID string,
) diag.Diagnostics {
	if err := d.Set("vpc_id", schemautil.BuildResourceID(project, vpcID)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("gcp_project_id", peeringConnection.PeerCloudAccount); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("peer_vpc", peeringConnection.PeerVPC); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", peeringConnection.State); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state_info", ConvertStateInfoToMap(peeringConnection.StateInfo)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
