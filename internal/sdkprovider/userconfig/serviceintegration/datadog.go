// Code generated by user config generator. DO NOT EDIT.

package serviceintegration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func datadogUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Datadog user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"datadog_dbm_enabled": {
				Description: "Enable Datadog Database Monitoring.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"datadog_pgbouncer_enabled": {
				Description: "Enable Datadog PgBouncer Metric Tracking.",
				Optional:    true,
				Type:        schema.TypeBool,
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
			"exclude_consumer_groups": {
				Description: "List of custom metrics.",
				Elem: &schema.Schema{
					Description: "Consumer groups to exclude.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"exclude_topics": {
				Description: "List of topics to exclude.",
				Elem: &schema.Schema{
					Description: "Topics to exclude.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"include_consumer_groups": {
				Description: "List of custom metrics.",
				Elem: &schema.Schema{
					Description: "Consumer groups to include.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"include_topics": {
				Description: "List of topics to include.",
				Elem: &schema.Schema{
					Description: "Topics to include.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"kafka_custom_metrics": {
				Description: "List of custom metrics.",
				Elem: &schema.Schema{
					Description: "Metric name.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"max_jmx_metrics": {
				Description: "Maximum number of JMX metrics to send.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"opensearch": {
				Description: "Datadog Opensearch Options",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"cluster_stats_enabled": {
						Description: "Enable Datadog Opensearch Cluster Monitoring.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"index_stats_enabled": {
						Description: "Enable Datadog Opensearch Index Monitoring.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"pending_task_stats_enabled": {
						Description: "Enable Datadog Opensearch Pending Task Monitoring.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"pshard_stats_enabled": {
						Description: "Enable Datadog Opensearch Primary Shard Monitoring.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"redis": {
				Description: "Datadog Redis Options",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"command_stats_enabled": {
					Description: "Enable command_stats option in the agent's configuration. The default value is `false`.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
