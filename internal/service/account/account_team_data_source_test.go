package account_test

import (
	"os"
	"testing"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAivenAccountTeamDataSource_basic(t *testing.T) {
	if _, ok := os.LookupEnv("AIVEN_ACCOUNT_NAME"); !ok {
		t.Skip("AIVEN_ACCOUNT_NAME env variable is required to run this test")
	}

	datasourceName := "data.aiven_account_team.team"
	resourceName := "aiven_account_team.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acc.TestAccPreCheck(t) },
		ProviderFactories: acc.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountTeamResource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(datasourceName, "team_id", resourceName, "team_id"),
					resource.TestCheckResourceAttrPair(datasourceName, "create_time", resourceName, "create_time"),
					resource.TestCheckResourceAttrPair(datasourceName, "update_time", resourceName, "update_time"),
				),
			},
		},
	})
}
