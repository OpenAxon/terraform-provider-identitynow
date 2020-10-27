package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testAccountSchemaAttributeConf      *AccountSchemaAttribute
	testAccountSchemaAttributeInterface map[string]interface{}
)

func init() {
	testAccountSchemaAttributeConf = &AccountSchemaAttribute{
		Description:       "test description",
		DisplayAttribute:  false,
		Entitlement:       true,
		IdentityAttribute: false,
		Managed:           false,
		Minable:           false,
		Multi:             false,
		Name:              "test",
		Type:              "string",
		ObjectType:        "account",
		SourceID:          "1234",
	}
	testAccountSchemaAttributeInterface = map[string]interface{}{
		"description":        "test description",
		"display_attribute":  false,
		"entitlement":        true,
		"identity_attribute": false,
		"managed":            false,
		"minable":            false,
		"multi":              false,
		"name":               "test",
		"type":               "string",
		"object_type":        "account",
		"source_id":          "1234",
	}
}

func TestFlattenAccountSchemaAttribute(t *testing.T) {
	cases := []struct {
		Input          *AccountSchemaAttribute
		ExpectedOutput map[string]interface{}
	}{
		{
			testAccountSchemaAttributeConf,
			testAccountSchemaAttributeInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, accountSchemaAttributeFields(), tc.ExpectedOutput)
		err := flattenAccountSchemaAttribute(output, tc.Input)
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
		ExpectedOutput *AccountSchemaAttribute
	}{
		{
			testAccountSchemaAttributeInterface,
			testAccountSchemaAttributeConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, accountSchemaAttributeFields(), tc.Input)
		output, err := expandAccountSchemaAttribute(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
