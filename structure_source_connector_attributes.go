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
	obj["iq_service_host"] = in.IQServiceHost
	obj["iq_service_port"] = in.IQServicePort
	obj["use_tls_for_iq_service"] = in.UseTLSForIQService
	obj["iq_service_user"] = in.IQServiceUser
	obj["iq_service_password"] = in.IQServicePassword
	obj["authorization_type"] = in.AuthorizationType
	obj["api_version"] = in.ApiVersion
	obj["exclude_aws_account_id_list"] = in.ExcludeAWSAccountIdList
	obj["include_aws_account_id_list"] = in.IncludeAWSAccountIdList
	obj["kid"] = in.Kid
	obj["secret"] = in.Secret
	obj["role_name"] = in.RoleName
	obj["manage_all_accounts_iam_data"] = in.ManageAllAccountsIAMData
	obj["connector_class"] = in.ConnectorClass
	obj["encrypted"] = in.Encrypted

	if in.DomainSettings != nil {
		obj["domain_settings"] = flattenSourceDomainSettings(in.DomainSettings)
	}

	if in.ForestSettings != nil {
		obj["forest_settings"] = flattenSourceForestSettings(in.ForestSettings)
	}

	if in.SearchDNs != nil {
		obj["search_dns"] = flattenSourceSearchDNs(in.SearchDNs)
	}

	if in.GroupSearchDNs != nil {
		obj["group_search_dns"] = flattenSourceGroupSearchDNs(in.GroupSearchDNs)
	}

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
	obj.IQServiceHost = in["iq_service_host"].(string)
	obj.IQServicePort = in["iq_service_port"].(string)
	obj.UseTLSForIQService = in["use_tls_for_iq_service"].(bool)
	obj.IQServiceUser = in["iq_service_user"].(string)
	obj.IQServicePassword = in["iq_service_password"].(string)
	obj.AuthorizationType = in["authorization_type"].(string)
	obj.ApiVersion = in["api_version"].(string)
	obj.ExcludeAWSAccountIdList = in["exclude_aws_account_id_list"].(string)
	obj.IncludeAWSAccountIdList = in["include_aws_account_id_list"].(string)
	obj.Kid = in["kid"].(string)
	obj.Secret = in["secret"].(string)
	obj.RoleName = in["role_name"].(string)
	obj.ManageAllAccountsIAMData = in["manage_all_accounts_iam_data"].(bool)
	obj.ConnectorClass = in["connector_class"].(string)
	obj.Encrypted = in["encrypted"].(string)

	if v, ok := in["forest_settings"].([]interface{}); ok && len(v) > 0 {
		obj.ForestSettings = expandSourceForestSettings(v)
	}

	if v, ok := in["domain_settings"].([]interface{}); ok && len(v) > 0 {
		obj.DomainSettings = expandSourceDomainSettings(v)
	}

	if v, ok := in["search_dns"].([]interface{}); ok && len(v) > 0 {
		obj.SearchDNs = expandSourceSearchDNs(v)
	}

	if v, ok := in["group_search_dns"].([]interface{}); ok && len(v) > 0 {
		obj.GroupSearchDNs = expandSourceGroupSearchDNs(v)
	}

	return &obj
}
