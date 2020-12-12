package main

type PasswordPolicy struct {
	AccountIDMinWordLength                *int                 `json:"accountIdMinWordLength,omitempty"`
	AccountNameMinWordLength              *int                 `json:"accountNameMinWordLength,omitempty"`
	ConnectedServices                     []*ConnectedServices `json:"connectedServices,omitempty"`
	DateCreated                           string               `json:"dateCreated,omitempty"`
	DefaultPolicy                         *bool                `json:"defaultPolicy,omitempty"`
	Description                           string               `json:"description,omitempty"`
	EnablePasswordExpiration              *bool                `json:"enablePasswdExpiration,omitempty"`
	FirstExpirationReminder               *int                 `json:"firstExpirationReminder,omitempty"`
	ID                                    string               `json:"id,omitempty"`
	LastUpdated                           string               `json:"lastUpdated,omitempty"`
	MaxLength                             *int                 `json:"maxLength,omitempty"`
	MaxRepeatedChars                      *int                 `json:"maxRepeatedChars,omitempty"`
	MinAlpha                              *int                 `json:"minAlpha,omitempty"`
	MinCharacterTypes                     *int                 `json:"minCharacterTypes,omitempty"`
	MinLength                             *int                 `json:"minLength,omitempty"`
	MinLower                              *int                 `json:"minLower,omitempty"`
	MinNumeric                            *int                 `json:"minNumeric,omitempty"`
	MinSpecial                            *int                 `json:"minSpecial,omitempty"`
	MinUpper                              *int                 `json:"minUpper,omitempty"`
	Name                                  string               `json:"name"`
	PasswordExpiration                    *int                 `json:"passwordExpiration,omitempty"`
	RequireStrongAuthOffNetwork           *bool                `json:"requireStrongAuthOffNetwork,omitempty"`
	RequireStrongAuthUntrustedGeographies *bool                `json:"requireStrongAuthUntrustedGeographies,omitempty"`
	RequireStrongAuthn                    *bool                `json:"requireStrongAuthn,omitempty"`
	UseAccountAttributes                  *bool                `json:"useAccountAttributes,omitempty"`
	UseDictionary                         *bool                `json:"useDictionary,omitempty"`
	UseHistory                            *int                 `json:"useHistory,omitempty"`
	UseIdentityAttributes                 *bool                `json:"useIdentityAttributes,omitempty"`
	ValidateAgainstAccountID              *bool                `json:"validateAgainstAccountId,omitempty"`
	ValidateAgainstAccountName            *bool                `json:"validateAgainstAccountName,omitempty"`
}

type ConnectedServices struct {
	ID                      string `json:"id,omitempty"`
	ExternalID              string `json:"externalId"`
	Name                    string `json:"name"`
	SupportsPasswordSetDate bool   `json:"supportsPasswordSetDate,omitempty"`
	AppCount                int    `json:"appCount,omitempty"`
}
