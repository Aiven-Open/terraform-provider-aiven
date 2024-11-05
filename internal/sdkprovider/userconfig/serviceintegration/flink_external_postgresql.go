// Code generated by user config generator. DO NOT EDIT.

package serviceintegration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func flinkExternalPostgresqlUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "FlinkExternalPostgresql user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{"stringtype": {
			Description:  "Enum: `unspecified`. If stringtype is set to unspecified, parameters will be sent to the server as untyped values.",
			Optional:     true,
			Type:         schema.TypeString,
			ValidateFunc: validation.StringInSlice([]string{"unspecified"}, false),
		}}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
