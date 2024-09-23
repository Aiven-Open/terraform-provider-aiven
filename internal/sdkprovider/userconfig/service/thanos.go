// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func thanosUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Thanos user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"compactor": {
				Description: "ThanosCompactor",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"retention_days": {
					Description: "Retention time for data in days for each resolution (5m, 1h, raw).",
					Optional:    true,
					Type:        schema.TypeInt,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"env": {
				Description: "Environmental variables.",
				Optional:    true,
				Type:        schema.TypeMap,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field. Example: `10.20.0.0/16`.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"ip_filter_object": {
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"description": {
						Description: "Description for IP filter list entry. Example: `Production service IP range`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"network": {
						Description: "CIDR address block. Example: `10.20.0.0/16`.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"ip_filter_string": {
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field. Example: `10.20.0.0/16`.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"object_storage_usage_alert_threshold_gb": {
				Description: "After exceeding the limit a service alert is going to be raised (0 means not set).",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"compactor": {
						Description: "Allow clients to connect to compactor from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"query": {
						Description: "Allow clients to connect to query from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"query_frontend": {
						Description: "Allow clients to connect to query_frontend from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"receiver_ingesting": {
						Description: "Allow clients to connect to receiver_ingesting from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"receiver_routing": {
						Description: "Allow clients to connect to receiver_routing from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"store": {
						Description: "Allow clients to connect to store from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"query": {
				Description: "ThanosQuery",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"query_default_evaluation_interval": {
						Description: "Set the default evaluation interval for subqueries. Default: `1m`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"query_lookback_delta": {
						Description: "The maximum lookback duration for retrieving metrics during expression evaluations in PromQL. PromQL always evaluates the query for a certain timestamp, and it looks back for the given amount of time to get the latest sample. If it exceeds the maximum lookback delta, it assumes the series is stale and returns none (a gap). The lookback delta should be set to at least 2 times the slowest scrape interval. If unset, it will use the promql default of 5m. Default: `5m`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"query_metadata_default_time_range": {
						Description: "The default metadata time range duration for retrieving labels through Labels and Series API when the range parameters are not specified. The zero value means the range covers the time since the beginning. Default: `0s`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"query_timeout": {
						Description: "Maximum time to process a query by the query node. Default: `2m`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"store_limits_request_samples": {
						Description: "The maximum samples allowed for a single Series request. The Series call fails if this limit is exceeded. Set to 0 for no limit. NOTE: For efficiency, the limit is internally implemented as 'chunks limit' considering each chunk contains a maximum of 120 samples. The default value is 100 * store.limits.request-series. Default: `0`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"store_limits_request_series": {
						Description: "The maximum series allowed for a single Series request. The Series call fails if this limit is exceeded. Set to 0 for no limit. The default value is 1000 * cpu_count. Default: `0`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"query_frontend": {
				Description: "ThanosQueryFrontend",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"query_range_align_range_with_step": {
					Description: "Whether to align the query range boundaries with the step. If enabled, the query range boundaries will be aligned to the step, providing more accurate results for queries with high-resolution data. Default: `true`.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"service_log": {
				Description: "Store logs for the service so that they are available in the HTTP API and console.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"static_ips": {
				Description: "Use static public IP addresses.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
