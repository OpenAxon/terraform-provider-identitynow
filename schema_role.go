package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func roleFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"access_profile_ids": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Access Profile Ids.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"approval_schemes": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "required approvers person or group. e.g. manager,owner",
		},
		"denied_comments_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Role description.",
		},
		"disabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"display_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"identity_count": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Role name",
		},
		"owner": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Owner of Role name",
		},
		"request_comments_required": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"requestable": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"revoke_request_approval_schemes": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	return s
}
