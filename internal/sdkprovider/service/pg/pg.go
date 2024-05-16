package pg

import (
	"context"
	"log"
	"time"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig/stateupgrader"
)

func aivenPGSchema() map[string]*schema.Schema {
	s := schemautil.ServiceCommonSchemaWithUserConfig(schemautil.ServiceTypePG)
	s[schemautil.ServiceTypePG] = &schema.Schema{
		Type:        schema.TypeList,
		MaxItems:    1,
		Computed:    true,
		Description: "PostgreSQL specific server provided values",
		Optional:    true,
		Sensitive:   true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// TODO: Remove `uri` in the next major version.
				"uri": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "PostgreSQL master connection URI",
					Optional:    true,
					Sensitive:   true,
				},
				"uris": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "PostgreSQL master connection URIs",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"bouncer": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Bouncer connection details",
				},
				// TODO: Remove `host` in the next major version.
				"host": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "PostgreSQL master node host IP or name",
				},
				// TODO: Remove `port` in the next major version.
				"port": {
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "PostgreSQL port",
				},
				// TODO: Remove `sslmode` in the next major version.
				"sslmode": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "PostgreSQL sslmode setting (currently always \"require\")",
				},
				// TODO: Remove `user` in the next major version.
				"user": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "PostgreSQL admin user name",
				},
				// TODO: Remove `password` in the next major version.
				"password": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "PostgreSQL admin user password",
					Sensitive:   true,
				},
				// TODO: Remove `dbname` in the next major version.
				"dbname": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "Primary PostgreSQL database name",
				},
				"params": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "PostgreSQL connection parameters",
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"host": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "PostgreSQL host IP or name",
							},
							"port": {
								Type:        schema.TypeInt,
								Computed:    true,
								Description: "PostgreSQL port",
							},
							"sslmode": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "PostgreSQL sslmode setting (currently always \"require\")",
							},
							"user": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "PostgreSQL admin user name",
							},
							"password": {
								Type:        schema.TypeString,
								Computed:    true,
								Sensitive:   true,
								Description: "PostgreSQL admin user password",
							},
							"database_name": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Primary PostgreSQL database name",
							},
						},
					},
				},
				"replica_uri": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "PostgreSQL replica URI for services with a replica",
					Sensitive:   true,
				},
				"standby_uris": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "PostgreSQL standby connection URIs",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"syncing_uris": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "PostgreSQL syncing connection URIs",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				// TODO: This isn't in the connection info, but it's in the metadata.
				//  We should move this to the other part of the schema in the next major version.
				"max_connections": {
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Connection limit",
				},
			},
		},
	}
	return s
}

func ResourcePG() *schema.Resource {
	return &schema.Resource{
		Description:   "The PG resource allows the creation and management of Aiven PostgreSQL services.",
		CreateContext: schemautil.ResourceServiceCreateWrapper(schemautil.ServiceTypePG),
		ReadContext:   schemautil.ResourceServiceRead,
		UpdateContext: resourceServicePGUpdate,
		DeleteContext: schemautil.ResourceServiceDelete,
		CustomizeDiff: customdiff.Sequence(
			schemautil.SetServiceTypeIfEmpty(schemautil.ServiceTypePG),
			schemautil.CustomizeDiffDisallowMultipleManyToOneKeys,
			customdiff.IfValueChange("tag",
				schemautil.TagsShouldNotBeEmpty,
				schemautil.CustomizeDiffCheckUniqueTag,
			),
			customdiff.IfValueChange("disk_space",
				schemautil.DiskSpaceShouldNotBeEmpty,
				schemautil.CustomizeDiffCheckDiskSpace,
			),
			customdiff.IfValueChange("additional_disk_space",
				schemautil.DiskSpaceShouldNotBeEmpty,
				schemautil.CustomizeDiffCheckDiskSpace,
			),
			customdiff.IfValueChange("service_integrations",
				schemautil.ServiceIntegrationShouldNotBeEmpty,
				schemautil.CustomizeDiffServiceIntegrationAfterCreation,
			),
			customdiff.Sequence(
				schemautil.CustomizeDiffCheckStaticIPDisassociation,
				schemautil.CustomizeDiffCheckPlanAndStaticIpsCannotBeModifiedTogether,
			),
		),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema:         aivenPGSchema(),
		SchemaVersion:  1,
		StateUpgraders: stateupgrader.PG(),
	}
}

func resourceServicePGUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	projectName, serviceName, err := schemautil.SplitResourceID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	userConfig, err := schemautil.ExpandService(schemautil.ServiceTypePG, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if userConfig["pg_version"] != nil {
		s, err := client.Services.Get(ctx, projectName, serviceName)
		if err != nil {
			return diag.Errorf("cannot get a common: %s", err)
		}

		if userConfig["pg_version"].(string) != s.UserConfig["pg_version"].(string) {
			t, err := client.ServiceTask.Create(ctx, projectName, serviceName, aiven.ServiceTaskRequest{
				TargetVersion: userConfig["pg_version"].(string),
				TaskType:      "upgrade_check",
			})
			if err != nil {
				return diag.Errorf("cannot create PG upgrade check task: %s", err)
			}

			w := &ServiceTaskWaiter{
				Context:     ctx,
				Client:      m.(*aiven.Client),
				Project:     projectName,
				ServiceName: serviceName,
				TaskID:      t.Task.Id,
			}

			taskI, err := w.Conf(d.Timeout(schema.TimeoutUpdate)).WaitForStateContext(ctx)
			if err != nil {
				return diag.Errorf("error waiting for Aiven service task to be DONE: %s", err)
			}

			task := taskI.(*aiven.ServiceTaskResponse)
			if !*task.Task.Success {
				return diag.Errorf(
					"PG service upgrade check error, version upgrade from %s to %s, result: %s",
					task.Task.SourcePgVersion, task.Task.TargetPgVersion, task.Task.Result)
			}

			log.Printf("[DEBUG] PG service upgrade check result: %s", task.Task.Result)
		}
	}

	return schemautil.ResourceServiceUpdate(ctx, d, m)
}

// ServiceTaskWaiter is used to refresh the Aiven Service Task endpoints when
// provisioning.
type ServiceTaskWaiter struct {
	Context     context.Context
	Client      *aiven.Client
	Project     string
	ServiceName string
	TaskID      string
}

// RefreshFunc will call the Aiven client and refresh its state.
func (w *ServiceTaskWaiter) RefreshFunc() retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		t, err := w.Client.ServiceTask.Get(
			w.Context,
			w.Project,
			w.ServiceName,
			w.TaskID,
		)
		if err != nil {
			return nil, "", err
		}

		if t.Task.Success == nil {
			return nil, "IN_PROGRESS", nil
		}

		return t, "DONE", nil
	}
}

// Conf sets up the configuration to refresh.
func (w *ServiceTaskWaiter) Conf(timeout time.Duration) *retry.StateChangeConf {
	return &retry.StateChangeConf{
		Pending:                   []string{"IN_PROGRESS"},
		Target:                    []string{"DONE"},
		Refresh:                   w.RefreshFunc(),
		Delay:                     common.DefaultStateChangeDelay,
		Timeout:                   timeout,
		MinTimeout:                common.DefaultStateChangeMinTimeout,
		ContinuousTargetOccurence: 3,
	}
}
