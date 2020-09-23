package main

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func identityFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"alias": {
			Type:     schema.TypeString,
			Required: true,
		},

		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Identity name",
		},
		"external_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"date_created": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"last_updated": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"email": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"enabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"uid": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"uuid": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"pending": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"encryption_key": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"encryption_check": {
			Type:     schema.TypeBool,
			Optional: true,
		},

		"password_reset_since_last_login": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"usage_cert_attested": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"alt_auth_via_integration_data": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"kba_answers": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"disable_password_reset": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"pta_source_id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"supports_password_push": {
			Type:     schema.TypeBool,
			Computed: true,
		},

		"role": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"alt_phone": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"alt_email": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"identity_flags": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"alt_auth_via": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"phone": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"employee_number": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"attributes": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	return s
}
