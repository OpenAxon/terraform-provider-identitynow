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
	"reflect"
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

func (c *Client) CreateSourceRequest(ctx context.Context, source *Source) (*Source, error) {
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

func (c *Client) AddConnectorAttributesToMicrosoftEntraSource(ctx context.Context, source *Source) (*Source, error) {
	if source == nil || source.ConnectorAttributes == nil {
		return nil, fmt.Errorf("source or ConnectorAttributes cannot be nil")
	}

	var updateSource []*UpdateSource

	// Reflect on the ConnectorAttributes struct
	val := reflect.ValueOf(source.ConnectorAttributes).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldName := getJSONFieldName(field)
		if fieldName == "" { // Skip fields without valid JSON tags
			continue
		}

		fieldValue := val.Field(i).Interface()

		// Skip empty or nil values
		if isEmptyValue(fieldValue) {
			continue
		}

		// Create the update source object
		updateSource = append(updateSource, &UpdateSource{
			Op:    "add",
			Path:  "/connectorAttributes/" + fieldName,
			Value: fieldValue,
		})
	}

	if len(updateSource) == 0 {
		log.Printf("No attributes to update")
		return source, nil // Return the original source if nothing to update
	}

	// Marshal the updateSource to JSON
	body, err := json.MarshalIndent(updateSource, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal updateSource: %v\n", err)
		return nil, fmt.Errorf("failed to marshal updateSource: %w", err)
	}
	log.Printf("updateSource: %s\n", string(body))

	// Create the HTTP PATCH request
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v3/sources/%s", c.BaseURL, source.ID), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to create HTTP request: %v\n", err)
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json-patch+json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req = req.WithContext(ctx)

	// Send the request and handle the response
	var res Source
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed updating source: %v\n", err)
		return nil, fmt.Errorf("failed to update source: %w", err)
	}

	resBody, _ := json.MarshalIndent(res, "", "  ")
	log.Printf("Response Body is: %s\n", string(resBody))

	return &res, nil
}

func (c *Client) CreateSource(ctx context.Context, source *Source) (*Source, error) {
	var res *Source

	if source.Connector == "Microsoft-Entra" {
		newSource := *source
		newSource.ConnectorAttributes = nil

		// Create source request
		sourceResponse, err := c.CreateSourceRequest(ctx, &newSource)
		if err != nil {
			return nil, err
		}
		source.ID = sourceResponse.ID
		// Add connector attributes
		res, err = c.AddConnectorAttributesToMicrosoftEntraSource(ctx, source)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		res, err = c.CreateSourceRequest(ctx, source)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v3/access-profiles/%s", c.BaseURL, id), nil)
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

func (c *Client) GetSourceEntitlements(ctx context.Context, id string) ([]*SourceEntitlement, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/beta/entitlements?filters=source.id", c.BaseURL)+url.QueryEscape(" eq ")+fmt.Sprintf("\"%s\"", id), nil)
	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}

	req = req.WithContext(ctx)

	var res []*SourceEntitlement
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Source Entitlements get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func (c *Client) CreateAccessProfile(ctx context.Context, accessProfile *AccessProfile) (*AccessProfile, error) {
	body, err := json.Marshal(&accessProfile)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v3/access-profiles", c.BaseURL), bytes.NewBuffer(body))
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

func (c *Client) UpdateAccessProfile(ctx context.Context, accessProfile []*UpdateAccessProfile, id interface{}) (*AccessProfile, error) {
	body, err := json.Marshal(&accessProfile)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v3/access-profiles/%s", c.BaseURL, id), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-patch+json; charset=utf-8")
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v3/access-profiles/%s", c.BaseURL, accessProfile.ID), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v3/roles/%s", c.BaseURL, id), nil)
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

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v3/roles", c.BaseURL), bytes.NewBuffer(body))
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

func (c *Client) UpdateRole(ctx context.Context, role []*UpdateRole, id interface{}) (*Role, error) {
	body, err := json.Marshal(&role)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v3/roles/%s", c.BaseURL, id), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Creation of new http request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-patch+json; charset=utf-8")
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v3/role/%s", c.BaseURL, role.ID), bytes.NewBuffer(body))
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

func (c *Client) GetIdentity(ctx context.Context, alias string) ([]*Identity, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/beta/identities?filters=alias", c.BaseURL)+url.QueryEscape(" eq ")+fmt.Sprintf("\"%s\"", alias), nil)

	if err != nil {
		log.Printf("Creation of new http request failed: %+v\n", err)
		return nil, err
	}
	log.Printf("GetIdentity Request is: %+v\n", req)

	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)

	var res []*Identity
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Identity get response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	log.Printf("GetIdentity Response is: %+v\n", res)

	return res, nil
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

func (c *Client) GetAccountSchema(ctx context.Context, sourceId string, id string) (*AccountSchema, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v3/sources/%s/schemas/%s", c.BaseURL, sourceId, id), nil)
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
	res.SourceID = sourceId

	return &res, nil
}

//func (c *Client) CreateAccountSchema(ctx context.Context, accountSchema *AccountSchema) (*AccountSchema, error) {
//for _, value := range updateAccountSchema {
//	log.Printf("arrBody: %+v, value: %+v", value, value.Value)
//}
//log.Printf("arrBody type: %+v", reflect.TypeOf(updateAccountSchema))
//body, err := json.Marshal(&updateAccountSchema)
//log.Printf("body: %+v", string(body))
//
//if err != nil {
//	return nil, err
//}
//req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v3/sources/%s/schemas/%s", c.BaseURL, sourceId, schemaId), bytes.NewBuffer(body))
//if err != nil {
//	log.Printf("New request failed:%+v\n", err)
//	return nil, err
//}
//
//req.Header.Set("Content-Type", "application/json-patch+json; charset=utf-8")
//req.Header.Set("Accept", "application/json; charset=utf-8")
//
//req = req.WithContext(ctx)
//res := AccountSchema{}
//if err := c.sendRequest(req, &res); err != nil {
//	log.Printf("get body: %+v\n", req.GetBody)
//
//	log.Printf("Failed Account Schema Attribute creation. response:%+v\n", res)
//	log.Fatal(err)
//	return nil, err
//}
//for _, value := range updateAccountSchema {
//	log.Printf("arrBody: %+v, value: %+v", value, value.Value)
//}
//return &res, nil
//}

func (c *Client) UpdateAccountSchema(ctx context.Context, accountSchema *AccountSchema) (*AccountSchema, error) {
	body, err := json.Marshal(&accountSchema)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/v3/sources/%s/schemas/%s", c.BaseURL, accountSchema.SourceID, accountSchema.ID), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("New request failed:%+v\n", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)
	res := AccountSchema{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed Account Schema Attribute updating. response:%+v\n", res)
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteAccountSchema(ctx context.Context, accountSchema *AccountSchema) error {
	endpoint := fmt.Sprintf("%s/v3/sources/%s/schemas/%s", c.BaseURL, accountSchema.SourceID, accountSchema.ID)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", endpoint, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	req = req.WithContext(ctx)
	res, err := client.Do(req)

	if err != nil {
		log.Printf("Failed Account Schema Attribute deletion. response:%+v\n", res)
		log.Fatal(err)
		return err
	}

	return nil
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
			if len(errRes.Messages) == 0 {
				return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
			}
			if res.StatusCode == http.StatusNotFound {
				// on the return statement, an interface value of type error is created by the compiler and bound to the pointer to satisfy the return argument.
				return &NotFoundError{errRes.Messages[0].Text}
			}
			return errors.New(errRes.Messages[0].Text)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Printf("Decoder error:%s", err)
		return err
	}

	return nil
}
