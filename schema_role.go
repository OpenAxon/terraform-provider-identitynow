package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func roleFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"description": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Role description.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Role name",
		},
		"owner": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: sourceOwnerFields(),
			},
		},
		"accessProfiles": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: roleAccessProfilesFields(),
			},
		},
		"requestable": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
	return s
}
