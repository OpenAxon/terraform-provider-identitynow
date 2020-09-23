package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

const (
	testAccIdentitynowSourceType = "identitynow_source"
)

var (
	testAccSource       string
	testAccSourceUpdate string
	rName               string
)

func init() {
	ownerId := os.Getenv("IDENTITYNOW_EXTERNAL_OWNER_ID")
	ownerName := os.Getenv("IDENTITYNOW_OWNER_NAME")
	clusterId := os.Getenv("IDENTITYNOW_CLUSTER_ID")
	clusterName := os.Getenv("IDENTITYNOW_CLUSTER_NAME")
	rName = acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	testAccSource = `
resource "` + testAccIdentitynowSourceType + `"  "foo" {
  name = "` + rName + `"
  description = "The Azure Active Directory acceptance test"
  connector = "azure-active-directory"
  authoritative = false
  delete_threshold = 10

  owner {
    id = "` + ownerId + `"
    name = "` + ownerName + `"
    type = "IDENTITY"
  }

  cluster {
    id = "` + clusterId + `"
    name = "` + clusterName + `"
    type = "CLUSTER"
  }

  connector_attributes {
    grant_type = "CLIENT_CREDENTIALS"
    client_id = "test-client-id"
    client_secret = "test-client-secret"
    domain_name = "us1.axonengineering.io"
    ms_graph_resource_base = "https://graph.microsoft.us"
    ms_graph_token_base = "https://login.microsoftonline.us"
    azure_ad_graph_resource_base = "https://graph.microsoftazure.us"
    azure_ad_graph_token_base = "https://login.microsoftonline.us"
  }
}
`
	testAccSourceUpdate = `
resource "identitynow_source" "foo" {
  name = "` + rName + `"
  description = "The Azure Active Directory acceptance test update"
  connector = "azure-active-directory"
  delete_threshold = 10
  authoritative = false

  owner {
    id = "` + ownerId + `"
    name = "` + ownerName + `"
    type = "IDENTITY"
  }

  cluster {
    id = "` + clusterId + `"
    name = "` + clusterName + `"
	type = "CLUSTER"
  }

  connector_attributes {
    grant_type = "CLIENT_CREDENTIALS"
    client_id = "test-client-id"
    client_secret = "test-client-secret"
    domain_name = "us1.axonengineering.io"
    ms_graph_resource_base = "https://graph.microsoft.us"
    ms_graph_token_base = "https://login.microsoftonline.us"
    azure_ad_graph_resource_base = "https://graph.microsoftazure.us"
    azure_ad_graph_token_base = "https://login.microsoftonline.us"
  }
}
`
}

func TestAccSource_basic(t *testing.T) {
	var source Source

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIdentitynowSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceExist(testAccIdentitynowSourceType+".foo", source),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "name", rName),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "description", "The Azure Active Directory acceptance test"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "connector", "azure-active-directory"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "authoritative", "false"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "delete_threshold", "10"),
				),
			},
			{
				Config: testAccSourceUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceExist(testAccIdentitynowSourceType+".foo", source),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "name", rName),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "description", "The Azure Active Directory acceptance test update"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "connector", "azure-active-directory"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "authoritative", "false"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "delete_threshold", "10"),
				),
			},
			{
				Config: testAccSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceExist(testAccIdentitynowSourceType+".foo", source),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "name", rName),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "description", "The Azure Active Directory acceptance test"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "connector", "azure-active-directory"),
					resource.TestCheckResourceAttr(testAccIdentitynowSourceType+".foo", "authoritative", "false"),
				),
			},
		},
	})

}

func testAccIdentitynowSourceDestroy(state *terraform.State) error {
	for _, rs := range state.RootModule().Resources {
		if rs.Type != testAccIdentitynowSourceType {
			continue
		}

		client, err := testAccProvider.Meta().(*Config).IdentityNowClient()
		if err != nil {
			return err
		}
		foundSource, _ := client.GetSource(context.Background(), rs.Primary.ID)
		if len(foundSource.ID) > 0 &&
			foundSource.ID != rs.Primary.ID {
			return fmt.Errorf("Source still exist")
		}
	}
	return nil
}

func testAccCheckSourceExist(name string, source Source) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No source ID is set")
		}

		client, err := testAccProvider.Meta().(*Config).IdentityNowClient()
		if err != nil {
			return err
		}

		foundSource, err := client.GetSource(context.Background(), rs.Primary.ID)
		// we expect a single Source by this ID. If we find zero
		// then we consider this an error
		if len(foundSource.ID) == 0 ||
			foundSource.ID != rs.Primary.ID || err != nil {
			return fmt.Errorf("Source not found.")
		}
		source = *foundSource

		return nil
	}

}
