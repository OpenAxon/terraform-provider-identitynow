package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func governanceGroupMembershipFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"group_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Associated governance group ID.",
		},
		"member_ids": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "Member IDs",
			Elem:        schema.TypeString,
		},
	}
	return s
}
