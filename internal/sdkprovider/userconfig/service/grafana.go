// Code generated by user config generator. DO NOT EDIT.

package service

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func grafanaUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Grafana user configurable settings",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"additional_backup_regions": {
				Description: "Additional Cloud Regions for Backup Replication.",
				Elem: &schema.Schema{
					Description: "Target cloud. Example: `aws-eu-central-1`.",
					Type:        schema.TypeString,
				},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"alerting_enabled": {
				Description: "Enable or disable Grafana legacy alerting functionality. This should not be enabled with unified_alerting_enabled.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"alerting_error_or_timeout": {
				Description:  "Enum: `alerting`, `keep_state`. Default error or timeout setting for new alerting rules.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"alerting", "keep_state"}, false),
			},
			"alerting_max_annotations_to_keep": {
				Description: "Max number of alert annotations that Grafana stores. 0 (default) keeps all alert annotations. Example: `0`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"alerting_nodata_or_nullvalues": {
				Description:  "Enum: `alerting`, `no_data`, `keep_state`, `ok`. Default value for 'no data or null values' for new alerting rules.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"alerting", "no_data", "keep_state", "ok"}, false),
			},
			"allow_embedding": {
				Description: "Allow embedding Grafana dashboards with iframe/frame/object/embed tags. Disabled by default to limit impact of clickjacking.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"auth_azuread": {
				Description: "Azure AD OAuth integration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allow_sign_up": {
						Description: "Automatically sign-up users on successful sign-in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"allowed_domains": {
						Description: "Allowed domains.",
						Elem: &schema.Schema{
							Description: "Allowed domain. Example: `mycompany.com`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
					"allowed_groups": {
						Description: "Require users to belong to one of given groups.",
						Elem: &schema.Schema{
							Description: "Group Object ID from Azure AD. Example: `c0ffee15-c01d-0000-1111-012345abcdef`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
					"auth_url": {
						Description: "Authorization URL. Example: `https://login.microsoftonline.com/<AZURE_TENANT_ID>/oauth2/v2.0/authorize`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_id": {
						Description: "Client ID from provider. Example: `b1ba0bf54a4c2c0a1c29`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "Client secret from provider. Example: `bfa6gea4f129076761dcba8ce5e1e406bd83af7b`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"token_url": {
						Description: "Token URL. Example: `https://login.microsoftonline.com/<AZURE_TENANT_ID>/oauth2/v2.0/token`.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"auth_basic_enabled": {
				Description: "Enable or disable basic authentication form, used by Grafana built-in login.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"auth_generic_oauth": {
				Description: "Generic OAuth integration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allow_sign_up": {
						Description: "Automatically sign-up users on successful sign-in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"allowed_domains": {
						Description: "Allowed domains.",
						Elem: &schema.Schema{
							Description: "Allowed domain. Example: `mycompany.com`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
					"allowed_organizations": {
						Description: "Require user to be member of one of the listed organizations.",
						Elem: &schema.Schema{
							Description: "Allowed organization. Example: `myorg`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
					"api_url": {
						Description: "API URL. Example: `https://yourprovider.com/api`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"auth_url": {
						Description: "Authorization URL. Example: `https://yourprovider.com/oauth/authorize`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"auto_login": {
						Description: "Allow users to bypass the login screen and automatically log in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"client_id": {
						Description: "Client ID from provider. Example: `b1ba0bf54a4c2c0a1c29`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "Client secret from provider. Example: `bfa6gea4f129076761dcba8ce5e1e406bd83af7b`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"name": {
						Description: "Name of the OAuth integration. Example: `My authentication`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"scopes": {
						Description: "OAuth scopes.",
						Elem: &schema.Schema{
							Description: "OAuth scope. Example: `email`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
					"token_url": {
						Description: "Token URL. Example: `https://yourprovider.com/oauth/token`.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"auth_github": {
				Description: "Github Auth integration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allow_sign_up": {
						Description: "Automatically sign-up users on successful sign-in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"allowed_organizations": {
						Description: "Require users to belong to one of given organizations.",
						Elem: &schema.Schema{
							Description: "Organization name. Example: `aiven`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
					"auto_login": {
						Description: "Allow users to bypass the login screen and automatically log in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"client_id": {
						Description: "Client ID from provider. Example: `b1ba0bf54a4c2c0a1c29`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "Client secret from provider. Example: `bfa6gea4f129076761dcba8ce5e1e406bd83af7b`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"skip_org_role_sync": {
						Description: "Stop automatically syncing user roles.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"team_ids": {
						Description: "Require users to belong to one of given team IDs.",
						Elem: &schema.Schema{
							Description: "Team ID. Example: `150`.",
							Type:        schema.TypeInt,
						},
						MaxItems: 50,
						Optional: true,
						Type:     schema.TypeList,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"auth_gitlab": {
				Description: "GitLab Auth integration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allow_sign_up": {
						Description: "Automatically sign-up users on successful sign-in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"allowed_groups": {
						Description: "Require users to belong to one of given groups.",
						Elem: &schema.Schema{
							Description: "Group or subgroup name. Example: `aiven/developers`.",
							Type:        schema.TypeString,
						},
						MaxItems: 50,
						Required: true,
						Type:     schema.TypeList,
					},
					"api_url": {
						Description: "API URL. This only needs to be set when using self hosted GitLab. Example: `https://gitlab.com/api/v4`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"auth_url": {
						Description: "Authorization URL. This only needs to be set when using self hosted GitLab. Example: `https://gitlab.com/oauth/authorize`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"client_id": {
						Description: "Client ID from provider. Example: `b1ba0bf54a4c2c0a1c29`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "Client secret from provider. Example: `bfa6gea4f129076761dcba8ce5e1e406bd83af7b`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"token_url": {
						Description: "Token URL. This only needs to be set when using self hosted GitLab. Example: `https://gitlab.com/oauth/token`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"auth_google": {
				Description: "Google Auth integration",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"allow_sign_up": {
						Description: "Automatically sign-up users on successful sign-in.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"allowed_domains": {
						Description: "Domains allowed to sign-in to this Grafana.",
						Elem: &schema.Schema{
							Description: "Domain. Example: `example.com`.",
							Type:        schema.TypeString,
						},
						MaxItems: 64,
						Required: true,
						Type:     schema.TypeList,
					},
					"client_id": {
						Description: "Client ID from provider. Example: `b1ba0bf54a4c2c0a1c29`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"client_secret": {
						Description: "Client secret from provider. Example: `bfa6gea4f129076761dcba8ce5e1e406bd83af7b`.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"cookie_samesite": {
				Description:  "Enum: `lax`, `strict`, `none`. Cookie SameSite attribute: `strict` prevents sending cookie for cross-site requests, effectively disabling direct linking from other sites to Grafana. `lax` is the default value.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"lax", "strict", "none"}, false),
			},
			"custom_domain": {
				Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name. Example: `grafana.example.org`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"dashboard_previews_enabled": {
				Description: "This feature is new in Grafana 9 and is quite resource intensive. It may cause low-end plans to work more slowly while the dashboard previews are rendering.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"dashboards_min_refresh_interval": {
				Description: "Signed sequence of decimal numbers, followed by a unit suffix (ms, s, m, h, d), e.g. 30s, 1h. Example: `5s`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"dashboards_versions_to_keep": {
				Description: "Dashboard versions to keep per dashboard. Example: `20`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"dataproxy_send_user_header": {
				Description: "Send `X-Grafana-User` header to data source.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"dataproxy_timeout": {
				Description: "Timeout for data proxy requests in seconds. Example: `30`.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"date_formats": {
				Description: "Grafana date format specifications",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"default_timezone": {
						Description: "Default time zone for user preferences. Value `browser` uses browser local time zone. Example: `Europe/Helsinki`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"full_date": {
						Description: "Moment.js style format string for cases where full date is shown. Example: `YYYY MM DD`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"interval_day": {
						Description: "Moment.js style format string used when a time requiring day accuracy is shown. Example: `MM/DD`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"interval_hour": {
						Description: "Moment.js style format string used when a time requiring hour accuracy is shown. Example: `MM/DD HH:mm`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"interval_minute": {
						Description: "Moment.js style format string used when a time requiring minute accuracy is shown. Example: `HH:mm`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"interval_month": {
						Description: "Moment.js style format string used when a time requiring month accuracy is shown. Example: `YYYY-MM`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"interval_second": {
						Description: "Moment.js style format string used when a time requiring second accuracy is shown. Example: `HH:mm:ss`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"interval_year": {
						Description: "Moment.js style format string used when a time requiring year accuracy is shown. Example: `YYYY`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"disable_gravatar": {
				Description: "Set to true to disable gravatar. Defaults to false (gravatar is enabled).",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"editors_can_admin": {
				Description: "Editors can manage folders, teams and dashboards created by them.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"external_image_storage": {
				Description: "External image store settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"access_key": {
						Description: "S3 access key. Requires permissions to the S3 bucket for the s3:PutObject and s3:PutObjectAcl actions. Example: `AAAAAAAAAAAAAAAAAAA`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"bucket_url": {
						Description: "Bucket URL for S3. Example: `https://grafana.s3-ap-southeast-2.amazonaws.com/`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"provider": {
						Description:  "Enum: `s3`. Provider type.",
						Required:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"s3"}, false),
					},
					"secret_key": {
						Description: "S3 secret key. Example: `AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`.",
						Required:    true,
						Type:        schema.TypeString,
					},
				}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"google_analytics_ua_id": {
				Description: "Google Analytics ID. Example: `UA-123456-4`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"ip_filter": {
				Deprecated:  "Deprecated. Use `ip_filter_string` instead.",
				Description: "Allow incoming connections from CIDR address block, e.g. `10.20.0.0/16`.",
				Elem: &schema.Schema{
					Description: "CIDR address block, either as a string, or in a dict with an optional description field. Example: `10.20.0.0/16`.",
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
						Description: "Description for IP filter list entry. Example: `Production service IP range`.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"network": {
						Description: "CIDR address block. Example: `10.20.0.0/16`.",
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
					Description: "CIDR address block, either as a string, or in a dict with an optional description field. Example: `10.20.0.0/16`.",
					Type:        schema.TypeString,
				},
				MaxItems: 1024,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"metrics_enabled": {
				Description: "Enable Grafana /metrics endpoint.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"oauth_allow_insecure_email_lookup": {
				Description: "Enforce user lookup based on email instead of the unique ID provided by the IdP.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"private_access": {
				Description: "Allow access to selected service ports from private networks",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"grafana": {
					Description: "Allow clients to connect to grafana with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"privatelink_access": {
				Description: "Allow access to selected service components through Privatelink",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"grafana": {
					Description: "Enable grafana.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"project_to_fork_from": {
				Description: "Name of another project to fork a service from. This has effect only when a new service is being created. Example: `anotherprojectname`.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"public_access": {
				Description: "Allow access to selected service ports from the public Internet",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"grafana": {
					Description: "Allow clients to connect to grafana from the public internet for service nodes that are in a project VPC or another type of private network.",
					Optional:    true,
					Type:        schema.TypeBool,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"recovery_basebackup_name": {
				Description: "Name of the basebackup to restore in forked service. Example: `backup-20191112t091354293891z`.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"service_log": {
				Description: "Store logs for the service so that they are available in the HTTP API and console.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"service_to_fork_from": {
				Description: "Name of another service to fork from. This has effect only when a new service is being created. Example: `anotherservicename`.",
				ForceNew:    true,
				Optional:    true,
				Type:        schema.TypeString,
			},
			"smtp_server": {
				Description: "SMTP server settings",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"from_address": {
						Description: "Address used for sending emails. Example: `yourgrafanauser@yourdomain.example.com`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"from_name": {
						Description: "Name used in outgoing emails, defaults to Grafana.",
						Optional:    true,
						Type:        schema.TypeString,
					},
					"host": {
						Description: "Server hostname or IP. Example: `smtp.example.com`.",
						Required:    true,
						Type:        schema.TypeString,
					},
					"password": {
						Description: "Password for SMTP authentication. Example: `ein0eemeev5eeth3Ahfu`.",
						Optional:    true,
						Sensitive:   true,
						Type:        schema.TypeString,
					},
					"port": {
						Description: "SMTP server port. Example: `25`.",
						Required:    true,
						Type:        schema.TypeInt,
					},
					"skip_verify": {
						Description: "Skip verifying server certificate. Defaults to false.",
						Optional:    true,
						Type:        schema.TypeBool,
					},
					"starttls_policy": {
						Description:  "Enum: `OpportunisticStartTLS`, `MandatoryStartTLS`, `NoStartTLS`. Either OpportunisticStartTLS, MandatoryStartTLS or NoStartTLS. Default is OpportunisticStartTLS.",
						Optional:     true,
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"OpportunisticStartTLS", "MandatoryStartTLS", "NoStartTLS"}, false),
					},
					"username": {
						Description: "Username for SMTP authentication. Example: `smtpuser`.",
						Optional:    true,
						Type:        schema.TypeString,
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
			"unified_alerting_enabled": {
				Description: "Enable or disable Grafana unified alerting functionality. By default this is enabled and any legacy alerts will be migrated on upgrade to Grafana 9+. To stay on legacy alerting, set unified_alerting_enabled to false and alerting_enabled to true. See https://grafana.com/docs/grafana/latest/alerting/set-up/migrating-alerts/ for more details.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"user_auto_assign_org": {
				Description: "Auto-assign new users on signup to main organization. Defaults to false.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"user_auto_assign_org_role": {
				Description:  "Enum: `Viewer`, `Admin`, `Editor`. Set role for new signups. Defaults to Viewer.",
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"Viewer", "Admin", "Editor"}, false),
			},
			"viewers_can_edit": {
				Description: "Users with view-only permission can edit but not save dashboards.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
