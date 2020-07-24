package main

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://axon-sb.api.identitynow.com/beta"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURL,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
