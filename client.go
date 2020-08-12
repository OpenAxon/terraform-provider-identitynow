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

type Client struct {
	BaseURL      string
	clientId     string
	clientSecret string
	accessToken  string
	HTTPClient   *http.Client
}

type errorResponse struct {
	DetailCode    	string    `json:"detailCode"`
	Messages	[]struct {
		Locale string `json:"locale"`
		LocaleOrigen   string `json:"localeOrigin"`
		Text string `json:"text"`
	} `json:"messages"`
}

func NewClient(baseURL string, clientId string, secret string) *Client {
	return &Client{
		BaseURL:      baseURL,
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

func (c *Client) GetSource(ctx context.Context, id string) (*Source, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Source{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateSource(ctx context.Context, source *Source) (*Source, error) {
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

	res := Source{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed source creation response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateSource(ctx context.Context, source *Source) (*Source, error) {
	body, err := json.Marshal(&source)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, source.ID), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Source{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed source update response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteSource(ctx context.Context, source *Source) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, source.ID), nil)
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return err
	}

	req = req.WithContext(ctx)

	var res interface{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed source update response:%+v\n", res)
		log.Fatal(err)
		return err
	}

	return nil
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

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err == nil {
			if res.StatusCode == http.StatusNotFound {
				// on the return statement, an interface value of type error is created by the compiler and bound to the pointer to satisfy the return argument.
				return &NotFoundError{errRes.Messages[0].Text}
			} else {
				return errors.New(errRes.Messages[0].Text)
			}
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Printf("Decoder error:%s", err)
		return err
	}

	return nil
}