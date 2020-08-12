package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceSchemaFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Id of schema",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of schema",
		},
		"type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Type of schema",
		},
	}

	return s
}
