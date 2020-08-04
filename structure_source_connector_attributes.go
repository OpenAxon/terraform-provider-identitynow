package main

// Flatteners

func flattenSourceConnectorAttributes(in *ConnectorAttributes, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	obj["grant_type"] = in.GrantType
	obj["client_id"] = in.ClientID
	obj["client_secret"] = in.ClientSecret
	obj["domain_name"] = in.DomainName
	obj["cloud_external_id"] = in.CloudExternalID
	obj["ms_graph_resource_base"] = in.MsGraphResourceBase
	obj["ms_graph_token_base"] = in.MsGraphTokenBase
	obj["azure_ad_graph_resource_base"] = in.AzureADGraphResourceBase
	obj["azure_ad_graph_token_base"] = in.AzureADGraphTokenBase

	return []interface{}{obj}

}

// Expanders

func expandSourceConnectorAttributes(p []interface{}) *ConnectorAttributes {
	obj := ConnectorAttributes{}

	if len(p) == 0 || p[0] == nil {
		return &obj
	}
	in := p[0].(map[string]interface{})

	obj.GrantType = in["grant_type"].(string)
	obj.ClientID = in["client_id"].(string)
	obj.ClientSecret = in["client_secret"].(string)
	obj.DomainName = in["domain_name"].(string)
	obj.CloudExternalID = in["cloud_external_id"].(string)
	obj.MsGraphResourceBase = in["ms_graph_resource_base"].(string)
	obj.MsGraphTokenBase = in["ms_graph_token_base"].(string)
	obj.AzureADGraphResourceBase = in["azure_ad_graph_resource_base"].(string)
	obj.AzureADGraphTokenBase = in["azure_ad_graph_token_base"].(string)

	return &obj
}