package vpc

import (
	"context"

	avngen "github.com/aiven/go-client-codegen"
	"github.com/aiven/go-client-codegen/handler/organizationvpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

func DataSourceOrganizationVPC() *schema.Resource {
	return &schema.Resource{
		Description: "Gets information about an existing VPC in an Aiven organization.",
		ReadContext: common.WithGenClient(datasourceOrganizationVPCRead),
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier of the organization.",
			},
			"organization_vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Aiven Organization VPC.",
			},
			"cloud_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: userconfig.Desc("The cloud provider and region where the service is hosted in the format `CLOUD_PROVIDER-REGION_NAME`. For example, `google-europe-west1` or `aws-us-east-2`.").ForceNew().Build(),
			},
			"network_cidr": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network address range used by the VPC.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: userconfig.Desc("State of the VPC.").PossibleValuesString(organizationvpc.VpcStateTypeChoices()...).Build(),
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time of creation of the VPC.",
			},
		},
	}
}

func datasourceOrganizationVPCRead(ctx context.Context, d *schema.ResourceData, client avngen.Client) error {
	var (
		orgID    = d.Get("organization_id").(string)
		orgVpcID = d.Get("organization_vpc_id").(string)
	)

	d.SetId(schemautil.BuildResourceID(orgID, orgVpcID))

	return resourceOrganizationVPCRead(ctx, d, client)
}
