package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
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
	DetailCode string `json:"detailCode"`
	Messages   []struct {
		Locale       string `json:"locale"`
		LocaleOrigen string `json:"localeOrigin"`
		Text         string `json:"text"`
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

	req.Header.Set("Accept", "application/json; charset=utf-8")

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

	req.Header.Set("Accept", "application/json; charset=utf-8")

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

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

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
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, source.ID), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

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

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	var res interface{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed source update response:%+v\n", res)
		log.Fatal(err)
		return err
	}

	return nil
}

func (c *Client) GetAccessProfile(ctx context.Context, id string) (*AccessProfile, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/access-profiles/%s", c.BaseURL, id), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := AccessProfile{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Access Profile get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetSourceEntitlements(ctx context.Context, id string) (*SourceEntitlements, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cc/api/entitlement/list?CISApplicationId=%s", c.BaseURL, id), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := SourceEntitlements{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Source Entitlements get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateAccessProfile(ctx context.Context, accessProfile *AccessProfile) (*AccessProfile, error) {
	body, err := json.Marshal(&accessProfile)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v2/access-profiles", c.BaseURL), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := AccessProfile{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Access Profile creation response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateAccessProfile(ctx context.Context, accessProfile *AccessProfile, id string) (*AccessProfile, error) {
	body, err := json.Marshal(&accessProfile)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v2/access-profiles/%s", c.BaseURL, id), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := AccessProfile{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Access Profile creation response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteAccessProfile(ctx context.Context, accessProfile *AccessProfile) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v2/access-profiles/%s", c.BaseURL, accessProfile.ID), nil)
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	var res interface{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed access profile update response:%+v\n", res)
		log.Fatal(err)
		return err
	}

	return nil
}

func (c *Client) GetRole(ctx context.Context, id string) (*Role, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cc/api/role/get?id=%s", c.BaseURL, id), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := Role{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Role get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateRole(ctx context.Context, role *Role) (*Role, error) {
	body, err := json.Marshal(&role)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cc/api/role/create", c.BaseURL), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("New request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := Role{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed source creation response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateRole(ctx context.Context, role *Role) (*Role, error) {
	body, err := json.Marshal(&role)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cc/api/role/update/?id=%s", c.BaseURL, role.ID), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := Role{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Role updating response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteRole(ctx context.Context, role *Role) (*Role, error) {
	body, err := json.Marshal(&role)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cc/api/role/delete/?id=%s", c.BaseURL, role.ID), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := Role{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Role deletion response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetIdentity(ctx context.Context, alias string) (*Identity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/identities/%s", c.BaseURL, alias), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := Identity{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Identity get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAccountAggregationSchedule(ctx context.Context, id string) (*AccountAggregationSchedule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cc/api/source/getAggregationSchedules/%s", c.BaseURL, id), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := []AccountAggregationSchedule{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Schedule Account Aggregation get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res[0], nil
}

func (c *Client) ManageAccountAggregationSchedule(ctx context.Context, scheduleAggregation *AccountAggregationSchedule, enable bool) (*AccountAggregationSchedule, error) {
	endpoint := fmt.Sprintf("%s/cc/api/source/scheduleAggregation/%s", c.BaseURL, scheduleAggregation.SourceID)
	data := url.Values{}
	data.Set("enable", fmt.Sprintf("%t", enable))
	data.Set("cronExp", scheduleAggregation.CronExpressions[0])
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("New request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	req = req.WithContext(ctx)

	res := AccountAggregationSchedule{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed schedule account aggregation response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetAccountSchemaAttributes(ctx context.Context, sourceId string) (*AccountSchema, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cc/api/source/getAccountSchema/%s", c.BaseURL, sourceId), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := AccountSchema{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Account Schema get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateAccountSchemaAttribute(ctx context.Context, attribute *AccountSchemaAttribute) (*AccountSchemaAttribute, error) {
	endpoint := fmt.Sprintf("%s/cc/api/source/createSchemaAttribute/%s", c.BaseURL, attribute.SourceID)
	data := url.Values{}
	data.Set("name", attribute.Name)
	data.Set("description", attribute.Description)
	data.Set("type", attribute.Type)
	data.Set("objectType", attribute.ObjectType)
	data.Set("displayAttribute", fmt.Sprintf("%t", attribute.DisplayAttribute))
	data.Set("entitlement", fmt.Sprintf("%t", attribute.Entitlement))
	data.Set("identityAttribute", fmt.Sprintf("%t", attribute.IdentityAttribute))
	data.Set("managed", fmt.Sprintf("%t", attribute.Managed))
	data.Set("minable", fmt.Sprintf("%t", attribute.Minable))
	data.Set("multi", fmt.Sprintf("%t", attribute.Multi))

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := AccountSchemaAttribute{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Account Schema Attribute creation. response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateAccountSchemaAttribute(ctx context.Context, attribute *AccountSchemaAttribute) (*AccountSchemaAttribute, error) {
	_, errDelete := c.DeleteAccountSchemaAttribute(ctx, attribute)
	if errDelete != nil {
		return nil, errDelete
	}
	res, errCreate := c.CreateAccountSchemaAttribute(ctx, attribute)
	if errCreate != nil {
		return nil, errDelete
	}
	return res, nil
}

func (c *Client) DeleteAccountSchemaAttribute(ctx context.Context, attribute *AccountSchemaAttribute) (*AccountSchemaAttribute, error) {
	endpoint := fmt.Sprintf("%s/cc/api/source/deleteSchemaAttribute/%s", c.BaseURL, attribute.SourceID)
	data := url.Values{}
	data.Set("names", attribute.Name)
	data.Set("objectType", attribute.ObjectType)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := AccountSchemaAttribute{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Account Schema Attribute deletion. response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreatePasswordPolicy(ctx context.Context, attributes *PasswordPolicy) (*PasswordPolicy, error) {
	endpoint := fmt.Sprintf("%s/cc/api/passwordPolicy/create", c.BaseURL)
	data, _ := setPasswordPolicyUrlValues(attributes)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := PasswordPolicy{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Password Policy creation. response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdatePasswordPolicy(ctx context.Context, attributes *PasswordPolicy) (*PasswordPolicy, error) {
	endpoint := fmt.Sprintf("%s/cc/api/passwordPolicy/set/%s", c.BaseURL, attributes.ID)
	data, _ := setPasswordPolicyUrlValues(attributes)
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	res := PasswordPolicy{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed to update Password Policy. response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetPasswordPolicy(ctx context.Context, passwordPolicyId string) (*PasswordPolicy, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cc/api/passwordPolicy/get/%s", c.BaseURL, passwordPolicyId), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	res := PasswordPolicy{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed to get Password Policy. response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeletePasswordPolicy(ctx context.Context, passwordPolicyId string) error {
	endpoint := fmt.Sprintf("%s/cc/api/passwordPolicy/delete/%s", c.BaseURL, passwordPolicyId)

	req, err := http.NewRequest("POST", endpoint, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "*/*")

	req = req.WithContext(ctx)

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
				return &NotFoundError{errRes.Messages[0].Text}
			} else {
				return errors.New(errRes.Messages[0].Text)
			}
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
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
