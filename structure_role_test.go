package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	reflect "reflect"
	"testing"
)

var (
	testRoleConf      *Role
	testRoleInterface map[string]interface{}
)

func init() {
	TRUE := true
	FALSE := false
	testRoleConf = &Role{
		AccessProfileIds:             []string{"1234", "5678"},
		ApprovalSchemes:              "manager",
		DeniedCommentsRequired:       &TRUE,
		Description:                  "test description",
		Disabled:                     &FALSE,
		DisplayName:                  "test name",
		IdentityCount:                1,
		Name:                         "test name",
		Owner:                        "test_identity",
		RequestCommentsRequired:      &FALSE,
		Requestable:                  &FALSE,
		RevokeRequestApprovalSchemes: "test",
	}
	testRoleInterface = map[string]interface{}{
		"access_profile_ids":              []interface{}{"1234", "5678"},
		"approval_schemes":                "manager",
		"denied_comments_required":        true,
		"description":                     "test description",
		"disabled":                        false,
		"display_name":                    "test name",
		"identity_count":                  1,
		"name":                            "test name",
		"owner":                           "test_identity",
		"request_comments_required":       false,
		"requestable":                     false,
		"revoke_request_approval_schemes": "test",
	}
}

func TestFlattenRole(t *testing.T) {
	cases := []struct {
		Input          *Role
		ExpectedOutput map[string]interface{}
	}{
		{
			testRoleConf,
			testRoleInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, roleFields(), tc.ExpectedOutput)
		err := flattenRole(output, tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		expectedOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			expectedOutput[k] = output.Get(k)
		}
		if !reflect.DeepEqual(expectedOutput, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, expectedOutput)
		}
	}
}

func TestExpandRole(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *Role
	}{
		{
			testRoleInterface,
			testRoleConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, roleFields(), tc.Input)
		output, err := expandRole(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
