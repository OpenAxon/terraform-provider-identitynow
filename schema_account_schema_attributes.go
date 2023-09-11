package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func accountSchemaAttributesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"schema": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: sourceSchemaFields(),
			},
		},

		"is_group": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"is_multi_valued": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"is_entitlement": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
	return s
}
