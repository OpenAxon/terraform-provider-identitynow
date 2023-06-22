package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func roleAccessProfilesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Id of role",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of AccessProfile",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "ACCESS_PROFILE",
			Description: "Access Profile Type",
		},
	}

	return s
}
