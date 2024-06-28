package flink

import (
	"context"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

// aivenFlinkApplicationSchema is the schema of the Flink Application resource.
var aivenFlinkApplicationSchema = map[string]*schema.Schema{
	"project": schemautil.CommonSchemaProjectReference,

	"service_name": schemautil.CommonSchemaServiceNameReference,

	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The name of the application.",
	},

	// Computed fields.
	"application_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Application ID.",
	},

	"created_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Application creation time.",
	},

	"created_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The user who created the application.",
	},

	"updated_at": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "When the application was updated.",
	},

	"updated_by": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The user who updated the application.",
	},
}

// ResourceFlinkApplication returns the Flink Application resource schema.
func ResourceFlinkApplication() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages an [Aiven for Apache Flink® application](https://aiven.io/docs/products/flink/concepts/flink-applications).",
		ReadContext:   resourceFlinkApplicationRead,
		CreateContext: resourceFlinkApplicationCreate,
		UpdateContext: resourceFlinkApplicationUpdate,
		DeleteContext: resourceFlinkApplicationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),
		Schema:   aivenFlinkApplicationSchema,
	}
}

// resourceFlinkApplicationRead is the read function for the Flink Application resource.
func resourceFlinkApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project, serviceName, ID, err := schemautil.SplitResourceID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.FlinkApplications.Get(ctx, project, serviceName, ID)
	if err != nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
	}

	if err := d.Set("project", project); err != nil {
		return diag.Errorf("error setting Flink Application `project` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("service_name", serviceName); err != nil {
		return diag.Errorf("error setting Flink Application `service_name` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("name", r.Name); err != nil {
		return diag.Errorf("error setting Flink Application `name` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("application_id", r.ID); err != nil {
		return diag.Errorf("error setting Flink Application `application_id` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("created_at", r.CreatedAt); err != nil {
		return diag.Errorf("error setting Flink Application `created_at` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("created_by", r.CreatedBy); err != nil {
		return diag.Errorf("error setting Flink Application `created_by` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("updated_at", r.UpdatedAt); err != nil {
		return diag.Errorf("error setting Flink Application `updated_at` for resource %s: %s", d.Id(), err)
	}

	if err := d.Set("updated_by", r.UpdatedBy); err != nil {
		return diag.Errorf("error setting Flink Application `updated_by` for resource %s: %s", d.Id(), err)
	}

	return nil
}

// resourceFlinkApplicationCreate is the create function for the Flink Application resource.
func resourceFlinkApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project := d.Get("project").(string)
	serviceName := d.Get("service_name").(string)

	r, err := client.FlinkApplications.Create(ctx, project, serviceName, aiven.CreateFlinkApplicationRequest{
		Name: d.Get("name").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(schemautil.BuildResourceID(project, serviceName, r.ID))

	return resourceFlinkApplicationRead(ctx, d, m)
}

func resourceFlinkApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project, serviceName, ID, err := schemautil.SplitResourceID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.FlinkApplications.Update(ctx, project, serviceName, ID, aiven.UpdateFlinkApplicationRequest{
		Name: d.Get("name").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceFlinkApplicationRead(ctx, d, m)
}

// resourceFlinkApplicationDelete is the delete function for the Flink Application resource.
func resourceFlinkApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	project, serviceName, ID, err := schemautil.SplitResourceID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.FlinkApplications.Delete(ctx, project, serviceName, ID)
	if common.IsCritical(err) {
		return diag.FromErr(err)
	}

	return nil
}
