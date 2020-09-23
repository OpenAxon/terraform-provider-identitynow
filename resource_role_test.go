package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"os"
	"testing"
)

const (
	testAccIdentitynowRoleType = "identitynow_role"
)

var (
	testAccRole       string
	testAccRoleUpdate string
	roleName          string
	ownerName         string
)

func init() {
	ownerName = os.Getenv("IDENTITYNOW_OWNER_NAME")
	roleName = acctest.RandomWithPrefix("random_role")
	testAccRole = `
resource "` + testAccIdentitynowRoleType + `" "foo" {
  name = "` + roleName + `"
  description = "test description"
  requestable = true
  owner = "` + ownerName + `"
  lifecycle {
    ignore_changes = [
      name,
      display_name
    ]
  }
}
`

	testAccRoleUpdate = `
resource "` + testAccIdentitynowRoleType + `" "foo" {
  name = "` + roleName + `"
  description = "test description changed"
  requestable = true
  owner = "` + ownerName + `"
  lifecycle {
    ignore_changes = [
      name,
      display_name
    ]
  }
}
`

}

func TestAccRole_basic(t *testing.T) {
	var role Role

	log.Printf("role:%s\n", testAccRole)
	log.Printf("updated role:%s\n", testAccRoleUpdate)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIdentitynowRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRole,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExist(testAccIdentitynowRoleType+".foo", role),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "description", "test description"),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "requestable", "true"),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "owner", ownerName),
				),
			},
			{
				Config: testAccRoleUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExist(testAccIdentitynowRoleType+".foo", role),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "description", "test description changed"),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "requestable", "true"),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "owner", ownerName),
				),
			},
			{
				Config: testAccRole,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoleExist(testAccIdentitynowRoleType+".foo", role),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "description", "test description"),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "requestable", "true"),
					resource.TestCheckResourceAttr(testAccIdentitynowRoleType+".foo", "owner", ownerName),
				),
			},
		},
	})

}

func testAccIdentitynowRoleDestroy(state *terraform.State) error {
	for _, rs := range state.RootModule().Resources {
		if rs.Type != testAccIdentitynowRoleType {
			continue
		}

		client, err := testAccProvider.Meta().(*Config).IdentityNowClient()
		if err != nil {
			return err
		}
		foundRole, _ := client.GetRole(context.Background(), rs.Primary.ID)
		if len(foundRole.ID) > 0 &&
			foundRole.ID != rs.Primary.ID {
			return fmt.Errorf("role still exist")
		}
	}
	return nil
}

func testAccCheckRoleExist(name string, role Role) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no role ID is set")
		}

		client, err := testAccProvider.Meta().(*Config).IdentityNowClient()
		if err != nil {
			return err
		}

		foundRole, err := client.GetRole(context.Background(), rs.Primary.ID)
		if err != nil {
			log.Fatal(err)
		}
		// we expect a single Source by this ID. If we find zero
		// then we consider this an error
		if len(foundRole.ID) == 0 ||
			foundRole.ID != rs.Primary.ID || err != nil {
			return fmt.Errorf("role not found")
		}
		role = *foundRole

		return nil
	}

}
