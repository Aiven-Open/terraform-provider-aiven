package thanos

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func DatasourceThanos() *schema.Resource {
	return &schema.Resource{
		ReadContext: schemautil.DatasourceServiceRead,
		Description: "Gets information about an Aiven for Thanos® service.",
		Schema:      schemautil.ResourceSchemaAsDatasourceSchema(thanosSchema(), "project", "service_name"),
	}
}
