package main

type SourceEntitlements struct {
	SourceID          string   `json:"id,omitempty"`
	EntitlementsCount int      `json:"count,omitempty"`
	Items             []*Items `json:"items,omitempty"`
}

type Items struct {
	SourceID   string `json:"applicationId"`
	SourceName string `json:"applicationName"`
	Attribute  string `json:"attribute,omitempty"`
	Attributes struct {
		DisplayName     string      `json:"displayName,omitempty"`
		GroupTypes      interface{} `json:"groupTypes,omitempty"`
		MailEnabled     bool        `json:"mailEnabled,omitempty"`
		MailNickname    string      `json:"mailNickname,omitempty"`
		Owners          interface{} `json:"owners,omitempty"`
		ProxyAddresses  interface{} `json:"proxyAddresses,omitempty"`
		SecurityEnabled bool        `json:"securityEnabled,omitempty"`
		TeamsEnabled    bool        `json:"teamsEnabled,omitempty"`
	} `json:"attributes,omitempty"`
	CreatedTime       interface{}   `json:"createdTime,omitempty"`
	DeletedTime       interface{}   `json:"deletedTime,omitempty"`
	Description       interface{}   `json:"description,omitempty"`
	DirectPermissions []interface{} `json:"directPermissions,omitempty"`
	DisplayName       interface{}   `json:"displayName,omitempty"`
	DisplayableName   string        `json:"displayableName,omitempty"`
	ID                string        `json:"id,omitempty"`
	LastModifiedTime  interface{}   `json:"lastModifiedTime,omitempty"`
	OwnerID           interface{}   `json:"ownerId,omitempty"`
	OwnerUID          interface{}   `json:"ownerUid,omitempty"`
	Privileged        bool          `json:"privileged,omitempty"`
	Schema            string        `json:"schema,omitempty"`
	Value             string        `json:"value,omitempty"`
}

type SourceEntitlementAggregationResult struct {
	Task struct {
		ID string `json:"id,omitempty"`
	} `json:"task"`
}

type SourceEntitlementAggregationStatus struct {
	CompletionStatus string `json:"completionStatus,omitempty"`
}
