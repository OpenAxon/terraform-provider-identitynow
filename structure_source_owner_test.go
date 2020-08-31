package main

import (
	"reflect"
	"testing"
)

var (
	testOwnerConf      *Owner
	testOwnerInterface []interface{}
)

func init() {
	testOwnerConf = &Owner{
		Type: "IDENTITY",
		ID:   "1234-id",
		Name: "test-name",
	}
	testOwnerInterface = []interface{}{
		map[string]interface{}{
			"type": "IDENTITY",
			"id":   "1234-id",
			"name": "test-name",
		},
	}
}

func TestFlattenSourceOwner(t *testing.T) {

	cases := []struct {
		Input          *Owner
		ExpectedOutput []interface{}
	}{
		{
			testOwnerConf,
			testOwnerInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceOwner(tc.Input, tc.ExpectedOutput)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceOwner(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *Owner
	}{
		{
			testOwnerInterface,
			testOwnerConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceOwner(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
