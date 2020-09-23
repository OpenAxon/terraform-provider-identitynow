package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testAccessProfileConf      *AccessProfile
	testAccessProfileInterface map[string]interface{}
)

func init() {
	TRUE := true
	FALSE := false
	testAccessProfileConf = &AccessProfile{
		Name:                         "test name",
		Description:                  "test Description",
		SourceID:                     1234,
		OwnerID:                      456,
		Entitlements:                 []string{"1234", "456"},
		DeniedCommentsRequired:       &FALSE,
		Disabled:                     &TRUE,
		Protected:                    &FALSE,
		RequestCommentsRequired:      &FALSE,
		ApprovalSchemes:              "test",
		RevokeRequestApprovalSchemes: "test",
	}
	testAccessProfileInterface = map[string]interface{}{
		"name":                            "test name",
		"description":                     "test Description",
		"source_id":                       1234,
		"owner_id":                        456,
		"entitlements":                    []interface{}{"1234", "456"},
		"denied_comments_required":        false,
		"disabled":                        true,
		"protected":                       false,
		"request_comments_required":       false,
		"approval_schemes":                "test",
		"revoke_request_approval_schemes": "test",
	}
}

func TestFlattenAccessProfile(t *testing.T) {
	cases := []struct {
		Input          *AccessProfile
		ExpectedOutput map[string]interface{}
	}{
		{
			testAccessProfileConf,
			testAccessProfileInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, accessProfileFields(), tc.ExpectedOutput)
		err := flattenAccessProfile(output, tc.Input)
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

func TestExpandAccessProfile(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *AccessProfile
	}{
		{
			testAccessProfileInterface,
			testAccessProfileConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, accessProfileFields(), tc.Input)
		output, err := expandAccessProfile(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
