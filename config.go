package main

import (
	"context"
	"fmt"
)

// Config is the configuration parameters for an IdentityNow API
type Config struct {
	URL          string `json:"url"`
	ClientId     string `json:"cacert"`
	ClientSecret string `json:"tokenKey"`
}

func (c *Config) IdentityNowClient() (*Client, error) {
	client := NewClient(c.URL, c.ClientId, c.ClientSecret)
	ctx := context.Background()

	if err := client.GetToken(ctx); err != nil {
		return nil, err
	}
	if len(client.accessToken) == 0 {
		return nil, fmt.Errorf("access token is empty")
	}
	return client, nil
}
