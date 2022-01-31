// Copyright (c) 2018-2022 Aiven, Helsinki, Finland. https://aiven.io/
package aiven

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/terraform-provider-aiven/aiven/internal/schemautil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAivenClickhouseUser_basic(t *testing.T) {
	t.Parallel()

	t.Run("clickhouse user creation", func(tt *testing.T) {
		resourceName := "aiven_clickhouse_user.foo"
		rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

		resource.ParallelTest(tt, resource.TestCase{
			PreCheck:          func() { testAccPreCheck(tt) },
			ProviderFactories: testAccProviderFactories,
			CheckDestroy:      testAccCheckAivenClickhouseUserResourceDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccClickhouseUserResource(rName),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckAivenClickhouseUserAttributes("data.aiven_clickhouse_user.user"),
						resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
						resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
						resource.TestCheckResourceAttr(resourceName, "username", fmt.Sprintf("user-%s", rName)),
						resource.TestCheckResourceAttrSet(resourceName, "password"),
						resource.TestCheckResourceAttrSet(resourceName, "uuid"),
						resource.TestCheckResourceAttrSet(resourceName, "required"),
					),
				},
			},
		})
	})
}

func testAccCheckAivenClickhouseUserResourceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*aiven.Client)

	// loop through the resources in state, verifying each aiven_clickhouse_user is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_clickhouse_user" {
			continue
		}

		projectName, serviceName, uuid := schemautil.SplitResourceID3(rs.Primary.ID)
		p, err := c.ClickhouseUser.Get(projectName, serviceName, uuid)
		if err != nil {
			if err.(aiven.Error).Status != 404 {
				return err
			}
		}

		if p != nil {
			return fmt.Errorf("clickhouse user (%s) still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccClickhouseUserResource(name string) string {
	return fmt.Sprintf(`
		data "aiven_project" "foo" {
		  project = "%s"
		}
		
		resource "aiven_clickhouse" "bar" {
		  project                 = data.aiven_project.foo.project
		  cloud_name              = "google-europe-west1"
		  plan                    = "business-beta-8"
		  service_name            = "test-acc-sr-%s"
		  maintenance_window_dow  = "monday"
		  maintenance_window_time = "10:00:00"
		}
		
		resource "aiven_clickhouse_user" "foo" {
		  service_name = aiven_clickhouse.bar.service_name
		  project      = aiven_clickhouse.bar.project
		  username     = "user-%s"
		}
		
		data "aiven_clickhouse_user" "user" {
		  service_name = aiven_clickhouse_user.foo.service_name
		  project      = aiven_clickhouse_user.foo.project
		  username     = aiven_clickhouse_user.foo.username
		}`,
		os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}

func testAccCheckAivenClickhouseUserAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		log.Printf("[DEBUG] clickhouse user attributes %v", a)

		if a["username"] == "" {
			return fmt.Errorf("expected to get a clikchouse user username from Aiven")
		}

		if a["project"] == "" {
			return fmt.Errorf("expected to get a clickhouse user project from Aiven")
		}

		if a["service_name"] == "" {
			return fmt.Errorf("expected to get a clickhouse user service_name from Aiven")
		}

		return nil
	}
}
