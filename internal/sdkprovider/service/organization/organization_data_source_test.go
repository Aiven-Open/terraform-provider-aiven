package organization_test

import (
	"testing"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAivenOrganizationDataSource_basic(t *testing.T) {
	datasourceName := "data.aiven_organization.organization"
	resourceName := "aiven_organization.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acc.TestAccPreCheck(t) },
		ProviderFactories: acc.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccOrganizationResource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(datasourceName, "tenant_id", resourceName, "tenant_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "create_time", resourceName, "create_time"),
					resource.TestCheckResourceAttrPair(datasourceName, "update_time", resourceName, "update_time"),
				),
			},
		},
	})
}
