package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func sourceEntitlementFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Source entitlements name",
		},
		"source_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "source id",
		},
		"aggregation_source_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "aggregation source id",
		},
		"source_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "source name",
		},
		"attribute": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "attribute",
		},
		"created_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"deleted_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"direct_permissions": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"display_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"displayable_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_modified_time": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"owner_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"owner_uid": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"privileged": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"schema": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"value": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return s
}
