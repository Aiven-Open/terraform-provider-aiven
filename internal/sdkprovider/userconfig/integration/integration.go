// Code generated by user config generator. DO NOT EDIT.

package integration

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func GetUserConfig(kind string) *schema.Schema {
	switch kind {
	case "clickhouse_kafka":
		return clickhouseKafkaUserConfig()
	case "clickhouse_postgresql":
		return clickhousePostgresqlUserConfig()
	case "datadog":
		return datadogUserConfig()
	case "external_aws_cloudwatch_logs":
		return externalAwsCloudwatchLogsUserConfig()
	case "external_aws_cloudwatch_metrics":
		return externalAwsCloudwatchMetricsUserConfig()
	case "external_elasticsearch_logs":
		return externalElasticsearchLogsUserConfig()
	case "external_opensearch_logs":
		return externalOpensearchLogsUserConfig()
	case "kafka_connect":
		return kafkaConnectUserConfig()
	case "kafka_logs":
		return kafkaLogsUserConfig()
	case "kafka_mirrormaker":
		return kafkaMirrormakerUserConfig()
	case "logs":
		return logsUserConfig()
	case "metrics":
		return metricsUserConfig()
	case "prometheus":
		return prometheusUserConfig()
	default:
		panic("unknown user config type: " + kind)
	}
}
