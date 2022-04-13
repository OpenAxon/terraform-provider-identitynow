package main

type Application struct {
	ID          string `json:"id,omitempty"`
	AppID       string `json:"appId,omitempty"`
	Alias       string `json:"alias,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
