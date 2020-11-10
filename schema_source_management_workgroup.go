package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceManagementWorkgroupFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Id of management group",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of management group",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "GOVERNANCE_GROUP",
			Description: "Type of management group",
		},
	}

	return s
}
