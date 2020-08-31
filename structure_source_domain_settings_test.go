package main

import (
	"reflect"
	"testing"
)

var (
	testDomainSettingsConf      []*DomainSettings
	testDomainSettingsInterface []interface{}
)

func init() {
	testDomainSettingsConf = []*DomainSettings{
		{
			Password:          "test-password",
			ForestName:        "test-forest-name",
			Port:              "1234",
			User:              "test-user",
			UseSSL:            true,
			AuthorizationType: "test-authorization-type",
			DomainDN:          "DC=test,DC=domain,DC=com",
			Servers:           []string{"test1-server.com", "test2-server.com"},
		},
	}
	testDomainSettingsInterface = []interface{}{
		map[string]interface{}{
			"password":           "test-password",
			"forest_name":        "test-forest-name",
			"port":               "1234",
			"user":               "test-user",
			"use_ssl":            true,
			"authorization_type": "test-authorization-type",
			"domain_dn":          "DC=test,DC=domain,DC=com",
			"servers":            []interface{}{"test1-server.com", "test2-server.com"},
		},
	}
}

func TestFlattenSourceDomainSettings(t *testing.T) {

	cases := []struct {
		Input          []*DomainSettings
		ExpectedOutput []interface{}
	}{
		{
			testDomainSettingsConf,
			testDomainSettingsInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceDomainSettings(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceDomainSettings(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*DomainSettings
	}{
		{
			testDomainSettingsInterface,
			testDomainSettingsConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceDomainSettings(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
