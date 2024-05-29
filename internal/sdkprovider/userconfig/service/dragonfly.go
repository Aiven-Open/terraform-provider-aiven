// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func dragonflyUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Dragonfly user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"cache_mode": {
				Description: "Evict entries when getting close to maxmemory limit. The default value is `false`.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"dragonfly_persistence": {
				Description:  "Enum: `off`, `rdb`, `dfs`. When persistence is `rdb` or `dfs`, Dragonfly does RDB or DFS dumps every 10 minutes. Dumps are done according to the backup schedule for backup purposes. When persistence is `off`, no RDB/DFS dumps or backups are done, so data can be lost at any moment if the service is restarted for any reason, or if the service is powered off. Also, the service can't be forked.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"off", "rdb", "dfs"}, false),
			},
			"dragonfly_ssl": {
				Description: "Require SSL to access Dragonfly. The default value is `true`.",
				Optional:    true,
				Type:        schema.TypeBool,
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
						Description:  "Enum: `dump`, `replication`. The migration method to be used (currently supported only by Redis, Dragonfly, MySQL and PostgreSQL service types).",
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
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"dragonfly": {
						Description: "Allow clients to connect to dragonfly with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
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
					"dragonfly": {
						Description: "Enable dragonfly.",
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
					"dragonfly": {
						Description: "Allow clients to connect to dragonfly from the public internet for service nodes that are in a project VPC or another type of private network.",
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
				Description: "Name of the basebackup to restore in forked service.",
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
