// Code generated by user config generator. DO NOT EDIT.

package serviceintegration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func kafkaConnectUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "KafkaConnect user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{"kafka_connect": {
			Description: "Kafka Connect service configuration values",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"config_storage_topic": {
					Description: "The name of the topic where connector and task configuration data are stored.This must be the same for all workers with the same group_id. Example: `__connect_configs`.",
					Optional:    true,
					Type:        schema.TypeString,
				},
				"group_id": {
					Description: "A unique string that identifies the Connect cluster group this worker belongs to. Example: `connect`.",
					Optional:    true,
					Type:        schema.TypeString,
				},
				"offset_storage_topic": {
					Description: "The name of the topic where connector and task configuration offsets are stored.This must be the same for all workers with the same group_id. Example: `__connect_offsets`.",
					Optional:    true,
					Type:        schema.TypeString,
				},
				"status_storage_topic": {
					Description: "The name of the topic where connector and task configuration status updates are stored.This must be the same for all workers with the same group_id. Example: `__connect_status`.",
					Optional:    true,
					Type:        schema.TypeString,
				},
			}},
			MaxItems: 1,
			Optional: true,
			Type:     schema.TypeList,
		}}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
