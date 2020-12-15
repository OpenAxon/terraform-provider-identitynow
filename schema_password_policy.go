package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func passwordPolicyFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"account_id_min_word_length": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Char length that disallow account ID fragments",
			Default:     -1,
		},
		"account_name_min_word_length": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Char length that disallow display name fragments",
			Default:     -1,
		},
		"connected_services": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: passwordPolicyConnectedServicesFields(),
			},
		},
		"date_created": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"default_policy": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Is the password policy default policy?",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Password policy description",
		},
		"enable_password_expiration": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"first_expiration_reminder": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"last_updated": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"max_length": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "password max length",
		},
		"max_repeated_chars": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"min_alpha": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "minimum letters in password",
		},
		"min_character_types": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"min_length": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "minimum password length",
		},
		"min_lower": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "minimum number of lowercase characters in password",
		},
		"min_numeric": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "minimum number in password",
		},
		"min_special": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "minimum number of special characters in password",
		},
		"min_upper": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "minimum number of uppercase characters in password",
		},
		"name": {
			Type:        schema.TypeString,
			ForceNew:    true,
			Required:    true,
			Description: "Password policy name",
		},
		"password_expiration": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  90,
		},
		"require_strong_auth_off_network": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"require_strong_auth_untrusted_geographies": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"require_strong_authn": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"use_account_attributes": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Prevent use of account attributes?",
		},
		"use_dictionary": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Prevent use of words in this site's password dictionary?",
		},
		"use_history": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"use_identity_attributes": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Prevent use of identity attributes?",
		},
		"validate_against_account_id": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Disallow account ID fragments?",
		},
		"validate_against_account_name": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Disallow account name fragments?",
		},
	}
	return s
}
