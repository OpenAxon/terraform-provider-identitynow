package main

import (
	"bytes"
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
	//log.Printf("%s/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", c.BaseURL, c.clientId, c.clientSecret)
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

func (c *Client) CreateSource(ctx context.Context, source *SourceAAD) (*SourceAAD, error) {
	body, err := json.Marshal(&source)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/beta/sources", c.BaseURL), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("New request failed:%+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := SourceAAD{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed source creation response:%+v\n", res)
		log.Fatal(err)
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
		log.Printf("Error After httpclient.do:%+v\n", err)
		return err
	}
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

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Printf("Decoder error:%s", err)
		return err
	}

	return nil
}