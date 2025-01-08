package vpc

import (
	"context"

	avngen "github.com/aiven/go-client-codegen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func DatasourceAzureOrgVPCPeeringConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: common.WithGenClientDiag(datasourceAzureOrgVPCPeeringConnectionRead),
		Description: "Gets information about about an Azure VPC peering connection.",
		Schema: schemautil.ResourceSchemaAsDatasourceSchema(aivenAzureOrgVPCPeeringConnectionSchema,
			"vpc_id", "azure_subscription_id", "peer_resource_group", "vnet_name"),
	}
}

func datasourceAzureOrgVPCPeeringConnectionRead(ctx context.Context, d *schema.ResourceData, client avngen.Client) diag.Diagnostics {
	var (
		subID = d.Get("azure_subscription_id").(string)
		vnet  = d.Get("vnet_name").(string)
		rg    = d.Get("peer_resource_group").(string)
	)

	orgID, vpcID, err := schemautil.SplitResourceID2(d.Get("vpc_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(schemautil.BuildResourceID(orgID, vpcID, subID, vnet, rg))

	return resourceAzureVPCPeeringConnectionRead(ctx, d, client)
}
