// Code generated by user config generator. DO NOT EDIT.

package serviceintegrationendpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func rsyslogUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Rsyslog user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"ca": {
				Description: "PEM encoded CA certificate. Example: `-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"cert": {
				Description: "PEM encoded client certificate. Example: `-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"format": {
				Description:  "Enum: `custom`, `rfc3164`, `rfc5424`. Message format. Default: `rfc5424`.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"custom", "rfc3164", "rfc5424"}, false),
			},
			"key": {
				Description: "PEM encoded client key. Example: `-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----\n`.",
				Optional:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"logline": {
				Description: "Custom syslog message format. Example: `<%pri%>%timestamp:::date-rfc3339% %HOSTNAME% %app-name% %msg%`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"max_message_size": {
				Description: "Rsyslog max message size. Default: `8192`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"port": {
				Description: "Rsyslog server port. Default: `514`.",
				Required:    true,
				Type:        schema.TypeInt,
			},
			"sd": {
				Description: "Structured data block for log message. Example: `TOKEN tag=\"LiteralValue\"`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"server": {
				Description: "Rsyslog server IP address or hostname. Example: `logs.example.com`.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"tls": {
				Description: "Require TLS. Default: `true`.",
				Required:    true,
				Type:        schema.TypeBool,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
