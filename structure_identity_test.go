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
		Alias:                       "test_alias",
		Name:                        "Test Name",
		ExternalID:                  "123ab456",
		Description:                 "test description",
		DateCreated:                 "2020-07-15T20:48:26Z",
		LastUpdated:                 "2020-07-15T20:48:26Z",
		Email:                       "test@email.com",
		Status:                      "ACTIVE",
		Enabled:                     true,
		UID:                         "678abc",
		UUID:                        "123-ab12-78ef",
		Pending:                     false,
		EncryptionKey:               nil,
		EncryptionCheck:             nil,
		PasswordResetSinceLastLogin: false,
		UsageCertAttested:           "2020-07-15T20:48:26Z",
		AltAuthViaIntegrationData:   nil,
		KbaAnswers:                  nil,
		DisablePasswordReset:        false,
		PtaSourceID:                 "123",
		SupportsPasswordPush:        false,
		Role:                        false,
		AltPhone:                    "1234567788",
		AltEmail:                    "test2@email.com",
		IdentityFlags:               nil,
		Phone:                       "1234560099",
		EmployeeNumber:              "KG234BG",
		Attributes:                  nil,
	}
	testIdentityInterface = map[string]interface{}{
		"alias":                           "test_alias",
		"name":                            "Test Name",
		"external_id":                     "123ab456",
		"description":                     "test description",
		"date_created":                    "2020-07-15T20:48:26Z",
		"last_updated":                    "2020-07-15T20:48:26Z",
		"email":                           "test@email.com",
		"status":                          "ACTIVE",
		"enabled":                         true,
		"uid":                             "678abc",
		"uuid":                            "123-ab12-78ef",
		"pending":                         false,
		"password_reset_since_last_login": false,
		"usage_cert_attested":             "2020-07-15T20:48:26Z",
		"disable_password_reset":          false,
		"pta_source_id":                   "123",
		"supports_password_push":          false,
		"alt_phone":                       "1234567788",
		"alt_email":                       "test2@email.com",
		"phone":                           "1234560099",
		"employee_number":                 "KG234BG",
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
