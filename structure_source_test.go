package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testSourceConf      *Source
	testSourceInterface map[string]interface{}
	d                   *schema.ResourceData
)

func init() {
	testSourceConf = &Source{
		Name:            "foo",
		Description:     "test description",
		Connector:       "active-directory",
		DeleteThreshold: 10,
		Authoritative:   false,
	}
	testSourceInterface = map[string]interface{}{
		"name":             "foo",
		"description":      "test description",
		"connector":        "active-directory",
		"delete_threshold": 10,
		"authoritative":    false,
	}
}

func TestFlattenSource(t *testing.T) {
	cases := []struct {
		Input          *Source
		ExpectedOutput map[string]interface{}
	}{
		{
			testSourceConf,
			testSourceInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, sourceFields(), tc.ExpectedOutput)
		err := flattenSource(output, tc.Input)
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

func TestExpandSource(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *Source
	}{
		{
			testSourceInterface,
			testSourceConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, sourceFields(), tc.Input)
		output, err := expandSource(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
