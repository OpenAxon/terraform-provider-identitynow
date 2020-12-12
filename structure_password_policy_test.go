package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"reflect"
	"testing"
)

var (
	testPassPolicyConf      *PasswordPolicy
	testPassPolicyInterface map[string]interface{}
)

func init() {
	TRUE := true
	FALSE := false
	num1 := 5
	num2 := 1
	testPassPolicyConf = &PasswordPolicy{
		AccountIDMinWordLength:                &num1,
		AccountNameMinWordLength:              &num1,
		DefaultPolicy:                         &FALSE,
		Description:                           "some description",
		EnablePasswordExpiration:              &FALSE,
		FirstExpirationReminder:               &num1,
		MaxLength:                             &num2,
		MaxRepeatedChars:                      &num2,
		MinAlpha:                              &num2,
		MinCharacterTypes:                     &num2,
		MinLength:                             &num2,
		MinLower:                              &num2,
		MinNumeric:                            &num2,
		MinSpecial:                            &num2,
		MinUpper:                              &num2,
		Name:                                  "some name",
		PasswordExpiration:                    &num1,
		RequireStrongAuthOffNetwork:           &TRUE,
		RequireStrongAuthUntrustedGeographies: &TRUE,
		RequireStrongAuthn:                    &TRUE,
		UseAccountAttributes:                  &TRUE,
		UseDictionary:                         &TRUE,
		UseHistory:                            &num2,
		UseIdentityAttributes:                 &TRUE,
		ValidateAgainstAccountID:              &TRUE,
		ValidateAgainstAccountName:            &FALSE,
	}
	testPassPolicyInterface = map[string]interface{}{
		"account_id_min_word_length":      num1,
		"account_name_min_word_length":    num1,
		"default_policy":                  FALSE,
		"description":                     "some description",
		"enable_password_expiration":      FALSE,
		"first_expiration_reminder":       num1,
		"max_length":                      num2,
		"max_repeated_chars":              num2,
		"min_alpha":                       num2,
		"min_character_types":             num2,
		"min_length":                      num2,
		"min_lower":                       num2,
		"min_numeric":                     num2,
		"min_special":                     num2,
		"min_upper":                       num2,
		"name":                            "some name",
		"password_expiration":             num1,
		"require_strong_auth_off_network": TRUE,
		"require_strong_auth_untrusted_geographies": TRUE,
		"require_strong_authn":                      TRUE,
		"use_account_attributes":                    TRUE,
		"use_dictionary":                            TRUE,
		"use_history":                               num2,
		"use_identity_attributes":                   TRUE,
		"validate_against_account_id":               TRUE,
		"validate_against_account_name":             FALSE,
	}
}

func TestFlattenPasswordPolicy(t *testing.T) {
	cases := []struct {
		Input          *PasswordPolicy
		ExpectedOutput map[string]interface{}
	}{
		{
			testPassPolicyConf,
			testPassPolicyInterface,
		},
	}
	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, passwordPolicyFields(), tc.ExpectedOutput)
		err := flattenPasswordPolicy(output, tc.Input)
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

func TestExpandPasswordPolicy(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *PasswordPolicy
	}{
		{
			testPassPolicyInterface,
			testPassPolicyConf,
		},
	}

	for _, tc := range cases {
		inputResourceData := schema.TestResourceDataRaw(t, passwordPolicyFields(), tc.Input)
		output, err := expandPasswordPolicy(inputResourceData)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
