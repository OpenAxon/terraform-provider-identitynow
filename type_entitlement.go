package main

type SourceEntitlement struct {
	Attribute              string        `json:"attribute,omitempty"`
	Value                  string        `json:"value,omitempty"`
	Description            interface{}   `json:"description,omitempty"`
	SourceSchemaObjectType string        `json:"sourceSchemaObjectType,omitempty"`
	Privileged             bool          `json:"privileged,omitempty"`
	CloudGoverned          bool          `json:"cloudGoverned,omitempty"`
	Requestable            bool          `json:"requestable,omitempty"`
	Attributes             *Attributes   `json:"attributes,omitempty"`
	Source                 *SourceInfo   `json:"source,omitempty"`
	Owner                  interface{}   `json:"owner,omitempty"`
	DirectPermissions      []interface{} `json:"directPermissions,omitempty"`
	Segments               []interface{} `json:"segments,omitempty"`
	ManuallyUpdatedFields  []interface{} `json:"manuallyUpdatedFields,omitempty"`
	Modified               interface{}   `json:"modified,omitempty"`
	Created                interface{}   `json:"created,omitempty"`
	ID                     string        `json:"id"`
	Name                   string        `json:"name"`
}

type Attributes struct {
	GroupType         string      `json:"groupType,omitempty"`
	SAMAccountName    string      `json:"sAMAccountName,omitempty"`
	ObjectGuid        interface{} `json:"objectguid,omitempty"`
	GroupScope        interface{} `json:"GroupScope,omitempty"`
	Description       interface{} `json:"description,omitempty"`
	ObjectSid         interface{} `json:"objectSid,omitempty"`
	Cn                interface{} `json:"cn,omitempty"`
	MsDSPrincipalName interface{} `json:"msDS-PrincipalName,omitempty"`
}

type SourceInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
