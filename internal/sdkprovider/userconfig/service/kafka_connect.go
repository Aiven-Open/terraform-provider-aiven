// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func kafkaConnectUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "KafkaConnect user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"additional_backup_regions": {
				Deprecated:  "This property is deprecated.",
				Description: "Additional Cloud Regions for Backup Replication.",
				Elem: &schema.Schema{
					Description: "Target cloud.",
					Type:        schema.TypeString,
				},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
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
				Type:     schema.TypeSet,
			},
			"ip_filter_string": {
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"kafka_connect": {
				Description: "Kafka Connect configuration values",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"connector_client_config_override_policy": {
						Description:  "Enum: `None`, `All`. Defines what client configurations can be overridden by the connector. Default is None.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"None", "All"}, false),
					},
					"consumer_auto_offset_reset": {
						Description:  "Enum: `earliest`, `latest`. What to do when there is no initial offset in Kafka or if the current offset does not exist any more on the server. Default is earliest.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"earliest", "latest"}, false),
					},
					"consumer_fetch_max_bytes": {
						Description: "Records are fetched in batches by the consumer, and if the first record batch in the first non-empty partition of the fetch is larger than this value, the record batch will still be returned to ensure that the consumer can make progress. As such, this is not a absolute maximum.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_isolation_level": {
						Description:  "Enum: `read_uncommitted`, `read_committed`. Transaction read isolation level. read_uncommitted is the default, but read_committed can be used if consume-exactly-once behavior is desired.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"read_uncommitted", "read_committed"}, false),
					},
					"consumer_max_partition_fetch_bytes": {
						Description: "Records are fetched in batches by the consumer.If the first record batch in the first non-empty partition of the fetch is larger than this limit, the batch will still be returned to ensure that the consumer can make progress.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_max_poll_interval_ms": {
						Description: "The maximum delay in milliseconds between invocations of poll() when using consumer group management (defaults to 300000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"consumer_max_poll_records": {
						Description: "The maximum number of records returned in a single call to poll() (defaults to 500).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"offset_flush_interval_ms": {
						Description: "The interval at which to try committing offsets for tasks (defaults to 60000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"offset_flush_timeout_ms": {
						Description: "Maximum number of milliseconds to wait for records to flush and partition offset data to be committed to offset storage before cancelling the process and restoring the offset data to be committed in a future attempt (defaults to 5000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_batch_size": {
						Description: "This setting gives the upper bound of the batch size to be sent. If there are fewer than this many bytes accumulated for this partition, the producer will `linger` for the linger.ms time waiting for more records to show up. A batch size of zero will disable batching entirely (defaults to 16384).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_buffer_memory": {
						Description: "The total bytes of memory the producer can use to buffer records waiting to be sent to the broker (defaults to 33554432).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_compression_type": {
						Description:  "Enum: `gzip`, `snappy`, `lz4`, `zstd`, `none`. Specify the default compression type for producers. This configuration accepts the standard compression codecs (`gzip`, `snappy`, `lz4`, `zstd`). It additionally accepts `none` which is the default and equivalent to no compression.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"gzip", "snappy", "lz4", "zstd", "none"}, false),
					},
					"producer_linger_ms": {
						Description: "This setting gives the upper bound on the delay for batching: once there is batch.size worth of records for a partition it will be sent immediately regardless of this setting, however if there are fewer than this many bytes accumulated for this partition the producer will `linger` for the specified time waiting for more records to show up. Defaults to 0.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"producer_max_request_size": {
						Description: "This setting will limit the number of record batches the producer will send in a single request to avoid sending huge requests.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"scheduled_rebalance_max_delay_ms": {
						Description: "The maximum delay that is scheduled in order to wait for the return of one or more departed workers before rebalancing and reassigning their connectors and tasks to the group. During this period the connectors and tasks of the departed workers remain unassigned. Defaults to 5 minutes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"session_timeout_ms": {
						Description: "The timeout in milliseconds used to detect failures when using Kafka’s group management facilities (defaults to 10000).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"kafka_connect": {
						Description: "Allow clients to connect to kafka_connect with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
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
					"jolokia": {
						Description: "Enable jolokia.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"kafka_connect": {
						Description: "Enable kafka_connect.",
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
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"kafka_connect": {
						Description: "Allow clients to connect to kafka_connect from the public internet for service nodes that are in a project VPC or another type of private network.",
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
