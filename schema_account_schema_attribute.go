package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func accountSchemaAttributeFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"object_type": {
			Type:     schema.TypeString,
			Required: true,
		},

		"source_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"display_attribute": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"identity_attribute": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"managed": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"minable": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"multi": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"entitlement": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
	return s
}
