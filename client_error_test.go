package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestClientMethodsReturnAPIError(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name string
		call func(*Client) error
	}{
		{"create source", func(c *Client) error { _, err := c.CreateSourceRequest(ctx, &Source{}); return err }},
		{"update source", func(c *Client) error { _, err := c.UpdateSource(ctx, &Source{}); return err }},
		{"delete source", func(c *Client) error { return c.DeleteSource(ctx, &Source{}) }},
		{"get access profile", func(c *Client) error { _, err := c.GetAccessProfile(ctx, "id"); return err }},
		{"get source entitlements", func(c *Client) error { _, err := c.GetSourceEntitlements(ctx, "id"); return err }},
		{"get source entitlement", func(c *Client) error { _, err := c.GetSourceEntitlement(ctx, "id", "name"); return err }},
		{"create access profile", func(c *Client) error { _, err := c.CreateAccessProfile(ctx, &AccessProfile{}); return err }},
		{"update access profile", func(c *Client) error { _, err := c.UpdateAccessProfile(ctx, nil, "id"); return err }},
		{"delete access profile", func(c *Client) error { return c.DeleteAccessProfile(ctx, &AccessProfile{}) }},
		{"get role", func(c *Client) error { _, err := c.GetRole(ctx, "id"); return err }},
		{"create role", func(c *Client) error { _, err := c.CreateRole(ctx, &Role{}); return err }},
		{"update role", func(c *Client) error { _, err := c.UpdateRole(ctx, nil, "id"); return err }},
		{"delete role", func(c *Client) error { _, err := c.DeleteRole(ctx, &Role{}); return err }},
		{"get identity", func(c *Client) error { _, err := c.GetIdentity(ctx, "alias"); return err }},
		{"get aggregation schedule", func(c *Client) error { _, err := c.GetAccountAggregationSchedule(ctx, "id"); return err }},
		{"manage aggregation schedule", func(c *Client) error {
			_, err := c.ManageAccountAggregationSchedule(ctx, &AccountAggregationSchedule{CronExpressions: []string{"0 0 * * *"}}, true)
			return err
		}},
		{"get account schema", func(c *Client) error { _, err := c.GetAccountSchema(ctx, "source-id", "schema-id"); return err }},
		{"update account schema", func(c *Client) error { _, err := c.UpdateAccountSchema(ctx, &AccountSchema{}); return err }},
		{"create password policy", func(c *Client) error { _, err := c.CreatePasswordPolicy(ctx, &PasswordPolicy{}); return err }},
		{"update password policy", func(c *Client) error { _, err := c.UpdatePasswordPolicy(ctx, &PasswordPolicy{}); return err }},
		{"get password policy", func(c *Client) error { _, err := c.GetPasswordPolicy(ctx, "id"); return err }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := NewClient("https://identitynow.test", "client-id", "client-secret")
			client.HTTPClient.Transport = roundTripFunc(func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(`{"messages":[{"text":"IdentityNow unavailable"}]}`)),
				}, nil
			})
			err := test.call(client)
			if err == nil {
				t.Fatal("expected an API error")
			}
			if !strings.Contains(err.Error(), "IdentityNow unavailable") {
				t.Fatalf("expected IdentityNow API error, got %q", err)
			}
		})
	}
}

func TestDeleteAccountSchemaReturnsTransportError(t *testing.T) {
	wantErr := errors.New("transport unavailable")
	originalTransport := http.DefaultTransport
	http.DefaultTransport = roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, wantErr
	})
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	client := NewClient("https://identitynow.test", "client-id", "client-secret")
	err := client.DeleteAccountSchema(context.Background(), &AccountSchema{})
	if !errors.Is(err, wantErr) {
		t.Fatalf("expected transport error %q, got %v", wantErr, err)
	}
}
