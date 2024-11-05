// Code generated by user config generator. DO NOT EDIT.

package serviceintegrationendpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func datadogUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Datadog user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"datadog_api_key": {
				Description: "Datadog API key. Example: `848f30907c15c55d601fe45487cce9b6`.",
				Required:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"datadog_tags": {
				Description: "Custom tags provided by user",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"comment": {
						Description: "Optional tag explanation. Example: `Used to tag primary replica metrics`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"tag": {
						Description: "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix `aiven-` are reserved for Aiven. Example: `replica:primary`.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 32,
				Optional: true,
				Type:     schema.TypeList,
			},
			"disable_consumer_stats": {
				Description: "Disable consumer group metrics.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"kafka_consumer_check_instances": {
				Description: "Number of separate instances to fetch kafka consumer statistics with. Example: `8`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"kafka_consumer_stats_timeout": {
				Description: "Number of seconds that datadog will wait to get consumer statistics from brokers. Example: `60`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"max_partition_contexts": {
				Description: "Maximum number of partition contexts to send. Example: `32000`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"site": {
				Description:  "Enum: `ap1.datadoghq.com`, `datadoghq.com`, `datadoghq.eu`, `ddog-gov.com`, `us3.datadoghq.com`, `us5.datadoghq.com`. Datadog intake site. Defaults to datadoghq.com.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ap1.datadoghq.com", "datadoghq.com", "datadoghq.eu", "ddog-gov.com", "us3.datadoghq.com", "us5.datadoghq.com"}, false),
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
