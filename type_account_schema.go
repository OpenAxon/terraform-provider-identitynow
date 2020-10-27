package main

type AccountSchema struct {
	Attributes        []*AccountSchemaAttribute `json:"attributes,omitempty"`
	DisplayAttribute  string                    `json:"displayAttribute,omitempty"`
	GroupAttribute    string                    `json:"groupAttribute,omitempty"`
	IdentityAttribute string                    `json:"identityAttribute,omitempty"`
	NativeObjectType  string                    `json:"nativeObjectType,omitempty"`
	ObjectType        string                    `json:"objectType,omitempty"`
}

type AccountSchemaAttribute struct {
	ID                string `json:"id,omitempty"`
	Description       string `json:"description,omitempty"`
	DisplayAttribute  bool   `json:"displayAttribute,omitempty"`
	Entitlement       bool   `json:"entitlement,omitempty"`
	IdentityAttribute bool   `json:"identityAttribute,omitempty"`
	Managed           bool   `json:"managed,omitempty"`
	Minable           bool   `json:"minable,omitempty"`
	Multi             bool   `json:"multi,omitempty"`
	Name              string `json:"name"`
	Type              string `json:"type,omitempty"`
	ObjectType        string `json:"objectType"`
	SourceID          string `json:"sourceId,omitempty"`
}
