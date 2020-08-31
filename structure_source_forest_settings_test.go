package main

import (
	"reflect"
	"testing"
)

var (
	testForestSettingsConf      []*ForestSettings
	testForestSettingsInterface []interface{}
)

func init() {
	testForestSettingsConf = []*ForestSettings{
		{
			Password:          "test-password",
			ForestName:        "test-forest-name",
			GcServer:          "test1-server.com:1234",
			User:              "test-user",
			UseSSL:            true,
			AuthorizationType: "test-authorization-type",
		},
	}
	testForestSettingsInterface = []interface{}{
		map[string]interface{}{
			"password":           "test-password",
			"forest_name":        "test-forest-name",
			"gc_server":          "test1-server.com:1234",
			"user":               "test-user",
			"use_ssl":            true,
			"authorization_type": "test-authorization-type",
		},
	}
}

func TestFlattenSourceForestSettings(t *testing.T) {

	cases := []struct {
		Input          []*ForestSettings
		ExpectedOutput []interface{}
	}{
		{
			testForestSettingsConf,
			testForestSettingsInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceForestSettings(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceForestSettings(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*ForestSettings
	}{
		{
			testForestSettingsInterface,
			testForestSettingsConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceForestSettings(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
