package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateSourceRequestReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/beta/sources" {
			t.Errorf("expected /beta/sources request, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"messages":[{"text":"IdentityNow unavailable"}]}`))
	}))
	t.Cleanup(server.Close)

	client := NewClient(server.URL, "client-id", "client-secret")
	_, err := client.CreateSourceRequest(context.Background(), &Source{})
	if err == nil {
		t.Fatal("expected an API error")
	}
	if !strings.Contains(err.Error(), "IdentityNow unavailable") {
		t.Fatalf("expected IdentityNow API error, got %q", err)
	}
}
