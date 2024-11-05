// Code generated by user config generator. DO NOT EDIT.

package serviceintegration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func logsUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Logs user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"elasticsearch_index_days_max": {
				Description: "Elasticsearch index retention limit. Default: `3`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"elasticsearch_index_prefix": {
				Description: "Elasticsearch index prefix. Default: `logs`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"selected_log_fields": {
				Description: "The list of logging fields that will be sent to the integration logging service. The MESSAGE and timestamp fields are always sent.",
				Elem: &schema.Schema{
					Description:  "Enum: `HOSTNAME`, `PRIORITY`, `REALTIME_TIMESTAMP`, `SYSTEMD_UNIT`, `service_name`. Log field name.",
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"HOSTNAME", "PRIORITY", "REALTIME_TIMESTAMP", "SYSTEMD_UNIT", "service_name"}, false),
				},
				MaxItems: 5,
				Optional: true,
				Type:     schema.TypeList,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
