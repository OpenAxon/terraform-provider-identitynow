package main

import (
	"reflect"
	"testing"
)

var (
	testSourcePassPoliciesConf      []*SourcePasswordPolicies
	testSourcePassPoliciesInterface []interface{}
)

func init() {
	testSourcePassPoliciesConf = []*SourcePasswordPolicies{
		{
			Type: "PASSWORD_POLICY",
			ID:   "1234-id",
			Name: "test-name",
		},
	}
	testSourcePassPoliciesInterface = []interface{}{
		map[string]interface{}{
			"type": "PASSWORD_POLICY",
			"id":   "1234-id",
			"name": "test-name",
		},
	}
}

func TestFlattenSourcePasswordPolicies(t *testing.T) {

	cases := []struct {
		Input          []*SourcePasswordPolicies
		ExpectedOutput []interface{}
	}{
		{
			testSourcePassPoliciesConf,
			testSourcePassPoliciesInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourcePasswordPolicies(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourcePasswordPolicies(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*SourcePasswordPolicies
	}{
		{
			testSourcePassPoliciesInterface,
			testSourcePassPoliciesConf,
		},
	}

	for _, tc := range cases {
		output := expandSourcePasswordPolicies(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
