package main

import (
	"reflect"
	"testing"
)

var (
	testAccountCorrelationConfigConf      *AccountCorrelationConfig
	testAccountCorrelationConfigInterface []interface{}
)

func init() {
	testAccountCorrelationConfigConf = &AccountCorrelationConfig{
		Type: "TYPE",
		ID:   "1234-id",
		Name: "test-name",
	}
	testAccountCorrelationConfigInterface = []interface{}{
		map[string]interface{}{
			"type": "TYPE",
			"id":   "1234-id",
			"name": "test-name",
		},
	}
}

func TestFlattenSourceAccountCorrelationConfig(t *testing.T) {

	cases := []struct {
		Input          *AccountCorrelationConfig
		ExpectedOutput []interface{}
	}{
		{
			testAccountCorrelationConfigConf,
			testAccountCorrelationConfigInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceAccountCorrelationConfig(tc.Input, tc.ExpectedOutput)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceAccountCorrelationConfig(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *AccountCorrelationConfig
	}{
		{
			testAccountCorrelationConfigInterface,
			testAccountCorrelationConfigConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceAccountCorrelationConfig(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
