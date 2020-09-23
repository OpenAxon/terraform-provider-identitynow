package main

type Role struct {
	AccessProfileIds             []string `json:"accessProfileIds,omitempty"`
	ApprovalSchemes              string   `json:"approvalSchemes,omitempty"`
	DeniedCommentsRequired       *bool    `json:"deniedCommentsRequired,omitempty"`
	Description                  string   `json:"description"`
	Disabled                     *bool    `json:"disabled,omitempty"`
	DisplayName                  string   `json:"displayName,omitempty"`
	ID                           string   `json:"id,omitempty"`
	IdentityCount                int      `json:"identityCount,omitempty"`
	Name                         string   `json:"name"`
	Owner                        string   `json:"owner"`
	RequestCommentsRequired      *bool    `json:"requestCommentsRequired,omitempty"`
	Requestable                  *bool    `json:"requestable,omitempty"`
	RevokeRequestApprovalSchemes string   `json:"revokeRequestApprovalSchemes,omitempty"`
	Selector                     struct {
		AliasList            []interface{} `json:"aliasList,omitempty"`
		ComplexRoleCriterion interface{}   `json:"complexRoleCriterion,omitempty"`
		EntitlementIds       []interface{} `json:"entitlementIds,omitempty"`
		RuleID               interface{}   `json:"ruleId,omitempty"`
		SourceID             interface{}   `json:"sourceId,omitempty"`
		Type                 string        `json:"type,omitempty"`
		ValueMap             []interface{} `json:"valueMap,omitempty"`
	} `json:"selector,omitempty"`
}
