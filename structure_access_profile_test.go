package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testAccessProfileConf      *AccessProfile
	testAccessProfileInterface map[string]interface{}
)

func init() {
	FALSE := false
	testAccessProfileConf = &AccessProfile{
		Name:        "test name",
		Description: "test Description",
		AccessProfileOwner: &ObjectInfo{
			ID:   "2c9180887412345678948078d29f2e46",
			Name: "SRE Test",
			Type: "IDENTITY",
		},
		AccessProfileSource: &ObjectInfo{
			ID:   "2c91808374bc866a0178948078d29f2e46",
			Name: "Product platform, Azure portal, AzureUSGovernment",
			Type: "SOURCE",
		},
		Entitlements: []*ObjectInfo{
			{
				ID:   "2c918088747654398948078d29f2e46",
				Name: "Operators_AG1",
				Type: "ENTITLEMENT",
			},
			{
				ID:   "2c918009437654398948078d29f2e46",
				Name: "Integrator_AG1",
				Type: "ENTITLEMENT",
			},
		},
		Enabled: &FALSE,
	}
	testAccessProfileInterface = map[string]interface{}{
		"name":        "test name",
		"description": "test Description",
		"source": []interface{}{
			map[string]interface{}{
				"id":   "2c91808374bc866a0178948078d29f2e46",
				"name": "Product platform, Azure portal, AzureUSGovernment",
				"type": "SOURCE",
			},
		},
		"owner": []interface{}{
			map[string]interface{}{
				"id":   "2c9180887412345678948078d29f2e46",
				"name": "SRE Test",
				"type": "IDENTITY",
			},
		},
		"entitlements": []interface{}{
			map[string]interface{}{
				"id":   "2c918088747654398948078d29f2e46",
				"name": "Operators_AG1",
				"type": "ENTITLEMENT",
			},
			map[string]interface{}{
				"id":   "2c918009437654398948078d29f2e46",
				"name": "Integrator_AG1",
				"type": "ENTITLEMENT",
			},
		},
		"enabled": false,
	}
}

func TestFlattenAccessProfile(t *testing.T) {
	cases := []struct {
		Input          *AccessProfile
		ExpectedOutput map[string]interface{}
	}{
		{
			testAccessProfileConf,
			testAccessProfileInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, accessProfileFields(), tc.ExpectedOutput)
		err := flattenAccessProfile(output, tc.Input)
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

func TestExpandAccessProfile(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *AccessProfile
	}{
		{
			testAccessProfileInterface,
			testAccessProfileConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, accessProfileFields(), tc.Input)
		output, err := expandAccessProfile(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
