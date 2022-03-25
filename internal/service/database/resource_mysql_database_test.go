package database_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("aiven_mysql_database", &resource.Sweeper{
		Name: "aiven_mysql_database",
		F:    acc.SweepDatabases,
		Dependencies: []string{
			"aiven_connection_pool",
		},
	})
}

func TestAccAivenDatabaseMySQL_basic(t *testing.T) {
	resourceName := "aiven_mysql_database.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rName2 := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acc.TestAccPreCheck(t) },
		ProviderFactories: acc.TestAccProviderFactories,
		CheckDestroy:      testAccCheckAivenDatabaseMySQLResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDatabaseMySQLResource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenDatabaseMySQLAttributes("data.aiven_mysql_database.database"),
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "database_name", fmt.Sprintf("test-acc-db-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "termination_protection", "false"),
				),
			},
			{
				Config:                    testAccDatabaseMySQLTerminationProtectionResource(rName2),
				PreventPostDestroyRefresh: true,
				ExpectNonEmptyPlan:        true,
				PlanOnly:                  true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName2)),
					resource.TestCheckResourceAttr(resourceName, "database_name", fmt.Sprintf("test-acc-db-%s", rName2)),
					resource.TestCheckResourceAttr(resourceName, "termination_protection", "true"),
				),
			},
		},
	})
}

func testAccCheckAivenDatabaseMySQLResourceDestroy(s *terraform.State) error {
	c := acc.TestAccProvider.Meta().(*aiven.Client)

	// loop through the resources in state, verifying each database is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_database" {
			continue
		}

		projectName, serviceName, databaseName := schemautil.SplitResourceID3(rs.Primary.ID)
		db, err := c.Databases.Get(projectName, serviceName, databaseName)
		if err != nil {
			if err.(aiven.Error).Status != 404 {
				return err
			}
		}

		if db != nil {
			return fmt.Errorf("database (%s) still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccDatabaseMySQLResource(name string) string {
	return fmt.Sprintf(`
		data "aiven_project" "foo" {
		  project = "%s"
		}
				
		resource "aiven_mysql" "bar" {
			project                 = data.aiven_project.foo.project
			cloud_name              = "google-europe-west1"
			plan                    = "startup-4"
			service_name            = "test-acc-sr-%s"
			maintenance_window_dow  = "monday"
			maintenance_window_time = "10:00:00"
		  
			mysql_user_config {
			  mysql {
				sql_mode                = "ANSI,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION,NO_ZERO_DATE,NO_ZERO_IN_DATE"
				sql_require_primary_key = true
			  }
		  
			  public_access {
				mysql = true
			  }
			}
		  }
  
		
		resource "aiven_mysql_database" "foo" {
		  project       = aiven_mysql.bar.project
		  service_name  = aiven_mysql.bar.service_name
		  database_name = "test-acc-db-%s"
		}
		
		data "aiven_mysql_database" "database" {
		  project       = aiven_mysql_database.foo.project
		  service_name  = aiven_mysql_database.foo.service_name
		  database_name = aiven_mysql_database.foo.database_name
		
		  depends_on = [aiven_mysql_database.foo]
		}`,
		os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}

func testAccDatabaseMySQLTerminationProtectionResource(name string) string {
	return fmt.Sprintf(`
		data "aiven_project" "foo" {
		  project = "%s"
		}
		
		resource "aiven_mysql" "bar" {
		  project                 = data.aiven_project.foo.project
		  cloud_name              = "google-europe-west1"
		  plan                    = "startup-4"
		  service_name            = "test-acc-sr-%s"
		  maintenance_window_dow  = "monday"
		  maintenance_window_time = "10:00:00"
		
		  mysql_user_config {
		    mysql {
		      sql_mode                = "ANSI,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION,NO_ZERO_DATE,NO_ZERO_IN_DATE"
		      sql_require_primary_key = true
		    }
		
		    public_access {
		      mysql = true
		    }
		  }
		}
		
		resource "aiven_mysql_database" "foo" {
		  project                = aiven_mysql.bar.project
		  service_name           = aiven_mysql.bar.service_name
		  database_name          = "test-acc-db-%s"
		  termination_protection = true
		}
		
		data "aiven_mysql_database" "database" {
		  project       = aiven_mysql_database.foo.project
		  service_name  = aiven_mysql_database.foo.service_name
		  database_name = aiven_mysql_database.foo.database_name
		
		  depends_on = [aiven_mysql_database.foo]
		}`,
		os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}

func testAccCheckAivenDatabaseMySQLAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["project"] == "" {
			return fmt.Errorf("expected to get a project name from Aiven")
		}

		if a["service_name"] == "" {
			return fmt.Errorf("expected to get a service_name from Aiven")
		}

		if a["database_name"] == "" {
			return fmt.Errorf("expected to get a database_name from Aiven")
		}

		return nil
	}
}
