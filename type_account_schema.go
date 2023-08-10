package main

type AccountSchema struct {
	Attributes         []*AccountSchemaAttribute `json:"attributes,omitempty"`
	DisplayAttribute   string                    `json:"displayAttribute,omitempty"`
	IdentityAttribute  string                    `json:"identityAttribute,omitempty"`
	NativeObjectType   string                    `json:"nativeObjectType,omitempty"`
	Features           []interface{}             `json:"features,omitempty"`
	Configuration      interface{}               `json:"configuration,omitempty"`
	HierarchyAttribute string                    `json:"hierarchyAttribute,omitempty"`
	IncludePermissions bool                      `json:"includePermissions,omitempty"`
	ID                 string                    `json:"id"`
	Name               string                    `json:"name"`
	Created            string                    `json:"created,omitempty"`
	Modified           string                    `json:"modified,omitempty"`
	SourceID           string                    `json:"sourceId,omitempty"`
}

type AccountSchemaAttribute struct {
	Description   string                        `json:"description,omitempty"`
	IsEntitlement bool                          `json:"isEntitlement,omitempty"`
	IsMultiValued bool                          `json:"isMultiValued,omitempty"`
	IsGroup       bool                          `json:"isGroup,omitempty"`
	Name          string                        `json:"name"`
	Type          string                        `json:"type,omitempty"`
	Schema        *AccountSchemaAttributeSchema `json: "schema,omitempty"`
}

type AccountSchemaAttributeSchema struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
