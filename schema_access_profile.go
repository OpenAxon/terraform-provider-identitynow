package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func accessProfileFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Access Profile name",
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Access Profile description",
		},

		"source_id": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Source Id that Access Profile is going to create for",
		},

		"source_name": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"owner_id": {
			Type:     schema.TypeInt,
			Required: true,
		},

		"entitlements": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Access Profile Entitlements.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},

		"denied_comments_required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Access Profile Denied Comments Required",
		},

		"approval_schemes": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"disabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"protected": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"request_comments_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"requestable": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"revoke_request_approval_schemes": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}
