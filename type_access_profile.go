package main

type AccessProfile struct {
	Description         string                   `json:"description"`
	Enabled             *bool                    `json:"enabled,omitempty"`
	Entitlements        []*ObjectInfo            `json:"entitlements,omitempty"`
	ID                  string                   `json:"id,omitempty"`
	Name                string                   `json:"name,omitempty"`
	AccessProfileOwner  *ObjectInfo              `json:"owner,omitempty"`
	AccessProfileSource *ObjectInfo              `json:"source,omitempty"`
	Requestable         *bool                    `json:"requestable,omitempty"`
	AccessRequestConfig *AccessRequestConfigList `json:"accessRequestConfig,omitempty"`
}

type AccessRequestConfigList struct {
	CommentsRequired        *bool       `json:"commentsRequired,omitempty"`
	DenialCommentsRequired  *bool       `json:"denialCommentsRequired,omitempty"`
	ApprovalSchemes         interface{} `json:"approvalSchemes,omitempty"`
	ReauthorizationRequired *bool       `json:"reauthorizationRequired,omitempty"`
}

type UpdateAccessProfile struct {
	Op    string        `json:"op"`
	Path  string        `json:"path"`
	Value []interface{} `json:"value"`
}
