package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	log "log"
	"os"
	"testing"
)

const (
	testAccIdentitynowAccessProfileType = "identitynow_access_profile"
)

var (
	testAccAccessProfile       string
	testAccAccessProfileUpdate string
	accessProfileName          string
	ownerId                    string
	sourceId                   string
	entitlement                string
)

func init() {
	ownerId = os.Getenv("IDENTITYNOW_OWNER_ID")
	sourceId = os.Getenv("IDENTITYNOW_SOURCE_ID")
	entitlement = os.Getenv("IDENTITYNOW_SOURCE_ENTITLEMENT")
	accessProfileName = acctest.RandomWithPrefix("access_profile")

	testAccAccessProfile = `
resource "` + testAccIdentitynowAccessProfileType + `" "foo" {
  name = "` + accessProfileName + `"
  description = "test description"
  entitlements = ["` + entitlement + `"]
  owner_id = ` + ownerId + `
  source_id = ` + sourceId + `
  approval_schemes = "manager"
  request_comments_required = "true"
  denied_comments_required = "true"
}
`

	testAccAccessProfileUpdate = `
resource "` + testAccIdentitynowAccessProfileType + `" "foo" {
  name = "` + accessProfileName + `"
  description = "test description changed"
  entitlements = ["` + entitlement + `"]
  owner_id = ` + ownerId + `
  source_id = ` + sourceId + `
  approval_schemes = "manager"
  request_comments_required = "true"
  denied_comments_required = "true"
}
`

}

func TestAccAccessProfile_basic(t *testing.T) {
	var accessProfile AccessProfile

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIdentitynowAccessProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAccessProfile,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessProfileExist(testAccIdentitynowAccessProfileType+".foo", accessProfile),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "name", accessProfileName),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "description", "test description"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "owner_id", ownerId),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "source_id", sourceId),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "approval_schemes", "manager"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "request_comments_required", "true"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "denied_comments_required", "true"),
				),
			},
			{
				Config: testAccAccessProfileUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessProfileExist(testAccIdentitynowAccessProfileType+".foo", accessProfile),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "name", accessProfileName),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "description", "test description changed"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "owner_id", ownerId),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "source_id", sourceId),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "approval_schemes", "manager"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "request_comments_required", "true"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "denied_comments_required", "true"),
				),
			},
			{
				Config: testAccAccessProfile,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccessProfileExist(testAccIdentitynowAccessProfileType+".foo", accessProfile),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "name", accessProfileName),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "description", "test description"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "owner_id", ownerId),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "source_id", sourceId),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "approval_schemes", "manager"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "request_comments_required", "true"),
					resource.TestCheckResourceAttr(testAccIdentitynowAccessProfileType+".foo", "denied_comments_required", "true"),
				),
			},
		},
	})

}

func testAccIdentitynowAccessProfileDestroy(state *terraform.State) error {
	for _, rs := range state.RootModule().Resources {
		if rs.Type != testAccIdentitynowAccessProfileType {
			continue
		}

		client, err := testAccProvider.Meta().(*Config).IdentityNowClient()
		if err != nil {
			return err
		}
		foundAccessProfile, _ := client.GetAccessProfile(context.Background(), rs.Primary.ID)
		if len(foundAccessProfile.ID) > 0 &&
			foundAccessProfile.ID != rs.Primary.ID {
			return fmt.Errorf("access profile still exist")
		}
	}
	return nil
}

func testAccCheckAccessProfileExist(name string, accessProfile AccessProfile) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Access Profile ID is set")
		}

		client, err := testAccProvider.Meta().(*Config).IdentityNowClient()
		if err != nil {
			return err
		}

		foundAccessProfile, err := client.GetAccessProfile(context.Background(), rs.Primary.ID)
		if err != nil {
			log.Fatal(err)
		}
		// we expect a single Source by this ID. If we find zero
		// then we consider this an error
		if len(foundAccessProfile.ID) == 0 ||
			foundAccessProfile.ID != rs.Primary.ID || err != nil {
			return fmt.Errorf("access profile not found")
		}
		accessProfile = *foundAccessProfile

		return nil
	}

}
