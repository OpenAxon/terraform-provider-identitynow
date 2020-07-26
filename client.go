package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	BaseURL = "https://axon-sb.api.identitynow.com"
)

type Client struct {
	BaseURL      string
	clientId     string
	clientSecret string
	accessToken  string
	HTTPClient   *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewClient(clientId string, secret string) *Client {
	return &Client{
		BaseURL:      BaseURL,
		clientId:     clientId,
		clientSecret: secret,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetToken(ctx context.Context) error {
	log.Printf("%s/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", c.BaseURL, c.clientId, c.clientSecret)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", c.BaseURL, c.clientId, c.clientSecret), nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	res := OauthToken{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("da err:%+v\n", err)
		return err
	}

	c.accessToken = res.AccessToken

	return nil
}

func (c *Client) GetSource(ctx context.Context, id string) (*SourceAAD, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := SourceAAD{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	log.Printf("da full response:%+v\n", res)
	//body, err := ioutil.ReadAll(res.Body)
	//log.Printf("da body:%s", string(body))

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	//fullResponse := successResponse{
	//	Data: v,
	//}
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Printf("Decoder error:%s", err)
		return err
	}
	log.Printf("Final response:%+v\n", v)

	return nil
}