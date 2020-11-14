package main

import (
	"reflect"
	"testing"
)

var (
	testConnectorAttributesConf      *ConnectorAttributes
	testConnectorAttributesInterface []interface{}
)

func init() {
	testConnectorAttributesConf = &ConnectorAttributes{
		GrantType:                "grant-type-value",
		ClientID:                 "client-id-value",
		ClientSecret:             "client-secret-value",
		DomainName:               "domain-name-value",
		CloudExternalID:          "cloud-external-value",
		MsGraphResourceBase:      "ms-graph-resource-base-value",
		MsGraphTokenBase:         "ms-graph-token-base-value",
		AzureADGraphResourceBase: "azure-ad-graph-resource-base-value",
		AzureADGraphTokenBase:    "azure-ad-graph-token-base",
		IQServicePort:            "1234",
		IQServicePassword:        "iq-service-password-value",
		IQServiceUser:            "iq-service-user-value",
		IQServiceHost:            "iq-service-host.com",
		UseTLSForIQService:       true,
		AuthorizationType:        "simple",
		ApiVersion:               "1.6",
		ExcludeAWSAccountIdList:  "123456789,987654321",
		IncludeAWSAccountIdList:  "234567890",
		Kid:                      "kid",
		Secret:                   "secret",
		RoleName:                 "role-name-123",
		ManageAllAccountsIAMData: true,
		ConnectorClass:           "openconnector.connector.aws.AWSConnectorSDK",
		Encrypted:                "test",
	}
	testConnectorAttributesInterface = []interface{}{
		map[string]interface{}{
			"grant_type":                   "grant-type-value",
			"client_id":                    "client-id-value",
			"client_secret":                "client-secret-value",
			"domain_name":                  "domain-name-value",
			"cloud_external_id":            "cloud-external-value",
			"ms_graph_resource_base":       "ms-graph-resource-base-value",
			"ms_graph_token_base":          "ms-graph-token-base-value",
			"azure_ad_graph_resource_base": "azure-ad-graph-resource-base-value",
			"azure_ad_graph_token_base":    "azure-ad-graph-token-base",
			"use_tls_for_iq_service":       true,
			"iq_service_port":              "1234",
			"iq_service_password":          "iq-service-password-value",
			"iq_service_user":              "iq-service-user-value",
			"iq_service_host":              "iq-service-host.com",
			"authorization_type":           "simple",
			"api_version":                  "1.6",
			"exclude_aws_account_id_list":  "123456789,987654321",
			"include_aws_account_id_list":  "234567890",
			"kid":                          "kid",
			"secret":                       "secret",
			"role_name":                    "role-name-123",
			"manage_all_accounts_iam_data": true,
			"connector_class":              "openconnector.connector.aws.AWSConnectorSDK",
			"encrypted":                    "test",
		},
	}
}

func TestFlattenSourceConnectorAttributes(t *testing.T) {

	cases := []struct {
		Input          *ConnectorAttributes
		ExpectedOutput []interface{}
	}{
		{
			testConnectorAttributesConf,
			testConnectorAttributesInterface,
		},
	}

	for _, tc := range cases {
		output := flattenSourceConnectorAttributes(tc.Input, []interface{}{})
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandSourceConnectorAttributes(t *testing.T) {
	cases := []struct {
		Input          []interface{}
		ExpectedOutput *ConnectorAttributes
	}{
		{
			testConnectorAttributesInterface,
			testConnectorAttributesConf,
		},
	}

	for _, tc := range cases {
		output := expandSourceConnectorAttributes(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
