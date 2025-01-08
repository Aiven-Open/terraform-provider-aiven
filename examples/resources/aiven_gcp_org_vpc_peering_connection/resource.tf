resource "aiven_gcp_org_vpc_peering_connection" "example" {
  vpc_id = aiven_organization_vpc.example.id  # Format: "organization_id/vpc_id"
  gcp_project_id = "my-gcp-project-123"       # Your GCP project ID
  peer_vpc = "my-vpc-network"                 # Your GCP VPC network name
}