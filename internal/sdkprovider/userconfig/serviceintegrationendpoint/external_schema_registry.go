// Code generated by user config generator. DO NOT EDIT.

package serviceintegrationendpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func externalSchemaRegistryUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "ExternalSchemaRegistry user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"authentication": {
				Description:  "Enum: `basic`, `none`. Authentication method.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"basic", "none"}, false),
			},
			"basic_auth_password": {
				Description: "Basic authentication password. Example: `Zm9vYg==`.",
				Optional:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"basic_auth_username": {
				Description: "Basic authentication user name. Example: `avnadmin`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"url": {
				Description: "Schema Registry URL. Example: `https://schema-registry.kafka.company.com:28419`.",
				Required:    true,
				Type:        schema.TypeString,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
