package main

type Identity struct {
	ID                          string      `json:"id,omitempty"`
	Name                        string      `json:"name,omitempty"`
	Description                 string      `json:"description,omitempty"`
	DateCreated                 interface{} `json:"dateCreated,omitempty"`
	LastUpdated                 interface{} `json:"lastUpdated,omitempty"`
	Alias                       string      `json:"alias"`
	Email                       string      `json:"email,omitempty"`
	Status                      string      `json:"status,omitempty"`
	Enabled                     bool        `json:"enabled,omitempty"`
	UID                         interface{} `json:"uid,omitempty"`
	UUID                        string      `json:"uuid,omitempty"`
	Pending                     bool        `json:"pending,omitempty"`
	EncryptionKey               interface{} `json:"encryptionKey,omitempty"`
	EncryptionCheck             interface{} `json:"encryptionCheck,omitempty"`
	PasswordResetSinceLastLogin bool        `json:"passwordResetSinceLastLogin,omitempty"`
	UsageCertAttested           string      `json:"usageCertAttested,omitempty"`
	AltAuthViaIntegrationData   interface{} `json:"altAuthViaIntegrationData,omitempty"`
	KbaAnswers                  interface{} `json:"kbaAnswers,omitempty"`
	DisablePasswordReset        bool        `json:"disablePasswordReset,omitempty"`
	PtaSourceID                 interface{} `json:"ptaSourceId,omitempty"`
	SupportsPasswordPush        bool        `json:"supportsPasswordPush,omitempty"`
	Role                        interface{} `json:"role,omitempty"`
	AltPhone                    interface{} `json:"altPhone,omitempty"`
	AltEmail                    interface{} `json:"altEmail,omitempty"`
	IdentityFlags               interface{} `json:"identityFlags,omitempty"`
	AltAuthVia                  string      `json:"altAuthVia,omitempty"`
	ExternalID                  string      `json:"externalId,omitempty"`
	Phone                       interface{} `json:"phone,omitempty"`
	EmployeeNumber              interface{} `json:"employeeNumber,omitempty"`
	Attributes                  interface{} `json:"attributes,omitempty"`
}
