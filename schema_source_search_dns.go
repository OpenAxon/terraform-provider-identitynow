package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceSearchDNsFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"search_dn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory search domain criteria.",
		},
		"iterate_search_filter": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory search filter.",
		},
		"group_membership_search_dn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory group membership search criteria.",
		},
		"group_member_filter_string": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory group membership search filter.",
		},
		"search_scope": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory search scope.",
		},
		"primary_group_search_dn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory primary group search.",
		},
	}

	return s
}
