package main

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func flattenPasswordPolicy(d *schema.ResourceData, in *PasswordPolicy) error {
	if in == nil {
		return nil
	}

	d.SetId(in.ID)
	d.Set("account_id_min_word_length", in.AccountIDMinWordLength)
	d.Set("account_name_min_word_length", in.AccountNameMinWordLength)
	d.Set("default_policy", in.DefaultPolicy)
	d.Set("description", in.Description)
	d.Set("enable_password_expiration", in.EnablePasswordExpiration)
	d.Set("first_expiration_reminder", in.FirstExpirationReminder)
	d.Set("max_length", in.MaxLength)
	d.Set("max_repeated_chars", in.MaxRepeatedChars)
	d.Set("min_alpha", in.MinAlpha)
	d.Set("min_character_types", in.MinCharacterTypes)
	d.Set("min_length", in.MinLength)
	d.Set("min_lower", in.MinLower)
	d.Set("min_numeric", in.MinNumeric)
	d.Set("min_special", in.MinSpecial)
	d.Set("min_upper", in.MinUpper)
	d.Set("name", in.Name)
	d.Set("password_expiration", in.PasswordExpiration)
	d.Set("require_strong_auth_off_network", in.RequireStrongAuthOffNetwork)
	d.Set("require_strong_auth_untrusted_geographies", in.RequireStrongAuthUntrustedGeographies)
	d.Set("require_strong_authn", in.RequireStrongAuthn)
	d.Set("use_account_attributes", in.UseAccountAttributes)
	d.Set("use_dictionary", in.UseDictionary)
	d.Set("use_history", in.UseHistory)
	d.Set("use_identity_attributes", in.UseIdentityAttributes)
	d.Set("validate_against_account_id", in.ValidateAgainstAccountID)
	d.Set("validate_against_account_name", in.ValidateAgainstAccountName)

	if in.ConnectedServices != nil {
		d.Set("connected_services", flattenPasswordPolicyConnectedServices(in.ConnectedServices))
	}
	return nil
}

func expandPasswordPolicy(in *schema.ResourceData) (*PasswordPolicy, error) {
	obj := PasswordPolicy{}
	if in == nil {
		return nil, fmt.Errorf("[ERROR] Expanding Password Policy: Schema Resource data is nil")
	}
	if v := in.Id(); len(v) > 0 {
		obj.ID = v
	}

	obj.Name = in.Get("name").(string)
	obj.Description = in.Get("description").(string)

	if v, ok := in.Get("account_name_min_word_length").(int); ok {
		obj.AccountNameMinWordLength = &v
	}

	if v, ok := in.Get("account_id_min_word_length").(int); ok {
		obj.AccountIDMinWordLength = &v
	}

	if v, ok := in.Get("connected_services").([]interface{}); ok && len(v) > 0 {
		obj.ConnectedServices = expandPasswordPolicyConnectedServices(v)
	}

	if v, ok := in.Get("default_policy").(bool); ok {
		obj.DefaultPolicy = &v
	}

	if v, ok := in.Get("enable_password_expiration").(bool); ok {
		obj.EnablePasswordExpiration = &v
	}

	if v, ok := in.Get("first_expiration_reminder").(int); ok {
		obj.FirstExpirationReminder = &v
	}

	if v, ok := in.Get("max_length").(int); ok {
		obj.MaxLength = &v
	}

	if v, ok := in.Get("max_repeated_chars").(int); ok {
		obj.MaxRepeatedChars = &v
	}

	if v, ok := in.Get("min_alpha").(int); ok {
		obj.MinAlpha = &v
	}

	if v, ok := in.Get("min_character_types").(int); ok {
		obj.MinCharacterTypes = &v
	}

	if v, ok := in.Get("min_length").(int); ok {
		obj.MinLength = &v
	}

	if v, ok := in.Get("min_lower").(int); ok {
		obj.MinLower = &v
	}

	if v, ok := in.Get("min_numeric").(int); ok {
		obj.MinNumeric = &v
	}

	if v, ok := in.Get("min_special").(int); ok {
		obj.MinSpecial = &v
	}

	if v, ok := in.Get("min_upper").(int); ok {
		obj.MinUpper = &v
	}

	if v, ok := in.Get("password_expiration").(int); ok {
		obj.PasswordExpiration = &v
	}

	if v, ok := in.Get("require_strong_auth_off_network").(bool); ok {
		obj.RequireStrongAuthOffNetwork = &v
	}

	if v, ok := in.Get("require_strong_auth_untrusted_geographies").(bool); ok {
		obj.RequireStrongAuthUntrustedGeographies = &v
	}

	if v, ok := in.Get("require_strong_authn").(bool); ok {
		obj.RequireStrongAuthn = &v
	}

	if v, ok := in.Get("use_account_attributes").(bool); ok {
		obj.UseAccountAttributes = &v
	}

	if v, ok := in.Get("use_dictionary").(bool); ok {
		obj.UseDictionary = &v
	}

	if v, ok := in.Get("use_history").(int); ok {
		obj.UseHistory = &v
	}

	if v, ok := in.Get("use_identity_attributes").(bool); ok {
		obj.UseIdentityAttributes = &v
	}

	if v, ok := in.Get("validate_against_account_id").(bool); ok {
		obj.ValidateAgainstAccountID = &v
	}

	if v, ok := in.Get("validate_against_account_name").(bool); ok {
		obj.ValidateAgainstAccountName = &v
	}

	return &obj, nil
}
