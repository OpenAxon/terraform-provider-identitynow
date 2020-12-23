package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceGroupSearchDNFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"search_dn": {
			Type:     schema.TypeString,
			Required: true,
		},
		"search_scope": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"iterate_search_filter": {
			Type:     schema.TypeString,
			Required: true,
		},
	}

	return s
}
