package main

type ApplicationConfiguration struct {
	ID                             string `json:"id,omitempty"`
	AccountServiceId               int    `json:"accountServiceId,omitempty"`
	AccountServiceMatchAllAccounts bool   `json:"accountServiceMatchAllAccounts"`
	// AccountServiceName                     string `json:"accountServiceName,omitempty"`
	// AccountServicePolicies                 string `json:"accountServicePolicies,omitempty"`
	// AccountServicePolicyId                 string `json:"accountServicePolicyId,omitempty"`
	// AccountServicePolicyName               string `json:"accountServicePolicyName,omitempty"`
	// AccountServiceUseForPasswordManagement string `json:"accountServiceUseForPasswordManagement,omitempty"`
	LaunchPadEnabled        bool          `json:"launchpadEnabled"`
	ProvisionRequestEnabled bool          `json:"provisionRequestEnabled"`
	AppCenterEnabled        bool          `json:"appCenterEnabled"`
	Icon                    string        `json:"icon,omitempty"`
	AccessProfileIDs        []interface{} `json:"accessProfileIds"`
	// AppType string `json:"appType"`
}
