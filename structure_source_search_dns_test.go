package main

import (
	"reflect"
	"testing"
)

var (
	testSearchDNsConf      []*SearchDNs
	testSearchDNsInterface []interface{}
)

func init() {
	testSearchDNsConf = []*SearchDNs{
		{
			SearchDN:                "OU=Users,DC=example,DC=com",
			IterateSearchFilter:     "(objectclass=person)",
			GroupMemberFilterString: "group_member_filter_string_value",
			SearchScope:             "SUBTREE",
			PrimaryGroupSearchDN:    "primary_group_search_dn_value",
			GroupMembershipSearchDN: "group_membership_search_dn_value",
		},
	}
	testSearchDNsInterface = []interface{}{
		map[string]interface{}{
			"search_dn":                  "OU=Users,DC=example,DC=com",
			"iterate_search_filter":      "(objectclass=person)",
			"group_member_filter_string": "group_member_filter_string_value",
			"primary_group_search_dn":    "primary_group_search_dn_value",
			"group_membership_search_dn": "group_membership_search_dn_value",
			"search_scope":               "SUBTREE",
		},
	}
}

func TestFlattenSourceSearchDNs(t *testing.T) {

	cases := []struct {
		Input          []*SearchDNs
		ExpectedOutput []interface{}
	}{
		{
			testSearchDNsConf,
			testSearchDNsInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceSearchDNs(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceSearchDNs(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput []*SearchDNs
	}{
		{
			testSearchDNsInterface,
			testSearchDNsConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceSearchDNs(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
