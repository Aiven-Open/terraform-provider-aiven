// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func opensearchUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Opensearch user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"additional_backup_regions": {
				Description: "Additional Cloud Regions for Backup Replication.",
				Elem: &schema.Schema{
					Description: "Target cloud. Example: `aws-eu-central-1`.",
					Type:        schema.TypeString,
				},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"azure_migration": {
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"account": {
						Description: "Azure account name.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"base_path": {
						Description: "The path to the repository data within its container. The value of this setting should not start or end with a /.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"chunk_size": {
						Description: "Big files can be broken down into chunks during snapshotting if needed. Should be the same as for the 3rd party repository.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"compress": {
						Description: "When set to true metadata files are stored in compressed format.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"container": {
						Description: "Azure container name.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"endpoint_suffix": {
						Description: "Defines the DNS suffix for Azure Storage endpoints.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"indices": {
						Description: "A comma-delimited list of indices to restore from the snapshot. Multi-index syntax is supported. By default, a restore operation includes all data streams and indices in the snapshot. If this argument is provided, the restore operation only includes the data streams and indices that you specify. Example: `metrics*,logs*,data-20240823`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"key": {
						Description: "Azure account secret key. One of key or sas_token should be specified.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sas_token": {
						Description: "A shared access signatures (SAS) token. One of key or sas_token should be specified.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"snapshot_name": {
						Description: "The snapshot name to restore from.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"custom_domain": {
				Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name. Example: `grafana.example.org`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"disable_replication_factor_adjustment": {
				Description: "Disable automatic replication factor adjustment for multi-node services. By default, Aiven ensures all indexes are replicated at least to two nodes. Note: Due to potential data loss in case of losing a service node, this setting can no longer be activated.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"gcs_migration": {
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"base_path": {
						Description: "The path to the repository data within its container. The value of this setting should not start or end with a /.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"bucket": {
						Description: "The path to the repository data within its container.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"chunk_size": {
						Description: "Big files can be broken down into chunks during snapshotting if needed. Should be the same as for the 3rd party repository.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"compress": {
						Description: "When set to true metadata files are stored in compressed format.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"credentials": {
						Description: "Google Cloud Storage credentials file content.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"indices": {
						Description: "A comma-delimited list of indices to restore from the snapshot. Multi-index syntax is supported. By default, a restore operation includes all data streams and indices in the snapshot. If this argument is provided, the restore operation only includes the data streams and indices that you specify. Example: `metrics*,logs*,data-20240823`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"snapshot_name": {
						Description: "The snapshot name to restore from.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"index_patterns": {
				Description: "Index patterns",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"max_index_count": {
						Description: "Maximum number of indexes to keep. Example: `3`.",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"pattern": {
						Description: "fnmatch pattern. Example: `logs_*_foo_*`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"sorting_algorithm": {
						Description:  "Enum: `alphabetical`, `creation_date`. Deletion sorting algorithm. Default: `creation_date`.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"alphabetical", "creation_date"}, false),
					},
				}},
				MaxItems: 512,
				Optional: true,
				Type:     schema.TypeList,
			},
			"index_rollup": {
				Description: "Index rollup settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"rollup_dashboards_enabled": {
						Description: "Whether rollups are enabled in OpenSearch Dashboards. Defaults to true.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"rollup_enabled": {
						Description: "Whether the rollup plugin is enabled. Defaults to true.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"rollup_search_backoff_count": {
						Description: "How many retries the plugin should attempt for failed rollup jobs. Defaults to 5.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"rollup_search_backoff_millis": {
						Description: "The backoff time between retries for failed rollup jobs. Defaults to 1000ms.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"rollup_search_search_all_jobs": {
						Description: "Whether OpenSearch should return all jobs that match all specified search terms. If disabled, OpenSearch returns just one, as opposed to all, of the jobs that matches the search terms. Defaults to false.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"index_template": {
				Description: "Template settings for all new indexes",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"mapping_nested_objects_limit": {
						Description: "The maximum number of nested JSON objects that a single document can contain across all nested types. This limit helps to prevent out of memory errors when a document contains too many nested objects. Default is 10000. Example: `10000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"number_of_replicas": {
						Description: "The number of replicas each primary shard has. Example: `1`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"number_of_shards": {
						Description: "The number of primary shards that an index should have. Example: `1`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
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
			"keep_index_refresh_interval": {
				Description: "Aiven automation resets index.refresh_interval to default value for every index to be sure that indices are always visible to search. If it doesn't fit your case, you can disable this by setting up this flag to true.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"max_index_count": {
				Description: "Use index_patterns instead. Default: `0`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"openid": {
				Description: "OpenSearch OpenID Connect Configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"client_id": {
						Description: "The ID of the OpenID Connect client configured in your IdP. Required.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "The client secret of the OpenID Connect client configured in your IdP. Required.",
						Required:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"connect_url": {
						Description: "The URL of your IdP where the Security plugin can find the OpenID Connect metadata/configuration settings. Example: `https://test-account.okta.com/app/exk491jujcVc83LEX697/sso/saml/metadata`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"enabled": {
						Description: "Enables or disables OpenID Connect authentication for OpenSearch. When enabled, users can authenticate using OpenID Connect with an Identity Provider. Default: `true`.",
						Required:    true,
						Type:        schema.TypeBool,
					},
					"header": {
						Description: "HTTP header name of the JWT token. Optional. Default is Authorization. Default: `Authorization`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"jwt_header": {
						Description: "The HTTP header that stores the token. Typically the Authorization header with the Bearer schema: Authorization: Bearer <token>. Optional. Default is Authorization. Example: `preferred_username`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"jwt_url_parameter": {
						Description: "If the token is not transmitted in the HTTP header, but as an URL parameter, define the name of the parameter here. Optional. Example: `preferred_username`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"refresh_rate_limit_count": {
						Description: "The maximum number of unknown key IDs in the time frame. Default is 10. Optional. Default: `10`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"refresh_rate_limit_time_window_ms": {
						Description: "The time frame to use when checking the maximum number of unknown key IDs, in milliseconds. Optional.Default is 10000 (10 seconds). Default: `10000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"roles_key": {
						Description: "The key in the JSON payload that stores the user’s roles. The value of this key must be a comma-separated list of roles. Required only if you want to use roles in the JWT. Example: `roles`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"scope": {
						Description: "The scope of the identity token issued by the IdP. Optional. Default is openid profile email address phone.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"subject_key": {
						Description: "The key in the JSON payload that stores the user’s name. If not defined, the subject registered claim is used. Most IdP providers use the preferred_username claim. Optional. Example: `preferred_username`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"opensearch": {
				Description: "OpenSearch settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"action_auto_create_index_enabled": {
						Description: "Explicitly allow or block automatic creation of indices. Defaults to true.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"action_destructive_requires_name": {
						Description: "Require explicit index names when deleting.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"auth_failure_listeners": {
						Description: "Opensearch Security Plugin Settings",
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"internal_authentication_backend_limiting": {
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"allowed_tries": {
										Description: "The number of login attempts allowed before login is blocked. Example: `10`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"authentication_backend": {
										Description:  "Enum: `internal`. internal_authentication_backend_limiting.authentication_backend.",
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"internal"}, false),
									},
									"block_expiry_seconds": {
										Description: "The duration of time that login remains blocked after a failed login. Example: `600`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_blocked_clients": {
										Description: "internal_authentication_backend_limiting.max_blocked_clients. Example: `100000`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_tracked_clients": {
										Description: "The maximum number of tracked IP addresses that have failed login. Example: `100000`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"time_window_seconds": {
										Description: "The window of time in which the value for `allowed_tries` is enforced. Example: `3600`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"type": {
										Description:  "Enum: `username`. internal_authentication_backend_limiting.type.",
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"username"}, false),
									},
								}},
								MaxItems: 1,
								Optional: true,
								Type:     schema.TypeList,
							},
							"ip_rate_limiting": {
								Description: "IP address rate limiting settings",
								Elem: &schema.Resource{Schema: map[string]*schema.Schema{
									"allowed_tries": {
										Description: "The number of login attempts allowed before login is blocked. Example: `10`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"block_expiry_seconds": {
										Description: "The duration of time that login remains blocked after a failed login. Example: `600`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_blocked_clients": {
										Description: "The maximum number of blocked IP addresses. Example: `100000`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"max_tracked_clients": {
										Description: "The maximum number of tracked IP addresses that have failed login. Example: `100000`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"time_window_seconds": {
										Description: "The window of time in which the value for `allowed_tries` is enforced. Example: `3600`.",
										Optional:    true,
										Type:        schema.TypeInt,
									},
									"type": {
										Description:  "Enum: `ip`. The type of rate limiting.",
										Optional:     true,
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"ip"}, false),
									},
								}},
								MaxItems: 1,
								Optional: true,
								Type:     schema.TypeList,
							},
						}},
						MaxItems: 1,
						Optional: true,
						Type:     schema.TypeList,
					},
					"cluster_max_shards_per_node": {
						Description: "Controls the number of shards allowed in the cluster per data node. Example: `1000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"cluster_routing_allocation_node_concurrent_recoveries": {
						Description: "How many concurrent incoming/outgoing shard recoveries (normally replicas) are allowed to happen on a node. Defaults to node cpu count * 2.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"email_sender_name": {
						Description: "Sender name placeholder to be used in Opensearch Dashboards and Opensearch keystore. Example: `alert-sender`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"email_sender_password": {
						Description: "Sender password for Opensearch alerts to authenticate with SMTP server. Example: `very-secure-mail-password`.",
						Optional:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"email_sender_username": {
						Description: "Sender username for Opensearch alerts. Example: `jane@example.com`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"enable_security_audit": {
						Description: "Enable/Disable security audit.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"http_max_content_length": {
						Description: "Maximum content length for HTTP requests to the OpenSearch HTTP API, in bytes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"http_max_header_size": {
						Description: "The max size of allowed headers, in bytes. Example: `8192`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"http_max_initial_line_length": {
						Description: "The max length of an HTTP URL, in bytes. Example: `4096`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_fielddata_cache_size": {
						Description: "Relative amount. Maximum amount of heap memory used for field data cache. This is an expert setting; decreasing the value too much will increase overhead of loading field data; too much memory used for field data cache will decrease amount of heap available for other operations.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_memory_index_buffer_size": {
						Description: "Percentage value. Default is 10%. Total amount of heap used for indexing buffer, before writing segments to disk. This is an expert setting. Too low value will slow down indexing; too high value will increase indexing performance but causes performance issues for query performance.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_memory_max_index_buffer_size": {
						Description: "Absolute value. Default is unbound. Doesn't work without indices.memory.index_buffer_size. Maximum amount of heap used for query cache, an absolute indices.memory.index_buffer_size maximum hard limit.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_memory_min_index_buffer_size": {
						Description: "Absolute value. Default is 48mb. Doesn't work without indices.memory.index_buffer_size. Minimum amount of heap used for query cache, an absolute indices.memory.index_buffer_size minimal hard limit.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_queries_cache_size": {
						Description: "Percentage value. Default is 10%. Maximum amount of heap used for query cache. This is an expert setting. Too low value will decrease query performance and increase performance for other operations; too high value will cause issues with other OpenSearch functionality.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_query_bool_max_clause_count": {
						Description: "Maximum number of clauses Lucene BooleanQuery can have. The default value (1024) is relatively high, and increasing it may cause performance issues. Investigate other approaches first before increasing this value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_recovery_max_bytes_per_sec": {
						Description: "Limits total inbound and outbound recovery traffic for each node. Applies to both peer recoveries as well as snapshot recoveries (i.e., restores from a snapshot). Defaults to 40mb.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"indices_recovery_max_concurrent_file_chunks": {
						Description: "Number of file chunks sent in parallel for each recovery. Defaults to 2.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_enabled": {
						Description: "Specifies whether ISM is enabled or not.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"ism_history_enabled": {
						Description: "Specifies whether audit history is enabled or not. The logs from ISM are automatically indexed to a logs document.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"ism_history_max_age": {
						Description: "The maximum age before rolling over the audit history index in hours. Example: `24`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_history_max_docs": {
						Description: "The maximum number of documents before rolling over the audit history index. Example: `2500000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_history_rollover_check_period": {
						Description: "The time between rollover checks for the audit history index in hours. Example: `8`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ism_history_rollover_retention_period": {
						Description: "How long audit history indices are kept in days. Example: `30`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"knn_memory_circuit_breaker_enabled": {
						Description: "Enable or disable KNN memory circuit breaker. Defaults to true.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"knn_memory_circuit_breaker_limit": {
						Description: "Maximum amount of memory that can be used for KNN index. Defaults to 50% of the JVM heap size.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"override_main_response_version": {
						Description: "Compatibility mode sets OpenSearch to report its version as 7.10 so clients continue to work. Default is false.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"plugins_alerting_filter_by_backend_roles": {
						Description: "Enable or disable filtering of alerting by backend roles. Requires Security plugin. Defaults to false.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"reindex_remote_whitelist": {
						Description: "Whitelisted addresses for reindexing. Changing this value will cause all OpenSearch instances to restart.",
						Elem: &schema.Schema{
							Description: "Address (hostname:port or IP:port). Example: `anotherservice.aivencloud.com:12398`.",
							Type:        schema.TypeString,
						},
						MaxItems: 32,
						Optional: true,
						Type:     schema.TypeList,
					},
					"script_max_compilations_rate": {
						Description: "Script compilation circuit breaker limits the number of inline script compilations within a period of time. Default is use-context. Example: `75/5m`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"search_max_buckets": {
						Description: "Maximum number of aggregation buckets allowed in a single response. OpenSearch default value is used when this is not defined. Example: `10000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_analyze_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_analyze_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_force_merge_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_get_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_get_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_throttled_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_search_throttled_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_write_queue_size": {
						Description: "Size for the thread pool queue. See documentation for exact details.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"thread_pool_write_size": {
						Description: "Size for the thread pool. See documentation for exact details. Do note this may have maximum value depending on CPU count - value is automatically lowered if set to higher than maximum value.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"opensearch_dashboards": {
				Description: "OpenSearch Dashboards settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Description: "Enable or disable OpenSearch Dashboards. Default: `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"max_old_space_size": {
						Description: "Limits the maximum amount of memory (in MiB) the OpenSearch Dashboards process can use. This sets the max_old_space_size option of the nodejs running the OpenSearch Dashboards. Note: the memory reserved by OpenSearch Dashboards is not available for OpenSearch. Default: `128`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"opensearch_request_timeout": {
						Description: "Timeout in milliseconds for requests made by OpenSearch Dashboards towards OpenSearch. Default: `30000`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"opensearch_version": {
				Description: "Enum: `1`, `2`, and newer. OpenSearch major version.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"opensearch": {
						Description: "Allow clients to connect to opensearch with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"opensearch_dashboards": {
						Description: "Allow clients to connect to opensearch_dashboards with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"privatelink_access": {
				Description: "Allow access to selected service components through Privatelink",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"opensearch": {
						Description: "Enable opensearch.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"opensearch_dashboards": {
						Description: "Enable opensearch_dashboards.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Enable prometheus.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"project_to_fork_from": {
				Description: "Name of another project to fork a service from. This has effect only when a new service is being created. Example: `anotherprojectname`.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"opensearch": {
						Description: "Allow clients to connect to opensearch from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"opensearch_dashboards": {
						Description: "Allow clients to connect to opensearch_dashboards from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"prometheus": {
						Description: "Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"recovery_basebackup_name": {
				Description: "Name of the basebackup to restore in forked service. Example: `backup-20191112t091354293891z`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"s3_migration": {
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"access_key": {
						Description: "AWS Access key.",
						Required:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"base_path": {
						Description: "The path to the repository data within its container. The value of this setting should not start or end with a /.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"bucket": {
						Description: "S3 bucket name.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"chunk_size": {
						Description: "Big files can be broken down into chunks during snapshotting if needed. Should be the same as for the 3rd party repository.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"compress": {
						Description: "When set to true metadata files are stored in compressed format.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"endpoint": {
						Description: "The S3 service endpoint to connect to. If you are using an S3-compatible service then you should set this to the service’s endpoint.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"indices": {
						Description: "A comma-delimited list of indices to restore from the snapshot. Multi-index syntax is supported. By default, a restore operation includes all data streams and indices in the snapshot. If this argument is provided, the restore operation only includes the data streams and indices that you specify. Example: `metrics*,logs*,data-20240823`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"region": {
						Description: "S3 region.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"secret_key": {
						Description: "AWS secret key.",
						Required:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"server_side_encryption": {
						Description: "When set to true files are encrypted on server side.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"snapshot_name": {
						Description: "The snapshot name to restore from.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"saml": {
				Description: "OpenSearch SAML configuration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Description: "Enables or disables SAML-based authentication for OpenSearch. When enabled, users can authenticate using SAML with an Identity Provider. Default: `true`.",
						Required:    true,
						Type:        schema.TypeBool,
					},
					"idp_entity_id": {
						Description: "The unique identifier for the Identity Provider (IdP) entity that is used for SAML authentication. This value is typically provided by the IdP. Example: `test-idp-entity-id`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"idp_metadata_url": {
						Description: "The URL of the SAML metadata for the Identity Provider (IdP). This is used to configure SAML-based authentication with the IdP. Example: `https://test-account.okta.com/app/exk491jujcVc83LEX697/sso/saml/metadata`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"idp_pemtrustedcas_content": {
						Description: "This parameter specifies the PEM-encoded root certificate authority (CA) content for the SAML identity provider (IdP) server verification. The root CA content is used to verify the SSL/TLS certificate presented by the server. Example: `-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"roles_key": {
						Description: "Optional. Specifies the attribute in the SAML response where role information is stored, if available. Role attributes are not required for SAML authentication, but can be included in SAML assertions by most Identity Providers (IdPs) to determine user access levels or permissions. Example: `RoleName`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sp_entity_id": {
						Description: "The unique identifier for the Service Provider (SP) entity that is used for SAML authentication. This value is typically provided by the SP. Example: `test-sp-entity-id`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"subject_key": {
						Description: "Optional. Specifies the attribute in the SAML response where the subject identifier is stored. If not configured, the NameID attribute is used by default. Example: `NameID`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"service_log": {
				Description: "Store logs for the service so that they are available in the HTTP API and console.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"service_to_fork_from": {
				Description: "Name of another service to fork from. This has effect only when a new service is being created. Example: `anotherservicename`.",
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
