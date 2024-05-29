// Code generated by user config generator. DO NOT EDIT.

package serviceintegrationendpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func datadogUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Datadog user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"datadog_api_key": {
				Description: "Datadog API key.",
				Required:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"datadog_tags": {
				Description: "Custom tags provided by user",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"comment": {
						Description: "Optional tag explanation.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"tag": {
						Description: "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix `aiven-` are reserved for Aiven.",
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
				Description: "Number of separate instances to fetch kafka consumer statistics with.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"kafka_consumer_stats_timeout": {
				Description: "Number of seconds that datadog will wait to get consumer statistics from brokers.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"max_partition_contexts": {
				Description: "Maximum number of partition contexts to send.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"site": {
				Description:  "Enum: `datadoghq.com`, `datadoghq.eu`, `us3.datadoghq.com`, `us5.datadoghq.com`, `ddog-gov.com`, `ap1.datadoghq.com`. Datadog intake site. Defaults to datadoghq.com.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"datadoghq.com", "datadoghq.eu", "us3.datadoghq.com", "us5.datadoghq.com", "ddog-gov.com", "ap1.datadoghq.com"}, false),
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
