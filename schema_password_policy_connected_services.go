package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func passwordPolicyConnectedServicesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source id",
		},
		"external_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source external id",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source name",
		},
		"supports_password_set_date": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"app_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
	return s
}
