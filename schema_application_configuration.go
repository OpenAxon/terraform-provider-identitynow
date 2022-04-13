package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func applicationConfigurationFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Associated application ID",
		},
		"icon": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Application icon URL",
		},
		"account_service_id": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Source ID",
		},
		"account_service_match_all_accounts": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Specific Users From Source",
		},
		"access_profile_ids": {
			Type: schema.TypeSet,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional:    true,
			Description: "Associated access profile IDs",
		},
		"app_center_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Visible in the Request Center",
		},
		"launch_pad_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable application for users",
		},
		"provision_request_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Allow access requests",
		},
		"default_values": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Store original configuration when provisioning the resource",
		},
	}
}
