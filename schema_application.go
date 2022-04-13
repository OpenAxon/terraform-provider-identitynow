package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func applicationFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Application name",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Application description",
		},
		"app_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "App ID",
		},
	}
}
