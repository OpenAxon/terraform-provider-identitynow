package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceAccountCorrelationConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Id of account correlation config",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of account correlation config",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of account correlation config",
		},
	}

	return s
}
