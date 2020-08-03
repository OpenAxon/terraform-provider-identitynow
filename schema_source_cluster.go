package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceClusterFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Id of cluster",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of cluster",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:	"CLUSTER",
			Description: "Type of cluster",
		},
	}

	return s
}
