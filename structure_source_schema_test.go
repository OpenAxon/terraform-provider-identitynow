package main

import (
	"reflect"
	"testing"
)

var (
	testSchemasConf      []*Schema
	testSchemasInterface []interface{}
)

func init() {
	testSchemasConf = []*Schema{
		{
			Type: "TYPE",
			ID:   "1234-id",
			Name: "test-name",
		},
	}
	testSchemasInterface = []interface{}{
		map[string]interface{}{
			"type": "TYPE",
			"id":   "1234-id",
			"name": "test-name",
		},
	}
}

func TestFlattenSourceSchemas(t *testing.T) {

	cases := []struct {
		Input          []*Schema
		ExpectedOutput []interface{}
	}{
		{
			testSchemasConf,
			testSchemasInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceSchema(tc.Input, []interface{}{})
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceSchemas(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*Schema
	}{
		{
			testSchemasInterface,
			testSchemasConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceSchema(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
