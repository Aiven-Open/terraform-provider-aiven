// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func pgUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Pg user configurable settings",
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
				Type:     schema.TypeList,
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
			"enable_ipv6": {
				Description: "Register AAAA DNS records for the service, and allow IPv6 packets to service ports.",
				Optional:    true,
				Type:        schema.TypeBool,
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
			"pg": {
				Description: "postgresql.conf configuration values",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"autovacuum_analyze_scale_factor": {
						Description: "Specifies a fraction of the table size to add to autovacuum_analyze_threshold when deciding whether to trigger an ANALYZE. The default is 0.2 (20% of table size).",
						Optional:    true,
						Type:        schema.TypeFloat,
					},
					"autovacuum_analyze_threshold": {
						Description: "Specifies the minimum number of inserted, updated or deleted tuples needed to trigger an  ANALYZE in any one table. The default is 50 tuples.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autovacuum_freeze_max_age": {
						Description: "Specifies the maximum age (in transactions) that a table's pg_class.relfrozenxid field can attain before a VACUUM operation is forced to prevent transaction ID wraparound within the table. Note that the system will launch autovacuum processes to prevent wraparound even when autovacuum is otherwise disabled. This parameter will cause the server to be restarted.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autovacuum_max_workers": {
						Description: "Specifies the maximum number of autovacuum processes (other than the autovacuum launcher) that may be running at any one time. The default is three. This parameter can only be set at server start.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autovacuum_naptime": {
						Description: "Specifies the minimum delay between autovacuum runs on any given database. The delay is measured in seconds, and the default is one minute.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autovacuum_vacuum_cost_delay": {
						Description: "Specifies the cost delay value that will be used in automatic VACUUM operations. If -1 is specified, the regular vacuum_cost_delay value will be used. The default value is 20 milliseconds.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autovacuum_vacuum_cost_limit": {
						Description: "Specifies the cost limit value that will be used in automatic VACUUM operations. If -1 is specified (which is the default), the regular vacuum_cost_limit value will be used.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autovacuum_vacuum_scale_factor": {
						Description: "Specifies a fraction of the table size to add to autovacuum_vacuum_threshold when deciding whether to trigger a VACUUM. The default is 0.2 (20% of table size).",
						Optional:    true,
						Type:        schema.TypeFloat,
					},
					"autovacuum_vacuum_threshold": {
						Description: "Specifies the minimum number of updated or deleted tuples needed to trigger a VACUUM in any one table. The default is 50 tuples.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"bgwriter_delay": {
						Description: "Specifies the delay between activity rounds for the background writer in milliseconds. Default is 200.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"bgwriter_flush_after": {
						Description: "Whenever more than bgwriter_flush_after bytes have been written by the background writer, attempt to force the OS to issue these writes to the underlying storage. Specified in kilobytes, default is 512. Setting of 0 disables forced writeback.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"bgwriter_lru_maxpages": {
						Description: "In each round, no more than this many buffers will be written by the background writer. Setting this to zero disables background writing. Default is 100.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"bgwriter_lru_multiplier": {
						Description: "The average recent need for new buffers is multiplied by bgwriter_lru_multiplier to arrive at an estimate of the number that will be needed during the next round, (up to bgwriter_lru_maxpages). 1.0 represents a “just in time” policy of writing exactly the number of buffers predicted to be needed. Larger values provide some cushion against spikes in demand, while smaller values intentionally leave writes to be done by server processes. The default is 2.0.",
						Optional:    true,
						Type:        schema.TypeFloat,
					},
					"deadlock_timeout": {
						Description: "This is the amount of time, in milliseconds, to wait on a lock before checking to see if there is a deadlock condition.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"default_toast_compression": {
						Description:  "Specifies the default TOAST compression method for values of compressible columns (the default is lz4).",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"lz4", "pglz"}, false),
					},
					"idle_in_transaction_session_timeout": {
						Description: "Time out sessions with open transactions after this number of milliseconds.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"jit": {
						Description: "Controls system-wide use of Just-in-Time Compilation (JIT).",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"log_autovacuum_min_duration": {
						Description: "Causes each action executed by autovacuum to be logged if it ran for at least the specified number of milliseconds. Setting this to zero logs all autovacuum actions. Minus-one (the default) disables logging autovacuum actions.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_error_verbosity": {
						Description:  "Controls the amount of detail written in the server log for each message that is logged.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"TERSE", "DEFAULT", "VERBOSE"}, false),
					},
					"log_line_prefix": {
						Description:  "Choose from one of the available log-formats. These can support popular log analyzers like pgbadger, pganalyze etc.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"'pid=%p,user=%u,db=%d,app=%a,client=%h '", "'%t [%p]: [%l-1] user=%u,db=%d,app=%a,client=%h '", "'%m [%p] %q[user=%u,db=%d,app=%a] '"}, false),
					},
					"log_min_duration_statement": {
						Description: "Log statements that take more than this number of milliseconds to run, -1 disables.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"log_temp_files": {
						Description: "Log statements for each temporary file created larger than this number of kilobytes, -1 disables.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_files_per_process": {
						Description: "PostgreSQL maximum number of files that can be open per process.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_locks_per_transaction": {
						Description: "PostgreSQL maximum locks per transaction.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_logical_replication_workers": {
						Description: "PostgreSQL maximum logical replication workers (taken from the pool of max_parallel_workers).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_parallel_workers": {
						Description: "Sets the maximum number of workers that the system can support for parallel queries.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_parallel_workers_per_gather": {
						Description: "Sets the maximum number of workers that can be started by a single Gather or Gather Merge node.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_pred_locks_per_transaction": {
						Description: "PostgreSQL maximum predicate locks per transaction.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_prepared_transactions": {
						Description: "PostgreSQL maximum prepared transactions.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_replication_slots": {
						Description: "PostgreSQL maximum replication slots.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_slot_wal_keep_size": {
						Description: "PostgreSQL maximum WAL size (MB) reserved for replication slots. Default is -1 (unlimited). wal_keep_size minimum WAL size setting takes precedence over this.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_stack_depth": {
						Description: "Maximum depth of the stack in bytes.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_standby_archive_delay": {
						Description: "Max standby archive delay in milliseconds.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_standby_streaming_delay": {
						Description: "Max standby streaming delay in milliseconds.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_wal_senders": {
						Description: "PostgreSQL maximum WAL senders.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"max_worker_processes": {
						Description: "Sets the maximum number of background processes that the system can support.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"pg_partman_bgw__dot__interval": {
						Description: "Sets the time interval to run pg_partman's scheduled tasks.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"pg_partman_bgw__dot__role": {
						Description: "Controls which role to use for pg_partman's scheduled background tasks.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"pg_stat_monitor__dot__pgsm_enable_query_plan": {
						Description: "Enables or disables query plan monitoring.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"pg_stat_monitor__dot__pgsm_max_buckets": {
						Description: "Sets the maximum number of buckets .",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"pg_stat_statements__dot__track": {
						Description:  "Controls which statements are counted. Specify top to track top-level statements (those issued directly by clients), all to also track nested statements (such as statements invoked within functions), or none to disable statement statistics collection. The default value is top.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"all", "top", "none"}, false),
					},
					"temp_file_limit": {
						Description: "PostgreSQL temporary file limit in KiB, -1 for unlimited.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"timezone": {
						Description: "PostgreSQL service timezone.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"track_activity_query_size": {
						Description: "Specifies the number of bytes reserved to track the currently executing command for each active session.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"track_commit_timestamp": {
						Description:  "Record commit time of transactions.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"off", "on"}, false),
					},
					"track_functions": {
						Description:  "Enables tracking of function call counts and time used.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"all", "pl", "none"}, false),
					},
					"track_io_timing": {
						Description:  "Enables timing of database I/O calls. This parameter is off by default, because it will repeatedly query the operating system for the current time, which may cause significant overhead on some platforms.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"off", "on"}, false),
					},
					"wal_sender_timeout": {
						Description: "Terminate replication connections that are inactive for longer than this amount of time, in milliseconds. Setting this value to zero disables the timeout.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"wal_writer_delay": {
						Description: "WAL flush interval in milliseconds. Note that setting this value to lower than the default 200ms may negatively impact performance.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"pg_qualstats": {
				Deprecated:  "This property is deprecated.",
				Description: "System-wide settings for the pg_qualstats extension",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"enabled": {
						Deprecated:  "This property is deprecated.",
						Description: "Enable / Disable pg_qualstats. The default value is `false`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"min_err_estimate_num": {
						Deprecated:  "This property is deprecated.",
						Description: "Error estimation num threshold to save quals. The default value is `0`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"min_err_estimate_ratio": {
						Deprecated:  "This property is deprecated.",
						Description: "Error estimation ratio threshold to save quals. The default value is `0`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"track_constants": {
						Deprecated:  "This property is deprecated.",
						Description: "Enable / Disable pg_qualstats constants tracking. The default value is `true`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"track_pg_catalog": {
						Deprecated:  "This property is deprecated.",
						Description: "Track quals on system catalogs too. The default value is `false`.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"pg_read_replica": {
				Description: "Should the service which is being forked be a read replica (deprecated, use read_replica service integration instead).",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"pg_service_to_fork_from": {
				Description: "Name of the PG Service from which to fork (deprecated, use service_to_fork_from). This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"pg_stat_monitor_enable": {
				Description: "Enable the pg_stat_monitor extension. Enabling this extension will cause the cluster to be restarted.When this extension is enabled, pg_stat_statements results for utility commands are unreliable. The default value is `false`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"pg_version": {
				Description:  "PostgreSQL major version.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"11", "12", "13", "14", "15", "10", "16"}, false),
			},
			"pgbouncer": {
				Description: "PGBouncer connection pooling settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"autodb_idle_timeout": {
						Description: "If the automatically created database pools have been unused this many seconds, they are freed. If 0 then timeout is disabled. (seconds).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autodb_max_db_connections": {
						Description: "Do not allow more than this many server connections per database (regardless of user). Setting it to 0 means unlimited.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"autodb_pool_mode": {
						Description:  "PGBouncer pool mode.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"session", "transaction", "statement"}, false),
					},
					"autodb_pool_size": {
						Description: "If non-zero then create automatically a pool of that size per user when a pool doesn't exist.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"ignore_startup_parameters": {
						Description: "List of parameters to ignore when given in startup packet.",
						Elem: &schema.Schema{
							Description:  "Enum of parameters to ignore when given in startup packet.",
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"extra_float_digits", "search_path"}, false),
						},
						MaxItems: 32,
						Optional: true,
						Type:     schema.TypeList,
					},
					"min_pool_size": {
						Description: "Add more server connections to pool if below this number. Improves behavior when usual load comes suddenly back after period of total inactivity. The value is effectively capped at the pool size.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"server_idle_timeout": {
						Description: "If a server connection has been idle more than this many seconds it will be dropped. If 0 then timeout is disabled. (seconds).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"server_lifetime": {
						Description: "The pooler will close an unused server connection that has been connected longer than this. (seconds).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"server_reset_query_always": {
						Description: "Run server_reset_query (DISCARD ALL) in all pooling modes.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"pglookout": {
				Description: "System-wide settings for pglookout",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"max_failover_replication_time_lag": {
					Description: "Number of seconds of master unavailability before triggering database failover to standby. The default value is `60`.",
					Optional:    true,
					Type:        schema.TypeInt,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"pg": {
						Description: "Allow clients to connect to pg with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"pgbouncer": {
						Description: "Allow clients to connect to pgbouncer with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
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
					"pg": {
						Description: "Enable pg.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"pgbouncer": {
						Description: "Enable pgbouncer.",
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
					"pg": {
						Description: "Allow clients to connect to pg from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"pgbouncer": {
						Description: "Allow clients to connect to pgbouncer from the public internet for service nodes that are in a project VPC or another type of private network.",
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
			"service_log": {
				Description: "Store logs for the service so that they are available in the HTTP API and console.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"service_to_fork_from": {
				Description: "Name of another service to fork from. This has effect only when a new service is being created.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"shared_buffers_percentage": {
				Description: "Percentage of total RAM that the database server uses for shared memory buffers. Valid range is 20-60 (float), which corresponds to 20% - 60%. This setting adjusts the shared_buffers configuration value.",
				Optional:    true,
				Type:        schema.TypeFloat,
			},
			"static_ips": {
				Description: "Use static public IP addresses.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"synchronous_replication": {
				Description:  "Synchronous replication type. Note that the service plan also needs to support synchronous replication.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"quorum", "off"}, false),
			},
			"timescaledb": {
				Description: "System-wide settings for the timescaledb extension",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"max_background_workers": {
					Description: "The number of background workers for timescaledb operations. You should configure this setting to the sum of your number of databases and the total number of concurrent background workers you want running at any given point in time. The default value is `16`.",
					Optional:    true,
					Type:        schema.TypeInt,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"variant": {
				Description:  "Variant of the PostgreSQL service, may affect the features that are exposed by default.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"aiven", "timescale"}, false),
			},
			"work_mem": {
				Description: "Sets the maximum amount of memory to be used by a query operation (such as a sort or hash table) before writing to temporary disk files, in MB. Default is 1MB + 0.075% of total RAM (up to 32MB).",
				Optional:    true,
				Type:        schema.TypeInt,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
