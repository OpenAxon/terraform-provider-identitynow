package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	reflect "reflect"
	"testing"
)

var (
	testRoleConf      *Role
	testRoleInterface map[string]interface{}
)

func init() {
	TRUE := true
	FALSE := false
	testRoleConf = &Role{
		Description: "test description",
		Enabled:     &TRUE,
		Name:        "test name",
		RoleOwner: &ObjectInfo{
			ID:   "2c9180887412345678948078d29f2e46",
			Name: "SRE Test",
			Type: "IDENTITY",
		},
		AccessProfiles: []*ObjectInfo{
			{
				ID:   "2c918088747654398948078d29f2e46",
				Name: "Test Developer",
				Type: "ACCESS_PROFILE",
			},
			{
				ID:   "2c918009437654398948078d29f2e46",
				Name: "Test Operator",
				Type: "ACCESS_PROFILE",
			},
		},
		Requestable: &FALSE,
	}
	testRoleInterface = map[string]interface{}{
		//"accessProfiles": []map[string]interface{}{
		//	{
		//		"id":   "2c918088747654398948078d29f2e46",
		//		"name": "Test Developer",
		//		"type": "ACCESS_PROFILE",
		//	},
		//	{
		//		"id":   "2c918009437654398948078d29f2e46",
		//		"name": "Test Operator",
		//		"type": "ACCESS_PROFILE",
		//	},
		//},
		"owner": map[string]interface{}{
			"id":   "2c9180887412345678948078d29f2e46",
			"name": "SRE Test",
			"type": "IDENTITY",
		},
		"description": "test description",
		"enabled":     true,
		"name":        "test name",
		"requestable": false,
	}
}

func TestFlattenRole(t *testing.T) {
	cases := []struct {
		Input          *Role
		ExpectedOutput map[string]interface{}
	}{
		{
			testRoleConf,
			testRoleInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, roleFields(), tc.ExpectedOutput)
		err := flattenRole(output, tc.Input)
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

func TestExpandRole(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *Role
	}{
		{
			testRoleInterface,
			testRoleConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, roleFields(), tc.Input)
		output, err := expandRole(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
