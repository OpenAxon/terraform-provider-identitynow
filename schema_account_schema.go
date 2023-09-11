package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func accountSchemaFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"source_id": {
			Type:     schema.TypeString,
			Required: true,
		},

		"schema_id": {
			Type:     schema.TypeString,
			Required: true,
		},

		"display_attribute": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"identity_attribute": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"native_object_type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"hierarchy_attribute": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"include_permissions": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"attributes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: accountSchemaAttributesFields(),
			},
		},

		"modified": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"created": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}
