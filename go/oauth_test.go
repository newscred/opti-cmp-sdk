package opticmp

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// tokenServer serves an OAuth token endpoint that counts requests.
func tokenServer(t *testing.T, calls *int32, body string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(calls, 1)
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		if r.Form.Get("client_id") == "" {
			t.Errorf("missing client_id")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(body))
	}))
}

func TestClientCredentialsFetchAndCache(t *testing.T) {
	var calls int32
	ts := tokenServer(t, &calls, `{"access_token":"abc","expires_in":3600,"token_type":"Bearer"}`)
	defer ts.Close()

	c := New(WithOAuth(OAuthConfig{
		GrantType:    GrantClientCredentials,
		ClientID:     "id",
		ClientSecret: "secret",
		TokenURL:     ts.URL,
	}))

	tok, err := c.OAuth.GetClientCredentialsToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if tok.AccessToken != "abc" {
		t.Fatalf("token = %q", tok.AccessToken)
	}

	// A second auth application reuses the cached token (no new token call).
	req, _ := http.NewRequestWithContext(context.Background(), "GET", "https://x.test", nil)
	if err := c.auth.apply(context.Background(), req); err != nil {
		t.Fatal(err)
	}
	if req.Header.Get("Authorization") != "Bearer abc" {
		t.Fatalf("auth header = %q", req.Header.Get("Authorization"))
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("token endpoint called %d times, want 1", got)
	}
}

func TestRefreshWhenExpired(t *testing.T) {
	var calls int32
	ts := tokenServer(t, &calls, `{"access_token":"fresh","expires_in":3600,"token_type":"Bearer"}`)
	defer ts.Close()

	state := newOAuthState(OAuthConfig{
		GrantType:    GrantClientCredentials,
		ClientID:     "id",
		ClientSecret: "secret",
		TokenURL:     ts.URL,
	})
	// Seed an already-expired token.
	state.token = &Token{AccessToken: "stale", ExpiresAt: time.Now().Add(-time.Hour)}

	tok, err := state.ensureToken(context.Background(), http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}
	if tok.AccessToken != "fresh" {
		t.Fatalf("token = %q, want fresh", tok.AccessToken)
	}
}

func TestRefreshWithinBuffer(t *testing.T) {
	var calls int32
	ts := tokenServer(t, &calls, `{"access_token":"fresh","expires_in":3600,"token_type":"Bearer"}`)
	defer ts.Close()

	state := newOAuthState(OAuthConfig{GrantType: GrantClientCredentials, ClientID: "id", TokenURL: ts.URL})
	// Expires in 30s, which is inside the 60s buffer, so it must refresh.
	state.token = &Token{AccessToken: "stale", ExpiresAt: time.Now().Add(30 * time.Second)}

	tok, _ := state.ensureToken(context.Background(), http.DefaultClient)
	if tok.AccessToken != "fresh" {
		t.Fatalf("token = %q, want refresh within buffer", tok.AccessToken)
	}
}

func TestAuthorizationCodeRefreshRequiresRefreshToken(t *testing.T) {
	state := newOAuthState(OAuthConfig{GrantType: GrantAuthorizationCode, ClientID: "id"})
	state.token = &Token{AccessToken: "stale", ExpiresAt: time.Now().Add(-time.Hour)} // no refresh token
	_, err := state.ensureToken(context.Background(), http.DefaultClient)
	if err == nil || !strings.Contains(err.Error(), "refresh token") {
		t.Fatalf("want refresh-token error, got %v", err)
	}
}

func TestTokenErrorScrubsSecret(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_client"}`))
	}))
	defer ts.Close()

	c := New(WithOAuth(OAuthConfig{GrantType: GrantClientCredentials, ClientID: "id", ClientSecret: "SECRET", TokenURL: ts.URL}))
	_, err := c.OAuth.GetClientCredentialsToken(context.Background())

	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("want *HTTPError, got %T", err)
	}
	if httpErr.Request == nil || httpErr.Request.Body != nil {
		t.Fatalf("token error must not carry a request body (secret leak): %+v", httpErr.Request)
	}
}

func TestOnTokenRefreshFires(t *testing.T) {
	var calls int32
	ts := tokenServer(t, &calls, `{"access_token":"abc","expires_in":3600,"token_type":"Bearer"}`)
	defer ts.Close()

	var got *Token
	c := New(WithOAuth(OAuthConfig{
		GrantType: GrantClientCredentials,
		ClientID:  "id",
		TokenURL:  ts.URL,
		OnTokenRefresh: func(_ context.Context, tok *Token) error {
			got = tok
			return nil
		},
	}))
	if _, err := c.OAuth.GetClientCredentialsToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got == nil || got.AccessToken != "abc" {
		t.Fatalf("onTokenRefresh token = %+v", got)
	}
}

func TestGetAuthorizationURL(t *testing.T) {
	c := New(WithOAuth(OAuthConfig{GrantType: GrantAuthorizationCode, ClientID: "cid"}))
	u, err := c.OAuth.GetAuthorizationURL(AuthorizationURLOptions{
		RedirectURI: "https://app.test/cb",
		Scope:       "read",
		State:       "xyz",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"client_id=cid", "response_type=code", "redirect_uri=https%3A%2F%2Fapp.test%2Fcb", "scope=read", "state=xyz"} {
		if !strings.Contains(u, want) {
			t.Fatalf("authorization URL %q missing %q", u, want)
		}
	}
}

func TestOAuthNotConfigured(t *testing.T) {
	c := New(WithToken("t"))
	if _, err := c.OAuth.GetClientCredentialsToken(context.Background()); err == nil {
		t.Fatal("want error when OAuth not configured")
	}
}

func TestAuthenticateSwitchesAndClears(t *testing.T) {
	authHeader := func(c *Client) string {
		req, _ := http.NewRequestWithContext(context.Background(), "GET", "https://x.test", nil)
		if err := c.auth.apply(context.Background(), req); err != nil {
			t.Fatal(err)
		}
		return req.Header.Get("Authorization")
	}

	c := New(WithToken("first"))
	if got := authHeader(c); got != "Bearer first" {
		t.Fatalf("initial auth = %q", got)
	}

	c.Authenticate(WithToken("second"))
	if got := authHeader(c); got != "Bearer second" {
		t.Fatalf("after switch = %q", got)
	}

	c.Authenticate()
	if got := authHeader(c); got != "" {
		t.Fatalf("after clear = %q, want empty", got)
	}
}

// Run with -race to catch unguarded credential swaps.
func TestAuthenticateConcurrent(t *testing.T) {
	c := New(WithToken("t"))
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			req, _ := http.NewRequestWithContext(context.Background(), "GET", "https://x.test", nil)
			_ = c.auth.apply(context.Background(), req)
		}()
		go func() {
			defer wg.Done()
			c.Authenticate(WithToken("t2"))
		}()
	}
	wg.Wait()
}
