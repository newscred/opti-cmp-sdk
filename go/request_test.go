package opticmp

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestExpandPath(t *testing.T) {
	got := expandPath("/assets/{asset_id}/fields/{field_id}", map[string]string{
		"asset_id": "a b",
		"field_id": "42",
	})
	want := "/assets/a%20b/fields/42"
	if got != want {
		t.Fatalf("expandPath = %q, want %q", got, want)
	}
}

func TestBuildURL(t *testing.T) {
	c := New(WithBaseURL("https://example.test/v3"))

	q := url.Values{}
	q.Set("offset", "10")
	got := c.buildURL(&request{path: "/assets/{id}", pathParams: map[string]string{"id": "x"}, query: q})
	if got != "https://example.test/v3/assets/x?offset=10" {
		t.Fatalf("buildURL relative = %q", got)
	}

	// An absolute URL (used by pagination) is passed through unchanged.
	abs := c.buildURL(&request{path: "https://api.test/next?offset=20"})
	if abs != "https://api.test/next?offset=20" {
		t.Fatalf("buildURL absolute = %q", abs)
	}
}

func TestDoSuccessAndAuthHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer tok" {
			t.Errorf("Authorization = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"123","title":"hi"}`))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL), WithToken("tok"))
	type asset struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}
	resp, err := do[asset](context.Background(), c, &request{method: "GET", path: "/assets/123"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Data.ID != "123" || resp.Data.Title != "hi" {
		t.Fatalf("decoded = %+v", resp.Data)
	}
	if resp.Status != 200 {
		t.Fatalf("status = %d", resp.Status)
	}
}

func TestDoHTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"message":"bad"}`))
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	_, err := do[any](context.Background(), c, &request{method: "POST", path: "/things", body: map[string]any{"x": 1}})

	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("want *HTTPError, got %T", err)
	}
	if httpErr.Status != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d", httpErr.Status)
	}
	data, _ := httpErr.Data.(map[string]any)
	if data["message"] != "bad" {
		t.Fatalf("error data = %v", httpErr.Data)
	}
	if httpErr.Request == nil || httpErr.Request.Method != "POST" {
		t.Fatalf("request spec = %+v", httpErr.Request)
	}
}

func TestDoNoContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c := New(WithBaseURL(srv.URL))
	resp, err := do[any](context.Background(), c, &request{method: "DELETE", path: "/things/1"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != http.StatusNoContent || resp.Data != nil {
		t.Fatalf("resp = %+v", resp)
	}
}

func TestDoContextCancelled(t *testing.T) {
	c := New(WithBaseURL("https://example.invalid"))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := do[any](ctx, c, &request{method: "GET", path: "/x"})
	if err == nil {
		t.Fatal("want error for cancelled context")
	}
}
