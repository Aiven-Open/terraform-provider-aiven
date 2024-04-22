// Code generated by user config generator. DO NOT EDIT.

package service

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func GetUserConfig(kind string) *schema.Schema {
	switch kind {
	case "cassandra":
		return cassandraUserConfig()
	case "clickhouse":
		return clickhouseUserConfig()
	case "dragonfly":
		return dragonflyUserConfig()
	case "flink":
		return flinkUserConfig()
	case "grafana":
		return grafanaUserConfig()
	case "influxdb":
		return influxdbUserConfig()
	case "kafka":
		return kafkaUserConfig()
	case "kafka_connect":
		return kafkaConnectUserConfig()
	case "kafka_mirrormaker":
		return kafkaMirrormakerUserConfig()
	case "m3aggregator":
		return m3aggregatorUserConfig()
	case "m3db":
		return m3dbUserConfig()
	case "mysql":
		return mysqlUserConfig()
	case "opensearch":
		return opensearchUserConfig()
	case "pg":
		return pgUserConfig()
	case "redis":
		return redisUserConfig()
	case "thanos":
		return thanosUserConfig()
	default:
		return nil
	}
}
func UserConfigTypes() []string {
	return []string{"cassandra", "clickhouse", "dragonfly", "flink", "grafana", "influxdb", "kafka", "kafka_connect", "kafka_mirrormaker", "m3aggregator", "m3db", "mysql", "opensearch", "pg", "redis", "thanos"}
}
