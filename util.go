package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func toArrayInterface(in []string) []interface{} {
	out := make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}

func toArrayString(in []interface{}) []string {
	out := make([]string, len(in))
	for i, v := range in {
		if v == nil {
			out[i] = ""
			continue
		}
		out[i] = v.(string)
	}
	return out
}

func splitAccountSchemaAttributeID(id string) (sourceId string, name string, err error) {
	separator := "-"

	result := strings.Split(id, separator)
	if len(result) == 2 {
		return result[0], result[1], nil
	}
	return "", "", fmt.Errorf("[ERROR Getting source id and name. id: %s", id)
}

func setPasswordPolicyUrlValues(attributes *PasswordPolicy) (url.Values, error) {
	data := url.Values{}
	data.Set("name", attributes.Name)
	data.Set("description", attributes.Description)

	if attributes.AccountIDMinWordLength != nil {
		data.Set("accountIdMinWordLength", fmt.Sprintf("%v", *attributes.AccountIDMinWordLength))
	}
	if attributes.AccountNameMinWordLength != nil {
		data.Set("accountNameMinWordLength", fmt.Sprintf("%v", *attributes.AccountNameMinWordLength))
	}
	if attributes.EnablePasswordExpiration != nil {
		data.Set("enablePasswdExpiration", fmt.Sprintf("%v", *attributes.EnablePasswordExpiration))
	}
	if attributes.FirstExpirationReminder != nil {
		data.Set("firstExpirationReminder", fmt.Sprintf("%v", *attributes.FirstExpirationReminder))
	}
	if attributes.MaxLength != nil {
		data.Set("maxLength", fmt.Sprintf("%v", *attributes.MaxLength))
	}
	if attributes.MaxRepeatedChars != nil {
		data.Set("maxRepeatedChars", fmt.Sprintf("%v", *attributes.MaxRepeatedChars))
	}
	if attributes.MinLength != nil {
		data.Set("minLength", fmt.Sprintf("%v", *attributes.MinLength))
		log.Printf("minLength: %v", *attributes.MinLength)
	}
	if attributes.MinAlpha != nil {
		data.Set("minAlpha", fmt.Sprintf("%v", *attributes.MinAlpha))
		log.Printf("minAlpha: %v", *attributes.MinAlpha)
	}
	if attributes.MinUpper != nil {
		data.Set("minUpper", fmt.Sprintf("%v", *attributes.MinUpper))
	}
	if attributes.MinLower != nil {
		data.Set("minLower", fmt.Sprintf("%v", *attributes.MinLower))
	}
	if attributes.MinNumeric != nil {
		data.Set("minNumeric", fmt.Sprintf("%v", *attributes.MinNumeric))
	}
	if attributes.MinSpecial != nil {
		data.Set("minSpecial", fmt.Sprintf("%v", *attributes.MinSpecial))
	}
	if attributes.MinCharacterTypes != nil {
		data.Set("minCharacterTypes", fmt.Sprintf("%v", *attributes.MinCharacterTypes))
	}

	if attributes.PasswordExpiration != nil {
		data.Set("passwordExpiration", fmt.Sprintf("%v", *attributes.PasswordExpiration))
	}
	if attributes.RequireStrongAuthOffNetwork != nil {
		data.Set("requireStrongAuthOffNetwork", fmt.Sprintf("%v", *attributes.RequireStrongAuthUntrustedGeographies))
	}
	if attributes.RequireStrongAuthUntrustedGeographies != nil {
		data.Set("requireStrongAuthUntrustedGeographies", fmt.Sprintf("%v", *attributes.RequireStrongAuthUntrustedGeographies))
	}
	if attributes.RequireStrongAuthn != nil {
		data.Set("requireStrongAuthn", fmt.Sprintf("%v", *attributes.RequireStrongAuthUntrustedGeographies))
	}
	if attributes.UseAccountAttributes != nil {
		data.Set("useAccountAttributes", fmt.Sprintf("%v", *attributes.UseAccountAttributes))
	}
	if attributes.UseDictionary != nil {
		data.Set("useDictionary", fmt.Sprintf("%v", *attributes.UseDictionary))
	}
	if attributes.UseHistory != nil {
		data.Set("useHistory", fmt.Sprintf("%v", *attributes.UseHistory))
	}
	if attributes.UseIdentityAttributes != nil {
		data.Set("useIdentityAttributes", fmt.Sprintf("%v", *attributes.UseIdentityAttributes))
	}
	if attributes.ValidateAgainstAccountID != nil {
		data.Set("validateAgainstAccountId", fmt.Sprintf("%v", *attributes.ValidateAgainstAccountID))
	}
	if attributes.ValidateAgainstAccountName != nil {
		data.Set("validateAgainstAccountName", fmt.Sprintf("%v", *attributes.ValidateAgainstAccountName))
	}
	return data, nil
}
