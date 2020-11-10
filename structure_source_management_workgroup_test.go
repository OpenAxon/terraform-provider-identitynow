package main

import (
	"reflect"
	"testing"
)

var (
	testManagementWorkgroupConf      *ManagementWorkgroup
	testManagementWorkgroupInterface []interface{}
)

func init() {
	testManagementWorkgroupConf = &ManagementWorkgroup{
		Type: "GOVERNANCE_GROUP",
		ID:   "1234-id",
		Name: "test-name",
	}
	testManagementWorkgroupInterface = []interface{}{
		map[string]interface{}{
			"type": "GOVERNANCE_GROUP",
			"id":   "1234-id",
			"name": "test-name",
		},
	}
}

func TestFlattenSourceManagementWorkgroup(t *testing.T) {

	cases := []struct {
		Input          *ManagementWorkgroup
		ExpectedOutput []interface{}
	}{
		{
			testManagementWorkgroupConf,
			testManagementWorkgroupInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceManagementWorkgroup(tc.Input, []interface{}{})
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceManagementWorkgroup(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ManagementWorkgroup
	}{
		{
			testManagementWorkgroupInterface,
			testManagementWorkgroupConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceManagementWorkgroup(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
