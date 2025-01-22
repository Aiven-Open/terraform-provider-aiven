package vpc_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/aiven/go-client-codegen/handler/organizationvpc"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

const (
	organizationVPCResource   = "aiven_organization_vpc"
	organizationVPCDataSource = "data_aiven_organization_vpc"
)

func TestAccAivenOrganizationVPC(t *testing.T) {
	var (
		orgName        = acc.SkipIfEnvVarsNotSet(t, "AIVEN_ORGANIZATION_NAME")["AIVEN_ORGANIZATION_NAME"]
		registry       = preSetOrgVPCTemplates(t)
		newComposition = func() *acc.CompositionBuilder {
			return registry.NewCompositionBuilder().
				Add("organization_data", map[string]interface{}{
					"organization_name": orgName})
		}

		resourceName   = fmt.Sprintf("%s.%s", organizationVPCResource, "test_org_vpc")
		dataSourceName = fmt.Sprintf("data.%s.%s", organizationVPCResource, "vpc_ds")
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOrganizationVPCResourceDestroy,
		Steps: []resource.TestStep{
			{
				// invalid CIDR range
				Config: `
				resource "aiven_organization_vpc" "test_validation" {
					organization_id = "test-org-id"
					cloud_name     = "aws-eu-west-1"
					network_cidr   = "256.256.256.256/24"
				}`,
				ExpectError: regexp.MustCompile(`expected "network_cidr" to be a valid CIDR Value`),
			},
			{
				// basic VPC creation
				Config: newComposition().
					Add(organizationVPCResource, map[string]interface{}{
						"resource_name": "test_org_vpc",
						"cloud_name":    "aws-eu-west-1",
						"network_cidr":  "10.0.0.0/24",
					}).MustRender(t),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cloud_name", "aws-eu-west-1"),
					resource.TestCheckResourceAttr(resourceName, "network_cidr", "10.0.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "state", string(organizationvpc.VpcStateTypeActive)),
					resource.TestCheckResourceAttrSet(resourceName, "organization_id"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
				),
			},
			{
				// test ForceNew on network_cidr change
				Config: newComposition().
					Add(organizationVPCResource, map[string]interface{}{
						"resource_name": "test_org_vpc",
						"cloud_name":    "aws-eu-west-1",
						"network_cidr":  "10.1.0.0/24", // Changed CIDR
					}).MustRender(t),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cloud_name", "aws-eu-west-1"),
					resource.TestCheckResourceAttr(resourceName, "network_cidr", "10.1.0.0/24"),
					resource.TestCheckResourceAttr(resourceName, "state", string(organizationvpc.VpcStateTypeActive)),
					resource.TestCheckResourceAttrSet(resourceName, "organization_id"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
				),
			},
			{
				// import state
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// test the data source
				Config: newComposition().
					Add(organizationVPCResource, map[string]interface{}{
						"resource_name": "test_org_vpc",
						"cloud_name":    "aws-eu-west-1",
						"network_cidr":  "10.0.0.0/24",
					}).
					Add(organizationVPCDataSource, map[string]interface{}{
						"resource_name": "vpc_ds",
					}).MustRender(t),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "organization_id", resourceName, "organization_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cloud_name", resourceName, "cloud_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "network_cidr", resourceName, "network_cidr"),
					resource.TestCheckResourceAttrPair(dataSourceName, "state", resourceName, "state"),
					resource.TestCheckResourceAttrPair(dataSourceName, "create_time", resourceName, "create_time"),
					resource.TestCheckResourceAttr(dataSourceName, "state", string(organizationvpc.VpcStateTypeActive)),
					resource.TestCheckResourceAttrSet(dataSourceName, "organization_vpc_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "organization_vpc_id", resourceName, "organization_vpc_id"),
				),
			},
		},
	})
}

func testAccCheckOrganizationVPCResourceDestroy(s *terraform.State) error {
	ctx := context.Background()

	c, err := acc.GetTestGenAivenClient()
	if err != nil {
		return fmt.Errorf("error initializing Aiven client: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != organizationVPCResource {
			continue
		}

		orgID, vpcID, err := schemautil.SplitResourceID2(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error splitting resource ID: %w", err)
		}

		orgVPC, err := c.OrganizationVpcGet(ctx, orgID, vpcID)
		if common.IsCritical(err) {
			return fmt.Errorf("error fetching VPC (%q): %w", vpcID, err)
		}

		if orgVPC != nil {
			return fmt.Errorf("VPC (%q) for organization (%q) still exists", vpcID, orgID)
		}
	}

	return nil
}

func preSetOrgVPCTemplates(t *testing.T) *acc.TemplateRegistry {
	t.Helper()

	var registry = acc.NewTemplateRegistry(organizationVPCResource)

	registry.MustAddTemplate(t, "organization_data", `
data "aiven_organization" "foo" {
	name = "{{ .organization_name }}"
}`)

	registry.MustAddTemplate(t, organizationVPCResource, `
resource "aiven_organization_vpc" "{{ .resource_name }}" {
    organization_id = data.aiven_organization.foo.id
    cloud_name     = "{{ .cloud_name }}"
    network_cidr   = "{{ .network_cidr }}"
}`)

	registry.MustAddTemplate(t, organizationVPCDataSource, `
data "aiven_organization_vpc" "{{ .resource_name }}" {
    organization_id = data.aiven_organization.foo.id
    organization_vpc_id          = split("/", aiven_organization_vpc.test_org_vpc.id)[1]
}`)

	return registry
}
