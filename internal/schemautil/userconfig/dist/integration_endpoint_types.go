// Code generated by internal/schemautil/userconfig/userconfig_test.go; DO NOT EDIT.

package dist

import (
	schemautil "github.com/aiven/terraform-provider-aiven/internal/schemautil"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IntegrationEndpointTypeDatadog is a generated function returning the schema of the datadog IntegrationEndpointType.
func IntegrationEndpointTypeDatadog() *schema.Schema {
	s := map[string]*schema.Schema{
		"datadog_api_key": {
			Description: "Datadog API key",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"datadog_tags": {
			Description: "Custom tags provided by user",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"comment": {
					Description: "Optional tag explanation",
					Optional:    true,
					Type:        schema.TypeString,
				},
				"tag": {
					Description: "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix 'aiven-' are reserved for Aiven.",
					Optional:    true,
					Type:        schema.TypeString,
				},
			}},
			MaxItems: 32,
			Optional: true,
			Type:     schema.TypeList,
		},
		"disable_consumer_stats": {
			Description: "Disable consumer group metrics",
			Optional:    true,
			Type:        schema.TypeBool,
		},
		"kafka_consumer_check_instances": {
			Description: "Number of separate instances to fetch kafka consumer statistics with",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"kafka_consumer_stats_timeout": {
			Description: "Number of seconds that datadog will wait to get consumer statistics from brokers",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"max_partition_contexts": {
			Description: "Maximum number of partition contexts to send",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"site": {
			Description: "Datadog intake site. Defaults to datadoghq.com",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "Datadog user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalAwsCloudwatchLogs is a generated function returning the schema of the external_aws_cloudwatch_logs IntegrationEndpointType.
func IntegrationEndpointTypeExternalAwsCloudwatchLogs() *schema.Schema {
	s := map[string]*schema.Schema{
		"access_key": {
			Description: "AWS access key. Required permissions are logs:CreateLogGroup, logs:CreateLogStream, logs:PutLogEvents and logs:DescribeLogStreams",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"log_group_name": {
			Description: "AWS CloudWatch log group name",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"region": {
			Description: "AWS region",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"secret_key": {
			Description: "AWS secret key",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalAwsCloudwatchLogs user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalAwsCloudwatchMetrics is a generated function returning the schema of the external_aws_cloudwatch_metrics IntegrationEndpointType.
func IntegrationEndpointTypeExternalAwsCloudwatchMetrics() *schema.Schema {
	s := map[string]*schema.Schema{
		"access_key": {
			Description: "AWS access key. Required permissions are cloudwatch:PutMetricData",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"namespace": {
			Description: "AWS CloudWatch Metrics Namespace",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"region": {
			Description: "AWS region",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"secret_key": {
			Description: "AWS secret key",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalAwsCloudwatchMetrics user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalElasticsearchLogs is a generated function returning the schema of the external_elasticsearch_logs IntegrationEndpointType.
func IntegrationEndpointTypeExternalElasticsearchLogs() *schema.Schema {
	s := map[string]*schema.Schema{
		"ca": {
			Description: "PEM encoded CA certificate",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"index_days_max": {
			Default:     "3",
			Description: "Maximum number of days of logs to keep The default value is `3`.",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"index_prefix": {
			Default:     "logs",
			Description: "Elasticsearch index prefix The default value is `logs`.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"timeout": {
			Default:     "10.0",
			Description: "Elasticsearch request timeout limit The default value is `10.0`.",
			Optional:    true,
			Type:        schema.TypeFloat,
		},
		"url": {
			Description: "Elasticsearch connection URL",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalElasticsearchLogs user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalGoogleCloudLogging is a generated function returning the schema of the external_google_cloud_logging IntegrationEndpointType.
func IntegrationEndpointTypeExternalGoogleCloudLogging() *schema.Schema {
	s := map[string]*schema.Schema{
		"log_id": {
			Description: "Google Cloud Logging log id",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"project_id": {
			Description: "GCP project id.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"service_account_credentials": {
			Description: "This is a JSON object with the fields documented in https://cloud.google.com/iam/docs/creating-managing-service-account-keys .",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalGoogleCloudLogging user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalKafka is a generated function returning the schema of the external_kafka IntegrationEndpointType.
func IntegrationEndpointTypeExternalKafka() *schema.Schema {
	s := map[string]*schema.Schema{
		"bootstrap_servers": {
			Description: "Bootstrap servers",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"sasl_mechanism": {
			Description: "The list of SASL mechanisms enabled in the Kafka server.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"sasl_plain_password": {
			Description: "Password for SASL PLAIN mechanism in the Kafka server.",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"sasl_plain_username": {
			Description: "Username for SASL PLAIN mechanism in the Kafka server.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"security_protocol": {
			Description: "Security protocol",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"ssl_ca_cert": {
			Description: "PEM-encoded CA certificate",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"ssl_client_cert": {
			Description: "PEM-encoded client certificate",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"ssl_client_key": {
			Description: "PEM-encoded client key",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"ssl_endpoint_identification_algorithm": {
			Description: "The endpoint identification algorithm to validate server hostname using server certificate.",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalKafka user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalOpensearchLogs is a generated function returning the schema of the external_opensearch_logs IntegrationEndpointType.
func IntegrationEndpointTypeExternalOpensearchLogs() *schema.Schema {
	s := map[string]*schema.Schema{
		"ca": {
			Description: "PEM encoded CA certificate",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"index_days_max": {
			Default:     "3",
			Description: "Maximum number of days of logs to keep The default value is `3`.",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"index_prefix": {
			Default:     "logs",
			Description: "OpenSearch index prefix The default value is `logs`.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"timeout": {
			Default:     "10.0",
			Description: "OpenSearch request timeout limit The default value is `10.0`.",
			Optional:    true,
			Type:        schema.TypeFloat,
		},
		"url": {
			Description: "OpenSearch connection URL",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalOpensearchLogs user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalPostgresql is a generated function returning the schema of the external_postgresql IntegrationEndpointType.
func IntegrationEndpointTypeExternalPostgresql() *schema.Schema {
	s := map[string]*schema.Schema{
		"host": {
			Description: "Hostname or IP address of the server",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"password": {
			Description: "Password",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"port": {
			Description: "Port number of the server",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"ssl_mode": {
			Default:     "verify-full",
			Description: "SSL Mode The default value is `verify-full`.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"ssl_root_cert": {
			Default:     "",
			Description: "SSL Root Cert",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"username": {
			Description: "User name",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalPostgresql user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeExternalSchemaRegistry is a generated function returning the schema of the external_schema_registry IntegrationEndpointType.
func IntegrationEndpointTypeExternalSchemaRegistry() *schema.Schema {
	s := map[string]*schema.Schema{
		"authentication": {
			Description: "Authentication method",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"basic_auth_password": {
			Description: "Basic authentication password",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"basic_auth_username": {
			Description: "Basic authentication user name",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"url": {
			Description: "Schema Registry URL",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "ExternalSchemaRegistry user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeJolokia is a generated function returning the schema of the jolokia IntegrationEndpointType.
func IntegrationEndpointTypeJolokia() *schema.Schema {
	s := map[string]*schema.Schema{
		"basic_auth_password": {
			Description: "Jolokia basic authentication password",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"basic_auth_username": {
			Description: "Jolokia basic authentication username",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "Jolokia user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypePrometheus is a generated function returning the schema of the prometheus IntegrationEndpointType.
func IntegrationEndpointTypePrometheus() *schema.Schema {
	s := map[string]*schema.Schema{
		"basic_auth_password": {
			Description: "Prometheus basic authentication password",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"basic_auth_username": {
			Description: "Prometheus basic authentication username",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "Prometheus user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeRsyslog is a generated function returning the schema of the rsyslog IntegrationEndpointType.
func IntegrationEndpointTypeRsyslog() *schema.Schema {
	s := map[string]*schema.Schema{
		"ca": {
			Description: "PEM encoded CA certificate",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"cert": {
			Description: "PEM encoded client certificate",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"format": {
			Default:     "rfc5424",
			Description: "message format The default value is `rfc5424`.",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"key": {
			Description: "PEM encoded client key",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"logline": {
			Description: "custom syslog message format",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"port": {
			Default:     "514",
			Description: "rsyslog server port The default value is `514`.",
			Optional:    true,
			Type:        schema.TypeInt,
		},
		"sd": {
			Description: "Structured data block for log message",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"server": {
			Description: "rsyslog server IP address or hostname",
			Optional:    true,
			Type:        schema.TypeString,
		},
		"tls": {
			Default:     true,
			Description: "Require TLS The default value is `true`.",
			Optional:    true,
			Type:        schema.TypeBool,
		},
	}

	return &schema.Schema{
		Description:      "Rsyslog user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}

// IntegrationEndpointTypeSignalfx is a generated function returning the schema of the signalfx IntegrationEndpointType.
func IntegrationEndpointTypeSignalfx() *schema.Schema {
	s := map[string]*schema.Schema{
		"enabled_metrics": {
			Description: "list of metrics to send",
			Elem:        &schema.Schema{Type: schema.TypeString},
			MaxItems:    256,
			Optional:    true,
			Type:        schema.TypeList,
		},
		"signalfx_api_key": {
			Description: "SignalFX API key",
			Optional:    true,
			Sensitive:   true,
			Type:        schema.TypeString,
		},
		"signalfx_realm": {
			Default:     "us0",
			Description: "SignalFX realm The default value is `us0`.",
			Optional:    true,
			Type:        schema.TypeString,
		},
	}

	return &schema.Schema{
		Description:      "Signalfx user configurable settings",
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFuncSkipArrays(s),
		Elem:             &schema.Resource{Schema: s},
		MaxItems:         1,
		Optional:         true,
		Type:             schema.TypeList,
	}
}
