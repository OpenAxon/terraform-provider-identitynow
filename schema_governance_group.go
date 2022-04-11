package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func governanceGroupFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Governance group name.",
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Governance group description.",
		},
		"owner_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Governance group owner ID.",
		},
		"approval_scheme": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Governance group approval scheme.",
		},
	}
	return s
}
