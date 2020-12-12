package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func sourcePasswordPoliciesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Id of password policy",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of password policy",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "IDENTITY",
			Description: "Type of password policy",
		},
	}

	return s
}
