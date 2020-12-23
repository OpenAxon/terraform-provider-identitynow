package main

import (
	"reflect"
	"testing"
)

var (
	testGroupSearchDNsConf      []*GroupSearchDNs
	testGroupSearchDNsInterface []interface{}
)

func init() {
	testGroupSearchDNsConf = []*GroupSearchDNs{
		{
			SearchDN:            "test-search-dn",
			SearchScope:         "test-search-scope",
			IterateSearchFilter: "test-iterate-search-filter",
		},
	}
	testGroupSearchDNsInterface = []interface{}{
		map[string]interface{}{
			"search_dn":             "test-search-dn",
			"search_scope":          "test-search-scope",
			"iterate_search_filter": "test-iterate-search-filter",
		},
	}
}

func TestFlattenSourceGroupSearchDNs(t *testing.T) {

	cases := []struct {
		Input          []*GroupSearchDNs
		ExpectedOutput []interface{}
	}{
		{
			testGroupSearchDNsConf,
			testGroupSearchDNsInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceGroupSearchDNs(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceGroupSearchDNs(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*GroupSearchDNs
	}{
		{
			testGroupSearchDNsInterface,
			testGroupSearchDNsConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceGroupSearchDNs(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
