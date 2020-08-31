package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var (
	testAccProvider               *schema.Provider
	testAccProviders              map[string]terraform.ResourceProvider
	testAccCheckIdentitynowConfig string
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"identitynow": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("Error: %s", err)
	}
}
func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IDENTITYNOW_URL"); v == "" {
		t.Fatal("IDENTITYNOW_URL must be set for acceptance tests")
	}
	if v := os.Getenv("IDENTITYNOW_CLIENT_ID"); v == "" {
		t.Fatal("IDENTITYNOW_CLIENT_ID must be set for acceptance tests")
	}
	if v := os.Getenv("IDENTITYNOW_CLIENT_SECRET"); v == "" {
		t.Fatal("IDENTITYNOW_CLIENT_SECRET must be set for acceptance tests")
	}
	if v := os.Getenv("IDENTITYNOW_OWNER_ID"); v == "" {
		t.Fatal("IDENTITYNOW_OWNER_ID must be set for acceptance tests")
	}
	if v := os.Getenv("IDENTITYNOW_OWNER_NAME"); v == "" {
		t.Fatal("IDENTITYNOW_OWNER_NAME must be set for acceptance tests")
	}
	if v := os.Getenv("IDENTITYNOW_CLUSTER_ID"); v == "" {
		t.Fatal("IDENTITYNOW_CLUSTER_ID must be set for acceptance tests")
	}
	if v := os.Getenv("IDENTITYNOW_CLUSTER_NAME"); v == "" {
		t.Fatal("IDENTITYNOW_CLUSTER_NAME must be set for acceptance tests")
	}
}
