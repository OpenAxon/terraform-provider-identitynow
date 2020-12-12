package main

import (
	"reflect"
	"testing"
)

var (
	testConnectedServicesConf      []*ConnectedServices
	testConnectedServicesInterface []interface{}
)

func init() {
	testConnectedServicesConf = []*ConnectedServices{
		{
			ID:                      "123",
			ExternalID:              "123abc",
			Name:                    "some name",
			AppCount:                1,
			SupportsPasswordSetDate: false,
		},
		{
			ID:                      "456",
			ExternalID:              "456abc",
			Name:                    "some other name",
			AppCount:                1,
			SupportsPasswordSetDate: false,
		},
	}
	testConnectedServicesInterface = []interface{}{
		map[string]interface{}{
			"id":                         "123",
			"external_id":                "123abc",
			"name":                       "some name",
			"app_count":                  1,
			"supports_password_set_date": false,
		},
		map[string]interface{}{
			"id":                         "456",
			"external_id":                "456abc",
			"name":                       "some other name",
			"app_count":                  1,
			"supports_password_set_date": false,
		},
	}
}

func TestFlattenPasswordPolicyConnectedServices(t *testing.T) {

	cases := []struct {
		Input          []*ConnectedServices
		ExpectedOutput []interface{}
	}{
		{
			testConnectedServicesConf,
			testConnectedServicesInterface,
		},
	}

	for _, tc := range cases {
		output := flattenPasswordPolicyConnectedServices(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandPasswordPolicyConnectedServices(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*ConnectedServices
	}{
		{
			testConnectedServicesInterface,
			testConnectedServicesConf,
		},
	}

	for _, tc := range cases {
		output := expandPasswordPolicyConnectedServices(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
