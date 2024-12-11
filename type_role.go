package main

type Role struct {
	Description          string        `json:"description"`
	ID                   string        `json:"id,omitempty"`
	Name                 string        `json:"name"`
	Requestable          *bool         `json:"requestable,omitempty"`
	RoleOwner            *ObjectInfo   `json:"owner,omitempty"`
	AccessProfiles       []*ObjectInfo `json:"accessProfiles,omitempty"`
	LegacyMembershipInfo interface{}   `json:"legacyMembershipInfo,omitempty"`
	Enabled              *bool         `json:"enabled,omitempty"`
	Segments             []interface{} `json:"segments,omitempty"`
	Membership           *struct {
		Type     string `json:"type,omitempty"`
		Criteria struct {
			Operation   string          `json:"operation,omitempty"`
			Key         interface{}     `json:"key,omitempty"`
			StringValue string          `json:"stringValue,omitempty"`
			Children    []*RoleChildren `json:"children,omitempty"`
		} `json:"criteria,omitempty"`
	} `json:"membership,omitempty"`
	AccessRequestConfig struct {
		CommentsRequired       *bool         `json:"commentsRequired,omitempty"`
		DenialCommentsRequired *bool         `json:"denialCommentsRequired,omitempty"`
		ApprovalSchemes        []interface{} `json:"approvalSchemes,omitempty"`
	} `json:"accessRequestConfig,omitempty"`
	RevocationRequestConfig struct {
		ApprovalSchemes []interface{} `json:"approvalSchemes,omitempty"`
	} `json:"revocationRequestConfig,omitempty"`
}
type ObjectInfo struct {
	ID   interface{} `json:"id,omitempty"`
	Type string      `json:"type,omitempty"`
	Name string      `json:"name"`
}

type RoleChildren struct {
	Operation   string        `json:"operation,omitempty"`
	Key         *RoleKey      `json:"key,omitempty"`
	StringValue string        `json:"stringValue,omitempty"`
	Children    *RoleChildren `json:"children,omitempty"`
}

type RoleKey struct {
	Type     string      `json:"type,omitempty"`
	Property interface{} `json:"property,omitempty"`
	SourceId interface{} `json:"sourceId,omitempty"`
}

type UpdateRole struct {
	Op    string        `json:"op"`
	Path  string        `json:"path"`
	Value []interface{} `json:"value"`
}
