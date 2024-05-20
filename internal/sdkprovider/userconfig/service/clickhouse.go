// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func clickhouseUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Clickhouse user configurable settings",
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
				Type:     schema.TypeSet,
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
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"clickhouse": {
						Description: "Allow clients to connect to clickhouse with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"clickhouse_https": {
						Description: "Allow clients to connect to clickhouse_https with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"clickhouse_mysql": {
						Description: "Allow clients to connect to clickhouse_mysql with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
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
					"clickhouse": {
						Description: "Enable clickhouse.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"clickhouse_https": {
						Description: "Enable clickhouse_https.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"clickhouse_mysql": {
						Description: "Enable clickhouse_mysql.",
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
					"clickhouse": {
						Description: "Allow clients to connect to clickhouse from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"clickhouse_https": {
						Description: "Allow clients to connect to clickhouse_https from the public internet for service nodes that are in a project VPC or another type of private network.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"clickhouse_mysql": {
						Description: "Allow clients to connect to clickhouse_mysql from the public internet for service nodes that are in a project VPC or another type of private network.",
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
