resource "aiven_clickhouse" "clickhouse" {
  project                 = data.aiven_project.pr1.project
  cloud_name              = "google-europe-west1"
  plan                    = "business-4"
  service_name            = "my-clickhouse"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}
