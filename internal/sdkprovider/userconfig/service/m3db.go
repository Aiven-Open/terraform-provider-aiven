// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func m3dbUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "M3db user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"additional_backup_regions": {
				Description: "Additional Cloud Regions for Backup Replication.",
				Elem: &schema.Schema{
					Description: "Target cloud.",
					Type:        schema.TypeString,
				},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"custom_domain": {
				Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"ip_filter_object": {
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"description": {
						Description: "Description for IP filter list entry.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"network": {
						Description: "CIDR address block.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeList,
			},
			"ip_filter_string": {
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"limits": {
				Description: "M3 limits",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"max_recently_queried_series_blocks": {
						Description: "The maximum number of blocks that can be read in a given lookback period.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_recently_queried_series_disk_bytes_read": {
						Description: "The maximum number of disk bytes that can be read in a given lookback period.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_recently_queried_series_lookback": {
						Description: "The lookback period for 'max_recently_queried_series_blocks' and 'max_recently_queried_series_disk_bytes_read'.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"query_docs": {
						Description: "The maximum number of docs fetched in single query.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"query_require_exhaustive": {
						Description: "When query limits are exceeded, whether to return error or return partial results.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"query_series": {
						Description: "The maximum number of series fetched in single query.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"m3": {
				Description: "M3 specific configuration options",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"tag_options": {
					Description: "M3 Tag Options",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"allow_tag_name_duplicates": {
							Description: "Allows for duplicate tags to appear on series (not allowed by default).",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"allow_tag_value_empty": {
							Description: "Allows for empty tags to appear on series (not allowed by default).",
							Optional:    true,
							Type:        schema.TypeBool,
						},
					}},
					MaxItems: 1,
					Optional: true,
					Type:     schema.TypeList,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"m3_version": {
				Description:  "M3 major version (deprecated, use m3db_version).",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1.1", "1.2", "1.5"}, false),
			},
			"m3coordinator_enable_graphite_carbon_ingest": {
				Description: "Enables access to Graphite Carbon plaintext metrics ingestion. It can be enabled only for services inside VPCs. The metrics are written to aggregated namespaces only.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"m3db_version": {
				Description:  "M3 major version (the minimum compatible version).",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"1.1", "1.2", "1.5"}, false),
			},
			"namespaces": {
				Description: "List of M3 namespaces",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"name": {
						Description: "The name of the namespace.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"options": {
						Description: "Namespace options",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"retention_options": {
								Description: "Retention options",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"block_data_expiry_duration": {
										Description: "Controls how long we wait before expiring stale data.",
										Optional:    true,
										Type:        schema.TypeString,
									},
									"blocksize_duration": {
										Description: "Controls how long to keep a block in memory before flushing to a fileset on disk.",
										Optional:    true,
										Type:        schema.TypeString,
									},
									"buffer_future_duration": {
										Description: "Controls how far into the future writes to the namespace will be accepted.",
										Optional:    true,
										Type:        schema.TypeString,
									},
									"buffer_past_duration": {
										Description: "Controls how far into the past writes to the namespace will be accepted.",
										Optional:    true,
										Type:        schema.TypeString,
									},
									"retention_period_duration": {
										Description: "Controls the duration of time that M3DB will retain data for the namespace.",
										Optional:    true,
										Type:        schema.TypeString,
									},
								}},
								MaxItems: 1,
								Required: true,
								Type:     schema.TypeList,
							},
							"snapshot_enabled": {
								Description: "Controls whether M3DB will create snapshot files for this namespace.",
								Optional:    true,
								Type:        schema.TypeBool,
							},
							"writes_to_commitlog": {
								Description: "Controls whether M3DB will include writes to this namespace in the commitlog.",
								Optional:    true,
								Type:        schema.TypeBool,
							},
						}},
						MaxItems: 1,
						Optional: true,
						Type:     schema.TypeList,
					},
					"resolution": {
						Description: "The resolution for an aggregated namespace.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"type": {
						Description:  "The type of aggregation (aggregated/unaggregated).",
						Required:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"aggregated", "unaggregated"}, false),
					},
				}},
				MaxItems: 2147483647,
				Optional: true,
				Type:     schema.TypeList,
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"m3coordinator": {
					Description: "Allow clients to connect to m3coordinator with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"project_to_fork_from": {
				Description: "Name of another project to fork a service from. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"m3coordinator": {
					Description: "Allow clients to connect to m3coordinator from the public internet for service nodes that are in a project VPC or another type of private network.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"rules": {
				Description: "M3 rules",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"mapping": {
					Description: "List of M3 mapping rules",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"aggregations": {
							Description: "List of aggregations to be applied.",
							Elem: &schema.Schema{
								Description:  "Aggregation to be applied.",
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"Count", "Last", "Max", "Mean", "Median", "Min", "P10", "P20", "P30", "P40", "P50", "P60", "P70", "P80", "P90", "P95", "P99", "P999", "P9999", "Stdev", "Sum", "SumSq"}, false),
							},
							MaxItems: 10,
							Optional: true,
							Type:     schema.TypeSet,
						},
						"drop": {
							Description: "Only store the derived metric (as specified in the roll-up rules), if any.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"filter": {
							Description: "Matching metric names with wildcards (using __name__:wildcard) or matching tags and their (optionally wildcarded) values. For value, ! can be used at start of value for negation, and multiple filters can be supplied using space as separator.",
							Required:    true,
							Type:        schema.TypeString,
						},
						"name": {
							Description: "The (optional) name of the rule.",
							Optional:    true,
							Type:        schema.TypeString,
						},
						"namespaces": {
							Deprecated:  "Deprecated. Use `namespaces_string` instead.",
							Description: "This rule will be used to store the metrics in the given namespace(s). If a namespace is target of rules, the global default aggregation will be automatically disabled. Note that specifying filters that match no namespaces whatsoever will be returned as an error. Filter the namespace by glob (=wildcards).",
							Elem: &schema.Schema{
								Description: "Filter the namespace by glob (=wildcards).",
								Type:        schema.TypeString,
							},
							MaxItems: 10,
							Optional: true,
							Type:     schema.TypeSet,
						},
						"namespaces_object": {
							Description: "This rule will be used to store the metrics in the given namespace(s). If a namespace is target of rules, the global default aggregation will be automatically disabled. Note that specifying filters that match no namespaces whatsoever will be returned as an error. Filter the namespace by exact match of retention period and resolution",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"resolution": {
									Description: "The resolution for the matching namespace.",
									Required:    true,
									Type:        schema.TypeString,
								},
								"retention": {
									Description: "The retention period of the matching namespace.",
									Optional:    true,
									Type:        schema.TypeString,
								},
							}},
							MaxItems: 10,
							Optional: true,
							Type:     schema.TypeList,
						},
						"namespaces_string": {
							Description: "This rule will be used to store the metrics in the given namespace(s). If a namespace is target of rules, the global default aggregation will be automatically disabled. Note that specifying filters that match no namespaces whatsoever will be returned as an error. Filter the namespace by glob (=wildcards).",
							Elem: &schema.Schema{
								Description: "Filter the namespace by glob (=wildcards).",
								Type:        schema.TypeString,
							},
							MaxItems: 10,
							Optional: true,
							Type:     schema.TypeSet,
						},
						"tags": {
							Description: "List of tags to be appended to matching metrics",
							Elem: &schema.Resource{Schema: map[string]*schema.Schema{
								"name": {
									Description: "Name of the tag.",
									Required:    true,
									Type:        schema.TypeString,
								},
								"value": {
									Description: "Value of the tag.",
									Required:    true,
									Type:        schema.TypeString,
								},
							}},
							MaxItems: 10,
							Optional: true,
							Type:     schema.TypeList,
						},
					}},
					MaxItems: 10,
					Optional: true,
					Type:     schema.TypeList,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"service_to_fork_from": {
				Description: "Name of another service to fork from. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
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
