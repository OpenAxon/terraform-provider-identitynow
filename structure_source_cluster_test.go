package main

import (
	"reflect"
	"testing"
)

var (
	testClusterConf      *Cluster
	testClusterInterface []interface{}
)

func init() {
	testClusterConf = &Cluster{
		Type: "CLUSTER",
		ID:   "1234-id",
		Name: "test-name",
	}
	testClusterInterface = []interface{}{
		map[string]interface{}{
			"type": "CLUSTER",
			"id":   "1234-id",
			"name": "test-name",
		},
	}
}

func TestFlattenSourceCluster(t *testing.T) {

	cases := []struct {
		Input          *Cluster
		ExpectedOutput []interface{}
	}{
		{
			testClusterConf,
			testClusterInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceCluster(tc.Input, tc.ExpectedOutput)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceCluster(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *Cluster
	}{
		{
			testClusterInterface,
			testClusterConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceCluster(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
