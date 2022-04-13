package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type header struct {
	Key   string
	Value string
}

var (
	headerContentTypeFormURLEncoded = header{"Content-Type", "application/x-www-form-urlencoded; charset=utf-8"}
	headerContentTypeJson           = header{"Content-Type", "application/json; charset=utf-8"}
	headerAcceptJson                = header{"Accept", "application/json; charset=utf-8"}
	headerAcceptAll                 = header{"Accept", "*/*"}
)

type Client struct {
	BaseURL      string
	clientId     string
	clientSecret string
	accessToken  string
	HTTPClient   *http.Client
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
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", c.BaseURL, c.clientId, c.clientSecret),
		nil,
		headerAcceptJson)
	if err != nil {
		return err
	}

	res := OauthToken{}
	if err := c.sendRequest(req, &res); err != nil {
		return err
	}

	c.accessToken = res.AccessToken
	return nil
}

func (c *Client) GetSource(ctx context.Context, id string) (*Source, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, id),
		nil,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Source{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) CreateSource(ctx context.Context, source *Source) (*Source, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/beta/sources", c.BaseURL),
		&source,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Source{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) UpdateSource(ctx context.Context, source *Source) (*Source, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, source.ID),
		&source,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Source{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) DeleteSource(ctx context.Context, source *Source) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/beta/sources/%s", c.BaseURL, source.ID),
		nil,
		headerAcceptJson)
	if err != nil {
		return err
	}

	return c.sendRequest(req, nil)
}

func (c *Client) GetAccessProfile(ctx context.Context, id string) (*AccessProfile, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/access-profiles/%s", c.BaseURL, id),
		nil)
	if err != nil {
		return nil, err
	}

	res := AccessProfile{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetSourceEntitlements(ctx context.Context, id string) (*SourceEntitlements, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/entitlement/list?CISApplicationId=%s", c.BaseURL, id),
		nil)
	if err != nil {
		return nil, err
	}

	res := SourceEntitlements{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) StartSourceEntitlementAggregation(ctx context.Context, id string) (*SourceEntitlementAggregationResult, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/source/loadEntitlements/%s", c.BaseURL, id),
		nil,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := SourceEntitlementAggregationResult{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetSourceEntitlementAggregationStatus(ctx context.Context, id string) (*SourceEntitlementAggregationStatus, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/taskResult/get/%s", c.BaseURL, id),
		nil,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := SourceEntitlementAggregationStatus{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) CreateApplication(ctx context.Context, application *Application) (*Application, error) {
	data := url.Values{}
	data.Set("name", application.Name)
	data.Set("description", application.Description)

	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/app/create/", c.BaseURL),
		strings.NewReader(data.Encode()),
		headerContentTypeFormURLEncoded,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := &Application{}
	return res, c.sendRequest(req, res)
}

func (c *Client) GetApplication(ctx context.Context, id string, v interface{}) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/app/get/%s", c.BaseURL, id),
		nil,
		headerAcceptJson)
	if err != nil {
		return err
	}
	return c.sendRequest(req, &v)
}

func (c *Client) UpdateApplication(ctx context.Context, id string, v interface{}) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/app/update/%s", c.BaseURL, id),
		v,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return err
	}
	return c.sendRequest(req, nil)
}

func (c *Client) DeleteApplication(ctx context.Context, id string) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/app/delete/%s", c.BaseURL, id),
		nil,
		headerAcceptAll)
	if err != nil {
		return err
	}
	return c.sendRequest(req, nil)
}

