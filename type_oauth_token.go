package main

type OauthToken struct {
	AccessToken         string `json:"access_token"`
	TokenType           string `json:"token_type"`
	ExpiresIn           int    `json:"expires_in"`
	Scope               string `json:"scope"`
	TenantID            string `json:"tenant_id"`
	Pod                 string `json:"pod"`
	StrongAuthSupported bool   `json:"strong_auth_supported"`
	Org                 string `json:"org"`
	IdentityID          string `json:"identity_id"`
	UserName            string `json:"user_name"`
	StrongAuth          bool   `json:"strong_auth"`
	Jti                 string `json:"jti"`
}