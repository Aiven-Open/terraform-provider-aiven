// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func kafkaMirrormakerUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "KafkaMirrormaker user configurable settings",
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
			"kafka_mirrormaker": {
				Description: "Kafka MirrorMaker configuration values",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"emit_checkpoints_enabled": {
						Description: "Whether to emit consumer group offset checkpoints to target cluster periodically (default: true).",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"emit_checkpoints_interval_seconds": {
						Description: "Frequency at which consumer group offset checkpoints are emitted (default: 60, every minute).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"refresh_groups_enabled": {
						Description: "Whether to periodically check for new consumer groups. Defaults to 'true'.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"refresh_groups_interval_seconds": {
						Description: "Frequency of consumer group refresh in seconds. Defaults to 600 seconds (10 minutes).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"refresh_topics_enabled": {
						Description: "Whether to periodically check for new topics and partitions. Defaults to 'true'.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"refresh_topics_interval_seconds": {
						Description: "Frequency of topic and partitions refresh in seconds. Defaults to 600 seconds (10 minutes).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"sync_group_offsets_enabled": {
						Description: "Whether to periodically write the translated offsets of replicated consumer groups (in the source cluster) to __consumer_offsets topic in target cluster, as long as no active consumers in that group are connected to the target cluster.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"sync_group_offsets_interval_seconds": {
						Description: "Frequency at which consumer group offsets are synced (default: 60, every minute).",
						Optional:    true,
						Type:        schema.TypeInt,
					},
					"sync_topic_configs_enabled": {
						Description: "Whether to periodically configure remote topics to match their corresponding upstream topics.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"tasks_max_per_cpu": {
						Description: "'tasks.max' is set to this multiplied by the number of CPUs in the service. The default value is `1`.",
						Optional:    true,
						Type:        schema.TypeInt,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
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
