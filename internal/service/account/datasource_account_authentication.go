package account

import (
	"context"

	"github.com/aiven/terraform-provider-aiven/internal/meta"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceAccountAuthentication() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceAccountAuthenticationRead,
		Description: "The Account Authentication data source provides information about the existing Aiven Account Authentication.",
		Schema: schemautil.ResourceSchemaAsDatasourceSchema(aivenAccountAuthenticationSchema,
			"account_id", "name"),
	}
}

func datasourceAccountAuthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*meta.Meta).Client

	name := d.Get("name").(string)
	accountId := d.Get("account_id").(string)

	r, err := client.AccountAuthentications.List(accountId)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, a := range r.AuthenticationMethods {
		if a.Name == name {
			d.SetId(schemautil.BuildResourceID(a.AccountId, a.Id))
			return resourceAccountAuthenticationRead(ctx, d, m)
		}
	}

	return diag.Errorf("account authentication %s not found", name)
}
