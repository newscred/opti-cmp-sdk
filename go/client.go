package opticmp

import (
	"net/http"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://api.cmp.optimizely.com/v3"
	userAgent      = "opti-cmp-sdk-go/" + Version
	defaultTimeout = 30 * time.Second
)

// Client is the Optimizely CMP API client. Endpoint methods are grouped into
// namespaces promoted from the embedded generated set (for example
// client.Campaign, client.Library).
type Client struct {
	baseURL    string
	httpClient *http.Client
	header     http.Header
	auth       *authProvider

	// OAuth exposes the OAuth2 helper methods. It is always set; the methods
	// return an error when the client was not created with WithOAuth.
	OAuth *OAuthMethods

	namespaces // generated: namespace fields + initNamespaces
}

// Option configures a [Client].
type Option func(*Client)

// New creates a client. Without an auth option the client sends unauthenticated
// requests.
func New(opts ...Option) *Client {
	c := &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{Timeout: defaultTimeout},
		header:     http.Header{},
		auth:       &authProvider{},
	}
	c.header.Set("Accept", "application/json")
	c.header.Set("User-Agent", userAgent)

	for _, opt := range opts {
		opt(c)
	}

	c.auth.httpClient = c.httpClient
	c.OAuth = &OAuthMethods{auth: c.auth}
	c.initNamespaces()
	return c
}

// WithBaseURL overrides the API base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) { c.baseURL = strings.TrimRight(baseURL, "/") }
}

// WithHTTPClient sets the underlying HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithHeader sets a default request header sent on every request.
func WithHeader(key, value string) Option {
	return func(c *Client) { c.header.Set(key, value) }
}

// WithToken authenticates with a static access token.
func WithToken(token string) Option {
	return func(c *Client) {
		c.auth.mu.Lock()
		defer c.auth.mu.Unlock()
		c.auth.token = token
		c.auth.oauth = nil
	}
}

// WithOAuth authenticates with OAuth2. The client acquires and refreshes tokens
// as needed.
func WithOAuth(cfg OAuthConfig) Option {
	return func(c *Client) {
		state := newOAuthState(cfg)
		c.auth.mu.Lock()
		defer c.auth.mu.Unlock()
		c.auth.oauth = state
		c.auth.token = ""
	}
}

// Authenticate changes the client's authentication after construction. Pass an
// auth option:
//
//	client.Authenticate(opticmp.WithToken(newToken))
//	client.Authenticate(opticmp.WithOAuth(cfg))
//
// Call it with no argument to clear authentication, so the client sends
// unauthenticated requests. It is safe to call while other goroutines make
// requests. Only auth options (WithToken, WithOAuth) are meaningful here.
func (c *Client) Authenticate(opts ...Option) {
	if len(opts) == 0 {
		c.auth.mu.Lock()
		c.auth.token = ""
		c.auth.oauth = nil
		c.auth.mu.Unlock()
		return
	}
	for _, opt := range opts {
		opt(c)
	}
}
