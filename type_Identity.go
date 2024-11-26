package main

type Identity struct {
	ID                 string              `json:"id,omitempty"`
	Name               string              `json:"name,omitempty"`
	Description        string              `json:"description,omitempty"`
	IsManager          bool                `json:"isManager,omitempty"`
	Alias              string              `json:"alias"`
	EmailAddress       string              `json:"emailAddress,omitempty"`
	IdentityStatus     string              `json:"identityStatus,omitempty"`
	Enabled            bool                `json:"enabled,omitempty"`
	IdentityAttributes *IdentityAttributes `json:"attributes,omitempty"`
}

type IdentityAttributes struct {
	AdpID     string `json:"adpId,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	Phone     string `json:"phone,omitempty"`
	UserType  string `json:"userType,omitempty"`
	UID       string `json:"uid,omitempty"`
	Email     string `json:"email,omitempty"`
	WorkdayId string `json:"workdayId,omitempty"`
}
