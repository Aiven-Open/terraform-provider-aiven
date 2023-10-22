package account_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func TestAccAivenAccountTeamProject_basic(t *testing.T) {
	if _, ok := os.LookupEnv("AIVEN_ACCOUNT_NAME"); !ok {
		t.Skip("AIVEN_ACCOUNT_NAME env variable is required to run this test")
	}

	resourceName := "aiven_account_team_project.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenAccountTeamProjectResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountTeamProjectResource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenAccountTeamProejctAttributes("data.aiven_account_team_project.project"),
					resource.TestCheckResourceAttr(resourceName, "project_name", fmt.Sprintf("test-acc-pr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "team_type", "admin"),
				),
			},
		},
	})
}

func testAccAccountTeamProjectResource(name string) string {
	orgName := os.Getenv("AIVEN_ACCOUNT_NAME")

	return fmt.Sprintf(`
data "aiven_account" "foo" {
  name = "%[1]s"
}

resource "aiven_account_team" "foo" {
  account_id = data.aiven_account.foo.account_id
  name       = "test-acc-team-%[2]s"
}

resource "aiven_project" "foo" {
  project    = "test-acc-pr-%[2]s"
  account_id = aiven_account_team.foo.account_id
}

resource "aiven_account_team_project" "foo" {
  account_id   = data.aiven_account.foo.account_id
  team_id      = aiven_account_team.foo.team_id
  project_name = aiven_project.foo.project
  team_type    = "admin"
}

data "aiven_account_team_project" "project" {
  team_id      = aiven_account_team_project.foo.team_id
  account_id   = aiven_account_team_project.foo.account_id
  project_name = aiven_account_team_project.foo.project_name

  depends_on = [aiven_account_team_project.foo]
}`, orgName, name)
}

func testAccCheckAivenAccountTeamProjectResourceDestroy(s *terraform.State) error {
	c := acc.GetTestAivenClient()

	ctx := context.Background()

	// loop through the resources in state, verifying each account team project is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_account_team_project" {
			continue
		}

		accountID, teamID, projectName, err := schemautil.SplitResourceID3(rs.Primary.ID)
		if err != nil {
			return err
		}

		r, err := c.Accounts.List(ctx)
		if err != nil {
			if err.(aiven.Error).Status != 404 {
				return err
			}

			return nil
		}

		for _, a := range r.Accounts {
			if a.Id == accountID {
				rp, err := c.AccountTeamProjects.List(ctx, accountID, teamID)
				if err != nil {
					if err.(aiven.Error).Status != 404 {
						return err
					}

					return nil
				}

				for _, p := range rp.Projects {
					if p.ProjectName == projectName {
						return fmt.Errorf("account team project (%s) still exists", rs.Primary.ID)
					}
				}
			}
		}
	}

	return nil
}

func testAccCheckAivenAccountTeamProejctAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		log.Printf("[DEBUG] account team project attributes %v", a)

		if a["account_id"] == "" {
			return fmt.Errorf("expected to get an account id from Aiven")
		}

		if a["team_id"] == "" {
			return fmt.Errorf("expected to get a team_id from Aiven")
		}

		if a["project_name"] == "" {
			return fmt.Errorf("expected to get a project_name from Aiven")
		}

		if a["team_type"] == "" {
			return fmt.Errorf("expected to get a team_type from Aiven")
		}

		return nil
	}
}
