package kafka

import (
	"context"

	avngen "github.com/aiven/go-client-codegen"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func DatasourceMirrorMakerReplicationFlowTopic() *schema.Resource {
	return &schema.Resource{
		ReadContext: common.WithGenClient(datasourceMirrorMakerReplicationFlowRead),
		Description: "Gets information about an [Aiven for Apache Kafka® MirrorMaker 2](https://aiven.io/docs/products/kafka/kafka-mirrormaker) replication flow.",
		Schema: schemautil.ResourceSchemaAsDatasourceSchema(
			aivenMirrorMakerReplicationFlowSchema,
			"project", "service_name", "source_cluster", "target_cluster"),
	}
}

func datasourceMirrorMakerReplicationFlowRead(ctx context.Context, d *schema.ResourceData, client avngen.Client) error {
	projectName := d.Get("project").(string)
	serviceName := d.Get("service_name").(string)
	sourceCluster := d.Get("source_cluster").(string)
	targetCluster := d.Get("target_cluster").(string)

	d.SetId(schemautil.BuildResourceID(projectName, serviceName, sourceCluster, targetCluster))

	return resourceMirrorMakerReplicationFlowRead(ctx, d, client)
}
