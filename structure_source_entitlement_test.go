package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testEntitlementConf      []*SourceEntitlement
	testEntitlementInterface map[string]interface{}
)

func init() {
	testEntitlementConf = []*SourceEntitlement{
		{
			Source: &SourceInfo{
				ID:   "2c9180887412345678948078d29f2e46",
				Name: "Active Directory test",
			},
			Attribute:              "memberOf",
			Created:                nil,
			Description:            "test description",
			Name:                   "test name",
			Modified:               "06/20/2020",
			Owner:                  nil,
			Privileged:             false,
			SourceSchemaObjectType: "group",
			Value:                  "CN=example,OU=Groups,DC=test,DC=com",
		},
	}
	testEntitlementInterface =
		map[string]interface{}{
			"source_name":               "Active Directory test",
			"source_id":                 "2c9180887412345678948078d29f2e46",
			"attribute":                 "memberOf",
			"description":               "test description",
			"name":                      "test name",
			"modified":                  "06/20/2020",
			"privileged":                false,
			"source_schema_object_type": "group",
			"value":                     "CN=example,OU=Groups,DC=test,DC=com",
		}
}

func TestFlattenSourceEntitlement(t *testing.T) {

	cases := []struct {
		Input          []*SourceEntitlement
		ExpectedOutput map[string]interface{}
	}{
		{
			testEntitlementConf,
			testEntitlementInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, sourceEntitlementFields(), tc.ExpectedOutput)
		err := flattenSourceEntitlement(output, tc.Input[0])
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		expectedOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			expectedOutput[k] = output.Get(k)
		}
		if !reflect.DeepEqual(expectedOutput, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, expectedOutput)
		}
	}
}
