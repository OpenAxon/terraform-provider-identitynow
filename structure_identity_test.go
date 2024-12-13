package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	reflect "reflect"
	"testing"
)

var (
	testIdentityConf      *Identity
	testIdentityInterface map[string]interface{}
)

func init() {
	testIdentityConf = &Identity{
		Alias:          "test_alias",
		Name:           "Test Name",
		Description:    "test description",
		EmailAddress:   "test@email.com",
		IdentityStatus: "ACTIVE",
		Enabled:        true,
		IsManager:      false,
		IdentityAttributes: &IdentityAttributes{
			AdpID:     "12345",
			LastName:  "Name",
			FirstName: "Test",
			Phone:     "+11234567890",
			UserType:  "Employee",
			UID:       "tname",
			Email:     "test@email.com",
			WorkdayId: "567890",
		},
	}
	testIdentityInterface = map[string]interface{}{
		"alias":           "test_alias",
		"name":            "Test Name",
		"description":     "test description",
		"email_address":   "test@email.com",
		"identity_status": "ACTIVE",
		"enabled":         true,
		"is_manager":      false,
		"attributes": []interface{}{
			map[string]interface{}{
				"adp_id":     "12345",
				"lastname":   "Name",
				"firstname":  "Test",
				"phone":      "11234567890",
				"user_type":  "Employee",
				"uid":        "tname",
				"email":      "test@email.com",
				"workday_id": "567890",
			},
		},
	}
}

func TestFlattenIdentity(t *testing.T) {
	cases := []struct {
		Input          *Identity
		ExpectedOutput map[string]interface{}
	}{
		{
			testIdentityConf,
			testIdentityInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, identityFields(), tc.ExpectedOutput)
		err := flattenIdentity(output, tc.Input)
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
