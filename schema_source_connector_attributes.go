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
			Sensitive:   true,
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
		"iq_service_host": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IQService host url for on-prem Active Directory.",
		},
		"iq_service_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "IQService port for on-prem Active Directory.",
		},
		"use_tls_for_iq_service": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Use TLS for IQService for on-prem Active Directory.",
		},
		"iq_service_user": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Service Account username for IQService host.",
		},
		"iq_service_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Service Account password for IQService host.",
		},
		"forest_settings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: sourceForestSettingsFields(),
			},
		},
		"domain_settings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: sourceDomainSettingsFields(),
			},
		},
		"search_dns": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: sourceSearchDNsFields(),
			},
		},
	}

	return s
}
