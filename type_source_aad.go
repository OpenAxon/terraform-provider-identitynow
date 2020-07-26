package main

import "time"

type SourceAAD struct {
	Description string `json:"description"`
	Owner       *Owner `json:"owner"`
	Cluster struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"cluster"`
	AccountCorrelationConfig struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"accountCorrelationConfig"`
	AccountCorrelationRule    interface{} `json:"accountCorrelationRule"`
	ManagerCorrelationMapping interface{} `json:"managerCorrelationMapping"`
	ManagerCorrelationRule    interface{} `json:"managerCorrelationRule"`
	BeforeProvisioningRule    interface{} `json:"beforeProvisioningRule"`
	Schemas                   []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"schemas"`
	PasswordPolicies []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"passwordPolicies"`
	Features            []string `json:"features"`
	Type                string   `json:"type"`
	Connector           string   `json:"connector"`
	ConnectorClass      string   `json:"connectorClass"`
	ConnectorAttributes struct {
		HealthCheckTimeout      int         `json:"healthCheckTimeout"`
		SupportsDeltaAgg        string      `json:"supportsDeltaAgg"`
		MsGraphResourceBase     string      `json:"msGraphResourceBase"`
		ClientID                string      `json:"clientID"`
		DeltaAggregationEnabled string      `json:"deltaAggregationEnabled"`
		AcctAggregationEnd      int64       `json:"acctAggregationEnd"`
		IQServicePort           string      `json:"IQServicePort"`
		AcctAggregationStart    int64       `json:"acctAggregationStart"`
		PageSize                string      `json:"pageSize"`
		AuthURL                 interface{} `json:"authURL"`
		SubscribedSkus          []struct {
			ConsumedUnits int `json:"consumedUnits"`
			PrepaidUnits  struct {
				Warning   int `json:"warning"`
				Enabled   int `json:"enabled"`
				Suspended int `json:"suspended"`
			} `json:"prepaidUnits"`
			SkuPartNumber    string `json:"skuPartNumber"`
			CapabilityStatus string `json:"capabilityStatus"`
			AppliesTo        string `json:"appliesTo"`
			ServicePlans     []struct {
				ServicePlanName    string `json:"servicePlanName"`
				ProvisioningStatus string `json:"provisioningStatus"`
				AppliesTo          string `json:"appliesTo"`
				ServicePlanID      string `json:"servicePlanId"`
			} `json:"servicePlans"`
			ObjectID string `json:"objectId"`
			SkuID    string `json:"skuId"`
		} `json:"subscribedSkus"`
		UseTLSForIQService          bool        `json:"useTLSForIQService"`
		IQServiceUser               interface{} `json:"IQServiceUser"`
		CloudAuthEnabled            bool        `json:"cloudAuthEnabled"`
		HasFullAggregationCompleted bool        `json:"hasFullAggregationCompleted"`
		MsGraphTokenBase            string      `json:"msGraphTokenBase"`
		DeltaAggregation            interface{} `json:"deltaAggregation"`
		CloudExternalID             string      `json:"cloudExternalId"`
		ClientSecret                string      `json:"clientSecret"`
		SamlRequestBody             interface{} `json:"samlRequestBody"`
		ManageO365Groups            bool        `json:"manageO365Groups"`
		AccountDeltaLink            string      `json:"accountDeltaLink"`
		AzureADGraphTokenBase       string      `json:"azureADGraphTokenBase"`
		DeleteThresholdPercentage   int         `json:"deleteThresholdPercentage"`
		UseForAccounts              string      `json:"useForAccounts"`
		IQServiceHost               interface{} `json:"IQServiceHost"`
		FormPath                    interface{} `json:"formPath"`
		CloudCacheUpdate            int64       `json:"cloudCacheUpdate"`
		TemplateApplication         string      `json:"templateApplication"`
		Encrypted                   string      `json:"encrypted"`
		IsB2CTenant                 bool        `json:"isB2CTenant"`
		DomainName                  string      `json:"domainName"`
		AzureADGraphResourceBase    string      `json:"azureADGraphResourceBase"`
		CloudDisplayName            string      `json:"cloudDisplayName"`
		GrantType                   string      `json:"grantType"`
		BeforeProvisioningRule      interface{} `json:"beforeProvisioningRule"`
		Md5                         string      `json:"md5"`
		Username                    interface{} `json:"username"`
	} `json:"connectorAttributes"`
	DeleteThreshold     int         `json:"deleteThreshold"`
	Authoritative       bool        `json:"authoritative"`
	ManagementWorkgroup interface{} `json:"managementWorkgroup"`
	ID                  string      `json:"id"`
	Name                string      `json:"name"`
	Created             time.Time   `json:"created"`
	Modified            time.Time   `json:"modified"`
}

type Owner       struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}
