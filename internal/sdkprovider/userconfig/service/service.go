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

// GetFieldMapping returns TF fields to Json fields mapping (in unix-path way)
func GetFieldMapping(kind string) map[string]string {
	return map[string]map[string]string{
		"cassandra": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"clickhouse": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"dragonfly": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"flink": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"grafana": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"influxdb": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"kafka": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"kafka_connect": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"kafka_mirrormaker": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"m3aggregator": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"m3db": {
			"ip_filter_object":                "ip_filter",
			"ip_filter_string":                "ip_filter",
			"rules/mapping/namespaces_object": "rules/mapping/namespaces",
			"rules/mapping/namespaces_string": "rules/mapping/namespaces",
		},
		"mysql": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"opensearch": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"pg": {
			"ip_filter_object":                                "ip_filter",
			"ip_filter_string":                                "ip_filter",
			"pg/pg_partman_bgw__dot__interval":                "pg/pg_partman_bgw.interval",
			"pg/pg_partman_bgw__dot__role":                    "pg/pg_partman_bgw.role",
			"pg/pg_stat_monitor__dot__pgsm_enable_query_plan": "pg/pg_stat_monitor.pgsm_enable_query_plan",
			"pg/pg_stat_monitor__dot__pgsm_max_buckets":       "pg/pg_stat_monitor.pgsm_max_buckets",
			"pg/pg_stat_statements__dot__track":               "pg/pg_stat_statements.track",
		},
		"redis": {
			"ip_filter_object": "ip_filter",
			"ip_filter_string": "ip_filter",
		},
		"thanos": {
			"compactor/retention_days":                         "compactor/retention.days",
			"ip_filter_object":                                 "ip_filter",
			"ip_filter_string":                                 "ip_filter",
			"query/query_default_evaluation_interval":          "query/query.default-evaluation-interval",
			"query/query_lookback_delta":                       "query/query.lookback-delta",
			"query/query_metadata_default_time_range":          "query/query.metadata.default-time-range",
			"query/query_timeout":                              "query/query.timeout",
			"query_frontend":                                   "query-frontend",
			"query_frontend/query_range_align_range_with_step": "query-frontend/query-range.align-range-with-step",
			"receiver_ingesting":                               "receiver-ingesting",
			"receiver_routing":                                 "receiver-routing",
		},
	}[kind]
}
func UserConfigTypes() []string {
	return []string{"cassandra", "clickhouse", "dragonfly", "flink", "grafana", "influxdb", "kafka", "kafka_connect", "kafka_mirrormaker", "m3aggregator", "m3db", "mysql", "opensearch", "pg", "redis", "thanos"}
}
