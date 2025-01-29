// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func valkeyUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Valkey user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
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
			"backup_hour": {
				Description: "The hour of day (in UTC) when backup for the service is started. New backup is only started if previous backup has already completed. Example: `3`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"backup_minute": {
				Description: "The minute of an hour when backup for the service is started. New backup is only started if previous backup has already completed. Example: `30`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"frequent_snapshots": {
				Description: "When enabled, Valkey will create frequent local RDB snapshots. When disabled, Valkey will only take RDB snapshots when a backup is created, based on the backup schedule. This setting is ignored when `valkey_persistence` is set to `off`. Default: `true`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field. Example: `10.20.0.0/16`.",
					Type:        schema.TypeString,
				},
				MaxItems: 2048,
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
				MaxItems: 2048,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"ip_filter_string": {
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field. Example: `10.20.0.0/16`.",
					Type:        schema.TypeString,
				},
				MaxItems: 2048,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"migration": {
				Description: "Migrate data from existing server",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"dbname": {
						Description: "Database name for bootstrapping the initial connection. Example: `defaultdb`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"host": {
						Description: "Hostname or IP address of the server where to migrate data from. Example: `my.server.com`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"ignore_dbs": {
						Description: "Comma-separated list of databases, which should be ignored during migration (supported by MySQL and PostgreSQL only at the moment). Example: `db1,db2`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"ignore_roles": {
						Description: "Comma-separated list of database roles, which should be ignored during migration (supported by PostgreSQL only at the moment). Example: `role1,role2`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"method": {
						Description:  "Enum: `dump`, `replication`. The migration method to be used (currently supported only by Redis, Dragonfly, MySQL and PostgreSQL service types).",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"dump", "replication"}, false),
					},
					"password": {
						Description: "Password for authentication with the server where to migrate data from. Example: `jjKk45Nnd`.",
						Optional:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"port": {
						Description: "Port number of the server where to migrate data from. Example: `1234`.",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"ssl": {
						Description: "The server where to migrate data from is secured with SSL. Default: `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"username": {
						Description: "User name for authentication with the server where to migrate data from. Example: `myname`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"prometheus": {
						Description: "Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"valkey": {
						Description: "Allow clients to connect to valkey with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
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
					"prometheus": {
						Description: "Enable prometheus.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"valkey": {
						Description: "Enable valkey.",
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
					"prometheus": {
						Description: "Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"valkey": {
						Description: "Allow clients to connect to valkey from the public internet for service nodes that are in a project VPC or another type of private network.",
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
			"valkey_acl_channels_default": {
				Description:  "Enum: `allchannels`, `resetchannels`. Determines default pub/sub channels' ACL for new users if ACL is not supplied. When this option is not defined, all_channels is assumed to keep backward compatibility. This option doesn't affect Valkey configuration acl-pubsub-default.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allchannels", "resetchannels"}, false),
			},
			"valkey_io_threads": {
				Description: "Set Valkey IO thread count. Changing this will cause a restart of the Valkey service. Example: `1`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"valkey_lfu_decay_time": {
				Description: "LFU maxmemory-policy counter decay time in minutes. Default: `1`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"valkey_lfu_log_factor": {
				Description: "Counter logarithm factor for volatile-lfu and allkeys-lfu maxmemory-policies. Default: `10`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"valkey_maxmemory_policy": {
				Description:  "Enum: `allkeys-lfu`, `allkeys-lru`, `allkeys-random`, `noeviction`, `volatile-lfu`, `volatile-lru`, `volatile-random`, `volatile-ttl`. Valkey maxmemory-policy. Default: `noeviction`.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allkeys-lfu", "allkeys-lru", "allkeys-random", "noeviction", "volatile-lfu", "volatile-lru", "volatile-random", "volatile-ttl"}, false),
			},
			"valkey_notify_keyspace_events": {
				Description: "Set notify-keyspace-events option.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"valkey_number_of_databases": {
				Description: "Set number of Valkey databases. Changing this will cause a restart of the Valkey service. Example: `16`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"valkey_persistence": {
				Description:  "Enum: `off`, `rdb`. When persistence is `rdb`, Valkey does RDB dumps each 10 minutes if any key is changed. Also RDB dumps are done according to backup schedule for backup purposes. When persistence is `off`, no RDB dumps and backups are done, so data can be lost at any moment if service is restarted for any reason, or if service is powered off. Also service can't be forked.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"off", "rdb"}, false),
			},
			"valkey_pubsub_client_output_buffer_limit": {
				Description: "Set output buffer limit for pub / sub clients in MB. The value is the hard limit, the soft limit is 1/4 of the hard limit. When setting the limit, be mindful of the available memory in the selected service plan. Example: `64`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"valkey_ssl": {
				Description: "Require SSL to access Valkey. Default: `true`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"valkey_timeout": {
				Description: "Valkey idle connection timeout in seconds. Default: `300`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
