package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceDomainSettingsFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Service Account password to login to on-prem active directory.",
			Sensitive:   true,
		},
		"user": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Service Account user to login to on-prem active directory.",
			Sensitive:   true,
		},
		"servers": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Active directory servers.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory host port.",
		},
		"forest_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory forest name.",
		},
		"domain_dn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory domain controller.",
		},
		"use_ssl": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Use ssl to connect to Active directory.",
		},
		"authorization_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Active Directory authorization type.",
		},
	}
	return s
}
