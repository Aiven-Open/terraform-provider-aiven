package vpc

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/aiven/go-client-codegen/handler/vpc"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

func DatasourceProjectVPC() *schema.Resource {
	aivenProjectVPCDataSourceSchema := map[string]*schema.Schema{
		"project": {
			Type:          schema.TypeString,
			ValidateFunc:  validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]*$"), "project name should be alphanumeric"),
			Description:   "Identifies the project this resource belongs to.",
			Optional:      true,
			ConflictsWith: []string{"vpc_id"},
		},
		"cloud_name": {
			Type:          schema.TypeString,
			Description:   "The cloud provider and region where the service is hosted in the format `CLOUD_PROVIDER-REGION_NAME`. For example, `google-europe-west1` or `aws-us-east-2`.",
			Optional:      true,
			ConflictsWith: []string{"vpc_id"},
		},
		"vpc_id": {
			Type:          schema.TypeString,
			Description:   "The ID of the VPC. This can be used to filter out the other VPCs if there are more than one for the project and cloud.",
			Optional:      true,
			ConflictsWith: []string{"project", "cloud_name"},
			ValidateDiagFunc: func(i interface{}, _ cty.Path) diag.Diagnostics {
				_, err := schemautil.SplitResourceID(i.(string), 2)
				if err != nil {
					return diag.Errorf("invalid vpc_id, should have the following format {project_name}/{project_vpc_id}: %s", err)
				}
				return nil
			},
		},
		"network_cidr": {
			Computed:    true,
			Type:        schema.TypeString,
			Description: "Network address range used by the VPC. For example, `192.168.0.0/24`.",
		},
		"state": {
			Computed:    true,
			Type:        schema.TypeString,
			Description: userconfig.Desc("State of the VPC.").PossibleValuesString(vpc.VpcStateTypeChoices()...).Build(),
		},
	}

	return &schema.Resource{
		ReadContext: datasourceProjectVPCRead,
		Description: "Gets information about the VPC for an Aiven project.",
		Schema:      aivenProjectVPCDataSourceSchema,
	}
}

func datasourceProjectVPCRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	var vpcID, projectName, cloudName string

	// This two branches are isolated by tf validation
	if s, hasID := d.GetOk("vpc_id"); hasID {
		chunks, err := schemautil.SplitResourceID(s.(string), 2)
		if err != nil {
			return diag.Errorf("error splitting vpc_id: %s:", err)
		}
		projectName = chunks[0]
		vpcID = chunks[1]
	} else {
		projectName = d.Get("project").(string)
		cloudName = d.Get("cloud_name").(string)
	}

	vpcList, err := client.VPCs.List(ctx, projectName)
	if err != nil {
		return diag.Errorf("error getting a list of project %q VPCs: %s", projectName, err)
	}

	// At this point we have strictly either vpcID OR cloudName
	// Because of ConflictsWith: []string{"project", "cloud_name"},
	vpc, err := getVPC(vpcList, vpcID, cloudName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(schemautil.BuildResourceID(projectName, vpc.ProjectVPCID))
	err = copyVPCPropertiesFromAPIResponseToTerraform(d, vpc, projectName)
	if err != nil {
		return diag.Errorf("error setting project VPC datasource values: %s", err)
	}

	return nil
}

// getVPC gets VPC by id or cloud name
func getVPC(vpcList []*aiven.VPC, vpcID, cloudName string) (vpc *aiven.VPC, err error) {
	//   A  xnor  B   | A | B | Out
	// ---------------|---|---|----
	// "foo" == ""    | 0 | 1 | 0
	//    "" == "foo" | 1 | 0 | 0
	//    "" == ""    | 1 | 1 | 1
	// "foo" == "foo" | 0 | 0 | 1
	if (vpcID == "") == (cloudName == "") {
		return nil, fmt.Errorf("provide exactly one: vpc_id or cloud_name")
	}

	for _, v := range vpcList {
		// Exact match
		if v.ProjectVPCID == vpcID {
			return v, nil
		}

		// cloudName can't be empty by this time
		if v.CloudName != cloudName {
			continue
		}

		// Cases:
		// 1. multiple active with same cloudName
		// 2. one is deleting and another one is creating (APPROVED)
		if vpc != nil {
			return nil, fmt.Errorf("multiple project VPC with cloud_name %q, use vpc_id instead", cloudName)
		}
		vpc = v
	}

	if vpc == nil {
		err = fmt.Errorf("not found project VPC with cloud_name %q", cloudName)
	}

	return vpc, err
}
