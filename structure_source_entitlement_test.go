package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testEntitlementConf      *Items
	testEntitlementInterface map[string]interface{}
)

func init() {
	testEntitlementConf = &Items{
		SourceID:          "2c9180887412345678948078d29f2e46",
		SourceName:        "Active Directory test",
		Attribute:         "memberOf",
		CreatedTime:       nil,
		DeletedTime:       "05/21/2020",
		Description:       "test description",
		DirectPermissions: []interface{}{"test"},
		DisplayName:       "test name",
		DisplayableName:   "test name",
		LastModifiedTime:  "06/20/2020",
		OwnerID:           "1234",
		OwnerUID:          "123a4",
		Privileged:        false,
		Schema:            "group",
		Value:             "CN=example,OU=Groups,DC=test,DC=com",
	}
	testEntitlementInterface = map[string]interface{}{
		"source_id":          "2c9180887412345678948078d29f2e46",
		"source_name":        "Active Directory test",
		"attribute":          "memberOf",
		"deleted_time":       "05/21/2020",
		"description":        "test description",
		"direct_permissions": []interface{}{"test"},
		"display_name":       "test name",
		"displayable_name":   "test name",
		"last_modified_time": "06/20/2020",
		"owner_id":           "1234",
		"owner_uid":          "123a4",
		"privileged":         false,
		"schema":             "group",
		"value":              "CN=example,OU=Groups,DC=test,DC=com",
	}
}

func TestFlattenSourceEntitlement(t *testing.T) {

	cases := []struct {
		Input          *Items
		ExpectedOutput map[string]interface{}
	}{
		{
			testEntitlementConf,
			testEntitlementInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, sourceEntitlementFields(), tc.ExpectedOutput)
		err := flattenSourceEntitlement(output, tc.Input)
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
