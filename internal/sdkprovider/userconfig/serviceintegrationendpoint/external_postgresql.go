// Code generated by user config generator. DO NOT EDIT.

package serviceintegrationendpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func externalPostgresqlUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "ExternalPostgresql user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"default_database": {
				Description: "Default database. Example: `testdb`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"host": {
				Description: "Hostname or IP address of the server. Example: `my.server.com`.",
				Required:    true,
				Type:        schema.TypeString,
			},
			"password": {
				Description: "Password. Example: `jjKk45Nnd`.",
				Optional:    true,
				Sensitive:   true,
				Type:        schema.TypeString,
			},
			"port": {
				Description: "Port number of the server. Example: `5432`.",
				Required:    true,
				Type:        schema.TypeInt,
			},
			"ssl_client_certificate": {
				Description: "Client certificate. Example: `-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"ssl_client_key": {
				Description: "Client key. Example: `-----BEGIN PRIVATE KEY-----\n...\n-----END PRIVATE KEY-----`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"ssl_mode": {
				Description:  "Enum: `disable`, `allow`, `prefer`, `require`, `verify-ca`, `verify-full`. SSL mode to use for the connection.  Please note that Aiven requires TLS for all connections to external PostgreSQL services. Default: `verify-full`.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "allow", "prefer", "require", "verify-ca", "verify-full"}, false),
			},
			"ssl_root_cert": {
				Description: "SSL Root Cert. Example: `-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----\n`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"username": {
				Description: "User name. Example: `myname`.",
				Required:    true,
				Type:        schema.TypeString,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
