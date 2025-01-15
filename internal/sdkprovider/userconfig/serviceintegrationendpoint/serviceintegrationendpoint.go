// Code generated by user config generator. DO NOT EDIT.

package serviceintegrationendpoint

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func GetUserConfig(kind string) *schema.Schema {
	switch kind {
	case "autoscaler":
		return autoscalerUserConfig()
	case "datadog":
		return datadogUserConfig()
	case "external_aws_cloudwatch_logs":
		return externalAwsCloudwatchLogsUserConfig()
	case "external_aws_cloudwatch_metrics":
		return externalAwsCloudwatchMetricsUserConfig()
	case "external_aws_s3":
		return externalAwsS3UserConfig()
	case "external_azure_blob_storage":
		return externalAzureBlobStorageUserConfig()
	case "external_clickhouse":
		return externalClickhouseUserConfig()
	case "external_elasticsearch_logs":
		return externalElasticsearchLogsUserConfig()
	case "external_google_cloud_bigquery":
		return externalGoogleCloudBigqueryUserConfig()
	case "external_google_cloud_logging":
		return externalGoogleCloudLoggingUserConfig()
	case "external_kafka":
		return externalKafkaUserConfig()
	case "external_mysql":
		return externalMysqlUserConfig()
	case "external_opensearch_logs":
		return externalOpensearchLogsUserConfig()
	case "external_postgresql":
		return externalPostgresqlUserConfig()
	case "external_prometheus":
		return externalPrometheusUserConfig()
	case "external_schema_registry":
		return externalSchemaRegistryUserConfig()
	case "jolokia":
		return jolokiaUserConfig()
	case "prometheus":
		return prometheusUserConfig()
	case "rsyslog":
		return rsyslogUserConfig()
	default:
		return nil
	}
}

// GetFieldMapping returns TF fields to Json fields mapping (in unix-path way)
func GetFieldMapping(kind string) map[string]string {
	return map[string]map[string]string{}[kind]
}
func UserConfigTypes() []string {
	return []string{"autoscaler", "datadog", "external_aws_cloudwatch_logs", "external_aws_cloudwatch_metrics", "external_aws_s3", "external_azure_blob_storage", "external_clickhouse", "external_elasticsearch_logs", "external_google_cloud_bigquery", "external_google_cloud_logging", "external_kafka", "external_mysql", "external_opensearch_logs", "external_postgresql", "external_prometheus", "external_schema_registry", "jolokia", "prometheus", "rsyslog"}
}
