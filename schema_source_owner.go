package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceOwnerFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Id of owner",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of owner",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:	"IDENTITY",
			Description: "Type of owner",
		},
	}

	return s
}
