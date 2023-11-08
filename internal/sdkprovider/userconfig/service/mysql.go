// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func mysqlUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Mysql user configurable settings",
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
			"admin_password": {
				Description: "Custom password for admin user. Defaults to random string. This must be set only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"admin_username": {
				Description: "Custom username for admin user. This must be set only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"backup_hour": {
				Description: "The hour of day (in UTC) when backup for the service is started. New backup is only started if previous backup has already completed.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"backup_minute": {
				Description: "The minute of an hour when backup for the service is started. New backup is only started if previous backup has already completed.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"binlog_retention_period": {
				Description: "The minimum amount of time in seconds to keep binlog entries before deletion. This may be extended for services that require binlog entries for longer than the default for example if using the MySQL Debezium Kafka connector.",
				Optional:    true,
				Type:        schema.TypeInt,
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
			"migration": {
				Description: "Migrate data from existing server",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"dbname": {
						Description: "Database name for bootstrapping the initial connection.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"host": {
						Description: "Hostname or IP address of the server where to migrate data from.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"ignore_dbs": {
						Description: "Comma-separated list of databases, which should be ignored during migration (supported by MySQL and PostgreSQL only at the moment).",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"method": {
						Description:  "The migration method to be used (currently supported only by Redis, Dragonfly, MySQL and PostgreSQL service types).",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"dump", "replication"}, false),
					},
					"password": {
						Description: "Password for authentication with the server where to migrate data from.",
						Optional:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"port": {
						Description: "Port number of the server where to migrate data from.",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"ssl": {
						Description: "The server where to migrate data from is secured with SSL. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"username": {
						Description: "User name for authentication with the server where to migrate data from.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"mysql": {
				Description: "mysql.conf configuration values",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"connect_timeout": {
						Description: "The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"default_time_zone": {
						Description: "Default server time zone as an offset from UTC (from -12:00 to +12:00), a time zone name, or 'SYSTEM' to use the MySQL server default.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"group_concat_max_len": {
						Description: "The maximum permitted result length in bytes for the GROUP_CONCAT() function.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"information_schema_stats_expiry": {
						Description: "The time, in seconds, before cached statistics expire.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_change_buffer_max_size": {
						Description: "Maximum size for the InnoDB change buffer, as a percentage of the total size of the buffer pool. Default is 25.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_flush_neighbors": {
						Description: "Specifies whether flushing a page from the InnoDB buffer pool also flushes other dirty pages in the same extent (default is 1): 0 - dirty pages in the same extent are not flushed,  1 - flush contiguous dirty pages in the same extent,  2 - flush dirty pages in the same extent.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_ft_min_token_size": {
						Description: "Minimum length of words that are stored in an InnoDB FULLTEXT index. Changing this parameter will lead to a restart of the MySQL service.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_ft_server_stopword_table": {
						Description: "This option is used to specify your own InnoDB FULLTEXT index stopword list for all InnoDB tables.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"innodb_lock_wait_timeout": {
						Description: "The length of time in seconds an InnoDB transaction waits for a row lock before giving up. Default is 120.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_log_buffer_size": {
						Description: "The size in bytes of the buffer that InnoDB uses to write to the log files on disk.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_online_alter_log_max_size": {
						Description: "The upper limit in bytes on the size of the temporary log files used during online DDL operations for InnoDB tables.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_print_all_deadlocks": {
						Description: "When enabled, information about all deadlocks in InnoDB user transactions is recorded in the error log. Disabled by default.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"innodb_read_io_threads": {
						Description: "The number of I/O threads for read operations in InnoDB. Default is 4. Changing this parameter will lead to a restart of the MySQL service.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_rollback_on_timeout": {
						Description: "When enabled a transaction timeout causes InnoDB to abort and roll back the entire transaction. Changing this parameter will lead to a restart of the MySQL service.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"innodb_thread_concurrency": {
						Description: "Defines the maximum number of threads permitted inside of InnoDB. Default is 0 (infinite concurrency - no limit).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"innodb_write_io_threads": {
						Description: "The number of I/O threads for write operations in InnoDB. Default is 4. Changing this parameter will lead to a restart of the MySQL service.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"interactive_timeout": {
						Description: "The number of seconds the server waits for activity on an interactive connection before closing it.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"internal_tmp_mem_storage_engine": {
						Description:  "The storage engine for in-memory internal temporary tables.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"TempTable", "MEMORY"}, false),
					},
					"long_query_time": {
						Description: "The slow_query_logs work as SQL statements that take more than long_query_time seconds to execute. Default is 10s.",
						Optional:    true,
						Type:        schema.TypeFloat,
					},
					"max_allowed_packet": {
						Description: "Size of the largest message in bytes that can be received by the server. Default is 67108864 (64M).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_heap_table_size": {
						Description: "Limits the size of internal in-memory tables. Also set tmp_table_size. Default is 16777216 (16M).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"net_buffer_length": {
						Description: "Start sizes of connection buffer and result buffer. Default is 16384 (16K). Changing this parameter will lead to a restart of the MySQL service.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"net_read_timeout": {
						Description: "The number of seconds to wait for more data from a connection before aborting the read.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"net_write_timeout": {
						Description: "The number of seconds to wait for a block to be written to a connection before aborting the write.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"slow_query_log": {
						Description: "Slow query log enables capturing of slow queries. Setting slow_query_log to false also truncates the mysql.slow_log table. Default is off.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"sort_buffer_size": {
						Description: "Sort buffer size in bytes for ORDER BY optimization. Default is 262144 (256K).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"sql_mode": {
						Description: "Global SQL mode. Set to empty to use MySQL server defaults. When creating a new service and not setting this field Aiven default SQL mode (strict, SQL standard compliant) will be assigned.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"sql_require_primary_key": {
						Description: "Require primary key to be defined for new tables or old tables modified with ALTER TABLE and fail if missing. It is recommended to always have primary keys because various functionality may break if any large table is missing them.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"tmp_table_size": {
						Description: "Limits the size of internal in-memory tables. Also set max_heap_table_size. Default is 16777216 (16M).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"wait_timeout": {
						Description: "The number of seconds the server waits for activity on a noninteractive connection before closing it.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"mysql_version": {
				Description:  "MySQL major version.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"8"}, false),
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"mysql": {
						Description: "Allow clients to connect to mysql with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"mysqlx": {
						Description: "Allow clients to connect to mysqlx with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
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
					"mysql": {
						Description: "Enable mysql.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"mysqlx": {
						Description: "Enable mysqlx.",
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
				Description: "Name of another project to fork a service from. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"mysql": {
						Description: "Allow clients to connect to mysql from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"mysqlx": {
						Description: "Allow clients to connect to mysqlx from the public internet for service nodes that are in a project VPC or another type of private network.",
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
			"recovery_target_time": {
				Description: "Recovery target time when forking a service. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
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
