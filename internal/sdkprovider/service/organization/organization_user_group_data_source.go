package organization

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func DatasourceOrganizationUserGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceOrganizationUserGroupRead,
		Description: "The Organization User Groupe data source provides information about the existing Aiven" +
			" Organization User Group.",
		Schema: schemautil.ResourceSchemaAsDatasourceSchema(
			aivenOrganizationUserGroupSchema, "organization_id", "name",
		),
	}
}

func datasourceOrganizationUserGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	organizationID := d.Get("organization_id").(string)
	name := d.Get("name").(string)

	client := m.(*aiven.Client)
	list, err := client.OrganizationUserGroups.List(organizationID)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, ug := range list.UserGroups {
		if ug.UserGroupName == name {
			d.SetId(schemautil.BuildResourceID(organizationID, ug.UserGroupID))
			return resourceOrganizationUserGroupRead(ctx, d, m)
		}
	}

	return diag.Errorf("organization user group %s not found", name)
}
