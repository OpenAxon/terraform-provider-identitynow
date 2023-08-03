package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testAccountSchemaConf      *AccountSchema
	testAccountSchemaInterface map[string]interface{}
)

func init() {
	testAccountSchemaConf = &AccountSchema{
		DisplayAttribute:   "distinguishedName",
		IdentityAttribute:  "sAMAccountName",
		NativeObjectType:   "User",
		Name:               "account",
		SourceID:           "2c9180835d191a86015d28455b4a2329",
		ID:                 "2c9180835d191a86015d28455b4a1234",
		HierarchyAttribute: "memberOf",
		IncludePermissions: false,
		Created:            "2019-12-24T22:32:58.104Z",
		Modified:           "2019-12-31T20:22:28.104Z",
		Attributes: []*AccountSchemaAttribute{
			{
				Name:          "sAMAccountName",
				Type:          "STRING",
				IsEntitlement: true,
				IsGroup:       false,
				IsMultiValued: false,
			},
			{
				Name:          "employeeId",
				Type:          "STRING",
				IsEntitlement: false,
				IsGroup:       false,
				IsMultiValued: false,
				Description:   "Employee Id",
			},
			{
				Name: "memberOf",
				Type: "STRING",
				Schema: &Schema{
					Type: "CONNECTOR_SCHEMA",
					ID:   "2c9180887671ff8c01767b4671fc7d60",
					Name: "group",
				},
				Description:   "Group membership",
				IsMultiValued: true,
				IsEntitlement: true,
				IsGroup:       true,
			},
		},
	}
	testAccountSchemaInterface = map[string]interface{}{
		"source_id":           "2c9180835d191a86015d28455b4a2329",
		"schema_id":           "2c9180835d191a86015d28455b4a1234",
		"name":                "account",
		"native_object_type":  "User",
		"identity_attribute":  "sAMAccountName",
		"display_attribute":   "distinguishedName",
		"hierarchy_attribute": "memberOf",
		"include_permissions": false,
		"attributes": []interface{}{
			map[string]interface{}{
				"name":           "sAMAccountName",
				"type":           "STRING",
				"is_multiValued": false,
				"is_entitlement": false,
				"is_group":       false,
			},
			map[string]interface{}{
				"name":            "employeeId",
				"type":            "STRING",
				"is_multi_valued": false,
				"is_entitlement":  false,
				"is_group":        false,
				"description":     "Employee ID",
			},
			map[string]interface{}{
				"name": "memberOf",
				"type": "STRING",
				"schema": map[string]interface{}{
					"type": "CONNECTOR_SCHEMA",
					"id":   "2c9180887671ff8c01767b4671fc7d60",
					"name": "group",
				},
				"description":     "Group membership",
				"is_multi_valued": true,
				"is_entitlement":  true,
				"is_group":        true,
			},
		},
		"created":  "2019-12-24T22:32:58.104Z",
		"modified": "2019-12-31T20:22:28.104Z",
	}
}

func TestFlattenAccountSchemaAttribute(t *testing.T) {
	cases := []struct {
		Input          *AccountSchema
		ExpectedOutput map[string]interface{}
	}{
		{
			testAccountSchemaConf,
			testAccountSchemaInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, accountSchemaFields(), tc.ExpectedOutput)
		err := flattenAccountSchema(output, tc.Input)
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

func TestExpandAccountSchemaAttribute(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *AccountSchema
	}{
		{
			testAccountSchemaInterface,
			testAccountSchemaConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, accountSchemaFields(), tc.Input)
		output, err := expandAccountSchema(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