func (c *Client) CreateAccessProfile(ctx context.Context, accessProfile *AccessProfile) (*AccessProfile, error) {
	req, err := c.newHttpRequest(ctx,
		http.MethodPost,
		fmt.Sprintf("%s/v2/access-profiles", c.BaseURL),
		&accessProfile,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := AccessProfile{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) UpdateAccessProfile(ctx context.Context, accessProfile *AccessProfile, id string) (*AccessProfile, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("%s/v2/access-profiles/%s", c.BaseURL, id),
		&accessProfile,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := AccessProfile{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) DeleteAccessProfile(ctx context.Context, accessProfile *AccessProfile) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/v2/access-profiles/%s", c.BaseURL, accessProfile.ID),
		nil,
		headerAcceptJson)
	if err != nil {
		return err
	}

	return c.sendRequest(req, nil)
}

func (c *Client) GetRole(ctx context.Context, id string) (*Role, error) {
	req, err := c.newHttpRequest(ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/role/get?id=%s", c.BaseURL, id),
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Role{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) CreateRole(ctx context.Context, role *Role) (*Role, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/role/create", c.BaseURL),
		&role,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Role{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) UpdateRole(ctx context.Context, role *Role) (*Role, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/role/update/?id=%s", c.BaseURL, role.ID),
		&role,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Role{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) DeleteRole(ctx context.Context, role *Role) (*Role, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/role/delete/?id=%s", c.BaseURL, role.ID),
		&role,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Role{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetIdentity(ctx context.Context, alias string) (*Identity, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/identities/%s", c.BaseURL, alias),
		nil,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := Identity{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetAccountAggregationSchedule(ctx context.Context, id string) (*AccountAggregationSchedule, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/source/getAggregationSchedules/%s", c.BaseURL, id),
		nil)
	if err != nil {
		return nil, err
	}

	res := []AccountAggregationSchedule{}
	return &res[0], c.sendRequest(req, &res)
}

func (c *Client) ManageAccountAggregationSchedule(ctx context.Context, scheduleAggregation *AccountAggregationSchedule, enable bool) (*AccountAggregationSchedule, error) {
	data := url.Values{}
	data.Set("enable", fmt.Sprintf("%t", enable))
	data.Set("cronExp", scheduleAggregation.CronExpressions[0])

	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/source/scheduleAggregation/%s", c.BaseURL, scheduleAggregation.SourceID),
		strings.NewReader(data.Encode()),
		headerContentTypeFormURLEncoded)
	if err != nil {
		return nil, err
	}

	res := AccountAggregationSchedule{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetAccountSchemaAttributes(ctx context.Context, sourceId string) (*AccountSchema, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/source/getAccountSchema/%s", c.BaseURL, sourceId),
		nil)
	if err != nil {
		return nil, err
	}

	res := AccountSchema{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) CreateAccountSchemaAttribute(ctx context.Context, attribute *AccountSchemaAttribute) (*AccountSchemaAttribute, error) {
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

	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/source/createSchemaAttribute/%s", c.BaseURL, attribute.SourceID),
		strings.NewReader(data.Encode()),
		headerContentTypeFormURLEncoded,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := AccountSchemaAttribute{}
	return &res, c.sendRequest(req, &res)
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
	data := url.Values{}
	data.Set("names", attribute.Name)
	data.Set("objectType", attribute.ObjectType)

	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/source/deleteSchemaAttribute/%s", c.BaseURL, attribute.SourceID),
		strings.NewReader(data.Encode()),
		headerContentTypeFormURLEncoded,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := AccountSchemaAttribute{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) CreatePasswordPolicy(ctx context.Context, attributes *PasswordPolicy) (*PasswordPolicy, error) {
	data, _ := setPasswordPolicyUrlValues(attributes)
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/passwordPolicy/create", c.BaseURL),
		strings.NewReader(data.Encode()),
		headerContentTypeFormURLEncoded,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := PasswordPolicy{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) UpdatePasswordPolicy(ctx context.Context, attributes *PasswordPolicy) (*PasswordPolicy, error) {
	data, _ := setPasswordPolicyUrlValues(attributes)
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/passwordPolicy/set/%s", c.BaseURL, attributes.ID),
		strings.NewReader(data.Encode()),
		headerContentTypeFormURLEncoded,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := PasswordPolicy{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetPasswordPolicy(ctx context.Context, passwordPolicyId string) (*PasswordPolicy, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/cc/api/passwordPolicy/get/%s", c.BaseURL, passwordPolicyId),
		nil)
	if err != nil {
		return nil, err
	}

	res := PasswordPolicy{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) DeletePasswordPolicy(ctx context.Context, passwordPolicyId string) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cc/api/passwordPolicy/delete/%s", c.BaseURL, passwordPolicyId),
		nil,
		headerAcceptAll)
	if err != nil {
		return err
	}

	return c.sendRequest(req, nil)
}

func (c *Client) CreateGovernanceGroup(ctx context.Context, group GovernanceGroup) (*GovernanceGroup, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/v2/workgroups/", c.BaseURL),
		&group,
		headerContentTypeJson,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := GovernanceGroup{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) GetGovernanceGroup(ctx context.Context, id string) (*GovernanceGroup, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/workgroups/%s", c.BaseURL, id),
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := GovernanceGroup{}
	return &res, c.sendRequest(req, &res)
}

func (c *Client) UpdateGovernanceGroup(ctx context.Context, group GovernanceGroup) (*GovernanceGroup, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("%s/v2/workgroups/%s", c.BaseURL, group.ID),
		&group,
		headerAcceptJson)
	if err != nil {
		return nil, err
	}

	res := GovernanceGroup{}
	if err := c.sendRequest(req, &res); err != nil {
		log.Printf("Failed to update Governance Group. response:%+v\n%+v\n", res, err)
		return nil, err
	}
	return &res, nil
}

func (c *Client) DeleteGovernanceGroup(ctx context.Context, id string) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/v2/workgroups/%s", c.BaseURL, id),
		nil,
		headerAcceptJson)
	if err != nil {
		return err
	}

	return c.sendRequest(req, nil)
}

func (c *Client) UpdateGovernanceGroupMemberships(ctx context.Context, groupID string, request GovernanceGroupMembershipRequest) error {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/v2/workgroups/%s/members", c.BaseURL, groupID),
		&request,
		headerContentTypeJson,
		headerAcceptAll)
	if err != nil {
		return err
	}

	return c.sendRequest(req, nil)
}

func (c *Client) GetGovernanceGroupMembership(ctx context.Context, groupID string) (*GovernanceGroupMembership, error) {
	req, err := c.newHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/v2/workgroups/%s/members", c.BaseURL, groupID),
		nil,
		headerContentTypeJson,
		headerAcceptAll)
	if err != nil {
		return nil, err
	}

	members := []GovernanceGroupMember{}
	err = c.sendRequest(req, &members)
	if err != nil {
		return nil, err
	}

	membership := &GovernanceGroupMembership{
		GroupID:   groupID,
		MemberIDs: []interface{}{},
	}
	for _, member := range members {
		membership.MemberIDs = append(membership.MemberIDs, member.ExternalID)
	}
	return membership, nil
}

func (c *Client) newHttpRequest(
	ctx context.Context,
	Method string,
	URL string,
	Body interface{},
	Headers ...header) (*http.Request, error) {

	log.Printf("[INFO] new http request: %s %s", Method, URL)
	var reader io.Reader

	if Body != nil {
		if suppliedReader, isReader := Body.(io.Reader); isReader {
			reader = suppliedReader
		} else {
			bytearr, err := json.Marshal(Body)
			if err != nil {
				return nil, err
			}
			log.Printf("[INFO] request body: %s", string(bytearr))
			reader = bytes.NewBuffer(bytearr)
		}
	}

	req, err := http.NewRequest(Method, URL, reader)
	if err != nil {
		log.Printf("[ERROR] new http request errror : %+v\n", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	for _, h := range Headers {
		req.Header.Set(h.Key, h.Value)
	}
	return req.WithContext(ctx), err
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] error sending http request:%+v\n", err)
		return err
	}
	defer res.Body.Close()
	responseBytes, _ := ioutil.ReadAll(res.Body)

	log.Printf("[DEBUG] response code = %d, response = %s", res.StatusCode, string(responseBytes))

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		if res.StatusCode == http.StatusNotFound {
			return NotFoundError{string(responseBytes)}
		}
		return errors.New(string(responseBytes))
	}

	if v != nil {
		if err = json.Unmarshal(responseBytes, v); err != nil {
			log.Printf("[ERROR] unmarshalling json from response: %s:%s", err, string(responseBytes))
			return err
		}
	}
	return nil
}
