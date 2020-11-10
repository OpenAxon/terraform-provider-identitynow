package main

import (
	"time"
)

type Source struct {
	Description               string                    `json:"description"`
	Owner                     *Owner                    `json:"owner"`
	Cluster                   *Cluster                  `json:"cluster,omitempty"`
	AccountCorrelationConfig  *AccountCorrelationConfig `json:"accountCorrelationConfig,omitempty"`
	AccountCorrelationRule    interface{}               `json:"accountCorrelationRule,omitempty"`
	ManagerCorrelationMapping interface{}               `json:"managerCorrelationMapping,omitempty"`
	ManagerCorrelationRule    interface{}               `json:"managerCorrelationRule,omitempty"`
	BeforeProvisioningRule    interface{}               `json:"beforeProvisioningRule,omitempty"`
	Schemas                   []*Schema                 `json:"schemas,omitempty"`
	PasswordPolicies          []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"passwordPolicies,omitempty"`
	Features            []string             `json:"features,omitempty"`
	Type                string               `json:"type,omitempty"`
	Connector           string               `json:"connector"`
	ConnectorClass      string               `json:"connectorClass,omitempty"`
	ConnectorAttributes *ConnectorAttributes `json:"connectorAttributes,omitempty"`
	DeleteThreshold     int                  `json:"deleteThreshold"`
	Authoritative       bool                 `json:"authoritative"`
	ManagementWorkgroup *ManagementWorkgroup `json:"managementWorkgroup,omitempty"`
	ID                  string               `json:"id,omitempty"`
	Name                string               `json:"name"`
	Created             time.Time            `json:"created,omitempty"`
	Modified            time.Time            `json:"modified,omitempty"`
}

type Owner struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Cluster struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountCorrelationConfig struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Schema struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ConnectorAttributes struct {
	HealthCheckTimeout      int         `json:"healthCheckTimeout,omitempty"`
	SupportsDeltaAgg        string      `json:"supportsDeltaAgg,omitempty"`
	MsGraphResourceBase     string      `json:"msGraphResourceBase,omitempty"`
	ClientID                string      `json:"clientID,omitempty"`
	DeltaAggregationEnabled string      `json:"deltaAggregationEnabled,omitempty"`
	AcctAggregationEnd      int64       `json:"acctAggregationEnd,omitempty"`
	IQServicePort           string      `json:"IQServicePort,omitempty"`
	AcctAggregationStart    int64       `json:"acctAggregationStart,omitempty"`
	PageSize                string      `json:"pageSize,omitempty"`
	AuthURL                 interface{} `json:"authURL,omitempty"`
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
	} `json:"subscribedSkus,omitempty"`
	UseTLSForIQService          bool              `json:"useTLSForIQService,omitempty"`
	IQServiceUser               interface{}       `json:"IQServiceUser,omitempty"`
	CloudAuthEnabled            bool              `json:"cloudAuthEnabled,omitempty"`
	HasFullAggregationCompleted bool              `json:"hasFullAggregationCompleted,omitempty"`
	MsGraphTokenBase            string            `json:"msGraphTokenBase,omitempty"`
	DeltaAggregation            interface{}       `json:"deltaAggregation,omitempty"`
	CloudExternalID             string            `json:"cloudExternalId,omitempty"`
	ClientSecret                string            `json:"clientSecret,omitempty"`
	SamlRequestBody             interface{}       `json:"samlRequestBody,omitempty"`
	ManageO365Groups            bool              `json:"manageO365Groups,omitempty"`
	AccountDeltaLink            string            `json:"accountDeltaLink,omitempty"`
	AzureADGraphTokenBase       string            `json:"azureADGraphTokenBase,omitempty"`
	DeleteThresholdPercentage   int               `json:"deleteThresholdPercentage,omitempty"`
	UseForAccounts              string            `json:"useForAccounts,omitempty"`
	IQServiceHost               interface{}       `json:"IQServiceHost,omitempty"`
	FormPath                    interface{}       `json:"formPath,omitempty"`
	TemplateApplication         string            `json:"templateApplication,omitempty"`
	Encrypted                   string            `json:"encrypted,omitempty"`
	IsB2CTenant                 bool              `json:"isB2CTenant,omitempty"`
	DomainName                  string            `json:"domainName,omitempty"`
	AzureADGraphResourceBase    string            `json:"azureADGraphResourceBase,omitempty"`
	CloudDisplayName            string            `json:"cloudDisplayName,omitempty"`
	GrantType                   string            `json:"grantType,omitempty"`
	BeforeProvisioningRule      interface{}       `json:"beforeProvisioningRule,omitempty"`
	Md5                         string            `json:"md5,omitempty"`
	Username                    interface{}       `json:"username,omitempty"`
	ForestSettings              []*ForestSettings `json:"forestSettings,omitempty"`
	IQServicePassword           interface{}       `json:"IQServicePassword,omitempty"`
	DomainSettings              []*DomainSettings `json:"domainSettings,omitempty"`
	SearchDNs                   []*SearchDNs      `json:"searchDNs,omitempty"`
	AuthorizationType           string            `json:"authorizationType,omitempty"`
	ApiVersion                  string            `json:"api-version,omitempty"`
}

type ForestSettings struct {
	Password          string      `json:"password,omitempty"`
	GcServer          string      `json:"gcServer,omitempty"`
	ForestName        interface{} `json:"forestName,omitempty"`
	User              string      `json:"user,omitempty"`
	UseSSL            bool        `json:"useSSL,omitempty"`
	AuthorizationType string      `json:"authorizationType,omitempty"`
}

type DomainSettings struct {
	Password          string   `json:"password,omitempty"`
	ForestName        string   `json:"forestName,omitempty"`
	Port              string   `json:"port,omitempty"`
	User              string   `json:"user,omitempty"`
	UseSSL            bool     `json:"useSSL,omitempty"`
	AuthorizationType string   `json:"authorizationType,omitempty"`
	DomainDN          string   `json:"domainDN,omitempty"`
	Servers           []string `json:"servers,omitempty"`
}

type SearchDNs struct {
	GroupMembershipSearchDN string `json:"groupMembershipSearchDN,omitempty"`
	SearchDN                string `json:"searchDN,omitempty"`
	GroupMemberFilterString string `json:"groupMemberFilterString,omitempty"`
	SearchScope             string `json:"searchScope,omitempty"`
	PrimaryGroupSearchDN    string `json:"primaryGroupSearchDN,omitempty"`
	IterateSearchFilter     string `json:"iterateSearchFilter,omitempty"`
}

type ManagementWorkgroup struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
