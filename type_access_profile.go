package main

type AccessProfile struct {
	ApprovalSchemes              interface{} `json:"approvalSchemes,omitempty"`
	DeniedCommentsRequired       *bool       `json:"deniedCommentsRequired,omitempty"`
	Description                  string      `json:"description"`
	Disabled                     *bool       `json:"disabled,omitempty"`
	Entitlements                 []string    `json:"entitlements,omitempty"`
	ID                           string      `json:"id,omitempty"`
	Name                         string      `json:"name,omitempty"`
	OwnerID                      int         `json:"ownerId"`
	Protected                    *bool       `json:"protected,omitempty"`
	RequestCommentsRequired      *bool       `json:"requestCommentsRequired,omitempty"`
	Requestable                  *bool       `json:"requestable,omitempty"`
	RevokeRequestApprovalSchemes interface{} `json:"revokeRequestApprovalSchemes,omitempty"`
	SourceID                     int         `json:"sourceId"`
	SourceName                   string      `json:"sourceName,omitempty"`
}
