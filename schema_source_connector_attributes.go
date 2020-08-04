package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Schemas

func sourceConnectorAttributesFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"grant_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Authentication grant type to use for communication to connected system",
		},
		"client_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Client id for the connector client credentials",
		},
		"client_secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Client id for the connector client credentials",
			Sensitive: 	 true,
		},
		"cloud_external_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Cloud external ID (related to the main id?)",
		},
		"domain_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Domain name for the connector client credentials",
		},
		"ms_graph_resource_base": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Base resource URL that is used for Microsoft Graph API REST calls",
		},
		"ms_graph_token_base": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Base token URL that is used to get access token for Microsoft Graph API REST calls",
		},
		"azure_ad_graph_resource_base": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Base resource URL that is used for Azure AD Graph API REST calls",
		},
		"azure_ad_graph_token_base": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Base token URL that is used to get an access token for Azure AD Graph API REST calls",
		},
	}

	return s
}
