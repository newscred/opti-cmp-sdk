package opticmp

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// OAuth2 default endpoints.
const (
	defaultAuthorizationURL = "https://accounts.cmp.optimizely.com/o/oauth2/v1/auth"
	defaultTokenURL         = "https://accounts.cmp.optimizely.com/o/oauth2/v1/token"
	defaultUserInfoURL      = "https://accounts.cmp.optimizely.com/o/oauth2/v1/userinfo"
)

// expiryBuffer refreshes a token slightly before it expires so an in-flight
// request does not use a token that expires mid-flight.
const expiryBuffer = time.Minute

// Grant types.
const (
	GrantAuthorizationCode = "authorization_code"
	GrantClientCredentials = "client_credentials"
)

// Token is an OAuth2 access token.
type Token struct {
	AccessToken  string
	ExpiresAt    time.Time
	RefreshToken string
}

// OAuthConfig configures OAuth2 authentication. GrantType is either
// [GrantAuthorizationCode] or [GrantClientCredentials]. URLs default to the
// Optimizely CMP endpoints when empty.
type OAuthConfig struct {
	GrantType    string
	ClientID     string
	ClientSecret string
	Scope        string // client_credentials only
	TokenURL     string
	// Token optionally seeds the client with an already-obtained token.
	Token *Token
	// OnTokenRefresh is called after each successful token refresh. Use it to
	// persist the token.
	OnTokenRefresh func(context.Context, *Token) error
}

type oauthState struct {
	cfg   OAuthConfig
	mu    sync.Mutex
	token *Token
}

func newOAuthState(cfg OAuthConfig) *oauthState {
	if cfg.TokenURL == "" {
		cfg.TokenURL = defaultTokenURL
	}
	return &oauthState{cfg: cfg, token: cfg.Token}
}

// ensureToken returns a valid token, refreshing it when missing or near expiry.
func (s *oauthState) ensureToken(ctx context.Context, hc *http.Client) (*Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.token == nil || time.Now().After(s.token.ExpiresAt.Add(-expiryBuffer)) {
		if err := s.refreshLocked(ctx, hc); err != nil {
			return nil, err
		}
	}
	return s.token, nil
}

func (s *oauthState) refreshLocked(ctx context.Context, hc *http.Client) error {
	prev := s.token

	var form url.Values
	if s.cfg.GrantType == GrantAuthorizationCode {
		if prev == nil || prev.RefreshToken == "" {
			return errors.New("opticmp: cannot refresh the access token without a refresh token")
		}
		form = url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {prev.RefreshToken},
			"client_id":     {s.cfg.ClientID},
			"client_secret": {s.cfg.ClientSecret},
		}
	} else {
		form = url.Values{
			"grant_type":    {GrantClientCredentials},
			"client_id":     {s.cfg.ClientID},
			"client_secret": {s.cfg.ClientSecret},
		}
		if s.cfg.Scope != "" {
			form.Set("scope", s.cfg.Scope)
		}
	}

	token, err := requestToken(ctx, hc, s.cfg.TokenURL, form)
	if err != nil {
		return err
	}
	// A refresh grant may not return a new refresh token; keep the previous one.
	if token.RefreshToken == "" && prev != nil {
		token.RefreshToken = prev.RefreshToken
	}
	return s.storeLocked(ctx, token)
}

func (s *oauthState) storeLocked(ctx context.Context, token *Token) error {
	s.token = token
	if s.cfg.OnTokenRefresh != nil {
		return s.cfg.OnTokenRefresh(ctx, token)
	}
	return nil
}

// requestToken performs a token request. On failure it returns an [HTTPError]
// whose Request omits the form body, which carries the client secret and any
// refresh token or code.
func requestToken(ctx context.Context, hc *http.Client, tokenURL string, form url.Values) (*Token, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	spec := &RequestSpec{Method: http.MethodPost, URL: tokenURL}

	resp, err := hc.Do(req)
	if err != nil {
		return nil, &HTTPError{Message: err.Error(), Request: spec}
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var data any
		if len(raw) > 0 && json.Unmarshal(raw, &data) != nil {
			data = string(raw)
		}
		return nil, &HTTPError{Status: resp.StatusCode, Message: resp.Status, Data: data, Header: resp.Header, Request: spec}
	}

	var body struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	return &Token{
		AccessToken:  body.AccessToken,
		ExpiresAt:    time.Now().Add(time.Duration(body.ExpiresIn) * time.Second),
		RefreshToken: body.RefreshToken,
	}, nil
}

// OAuthMethods exposes OAuth2 helpers. Obtain it from Client.OAuth.
type OAuthMethods struct {
	auth *authProvider
}

func (o *OAuthMethods) require() (*oauthState, error) {
	state := o.auth.currentOAuth()
	if state == nil {
		return nil, errors.New("opticmp: OAuth is not configured; create the client with WithOAuth(...)")
	}
	return state, nil
}

// AuthorizationURLOptions configures [OAuthMethods.GetAuthorizationURL].
type AuthorizationURLOptions struct {
	RedirectURI      string
	Scope            string
	State            string
	AuthorizationURL string // defaults to the CMP authorization endpoint
}

// GetAuthorizationURL builds the URL to redirect a user to for the
// authorization-code grant.
func (o *OAuthMethods) GetAuthorizationURL(opts AuthorizationURLOptions) (string, error) {
	state, err := o.require()
	if err != nil {
		return "", err
	}
	base := opts.AuthorizationURL
	if base == "" {
		base = defaultAuthorizationURL
	}
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("client_id", state.cfg.ClientID)
	q.Set("redirect_uri", opts.RedirectURI)
	q.Set("response_type", "code")
	if opts.Scope != "" {
		q.Set("scope", opts.Scope)
	}
	if opts.State != "" {
		q.Set("state", opts.State)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

// ExchangeCodeOptions configures [OAuthMethods.ExchangeCode].
type ExchangeCodeOptions struct {
	Code        string
	RedirectURI string
}

// ExchangeCode exchanges an authorization code for a token and stores it.
func (o *OAuthMethods) ExchangeCode(ctx context.Context, opts ExchangeCodeOptions) (*Token, error) {
	state, err := o.require()
	if err != nil {
		return nil, err
	}
	token, err := requestToken(ctx, o.auth.httpClient, state.cfg.TokenURL, url.Values{
		"grant_type":    {GrantAuthorizationCode},
		"code":          {opts.Code},
		"redirect_uri":  {opts.RedirectURI},
		"client_id":     {state.cfg.ClientID},
		"client_secret": {state.cfg.ClientSecret},
	})
	if err != nil {
		return nil, err
	}
	state.mu.Lock()
	defer state.mu.Unlock()
	if err := state.storeLocked(ctx, token); err != nil {
		return nil, err
	}
	return token, nil
}

// GetClientCredentialsToken requests a token with the client-credentials grant
// and stores it.
func (o *OAuthMethods) GetClientCredentialsToken(ctx context.Context) (*Token, error) {
	state, err := o.require()
	if err != nil {
		return nil, err
	}
	form := url.Values{
		"grant_type":    {GrantClientCredentials},
		"client_id":     {state.cfg.ClientID},
		"client_secret": {state.cfg.ClientSecret},
	}
	if state.cfg.Scope != "" {
		form.Set("scope", state.cfg.Scope)
	}
	token, err := requestToken(ctx, o.auth.httpClient, state.cfg.TokenURL, form)
	if err != nil {
		return nil, err
	}
	state.mu.Lock()
	defer state.mu.Unlock()
	if err := state.storeLocked(ctx, token); err != nil {
		return nil, err
	}
	return token, nil
}

// RefreshToken forces a token refresh and returns the new token.
func (o *OAuthMethods) RefreshToken(ctx context.Context) (*Token, error) {
	state, err := o.require()
	if err != nil {
		return nil, err
	}
	state.mu.Lock()
	defer state.mu.Unlock()
	if err := state.refreshLocked(ctx, o.auth.httpClient); err != nil {
		return nil, err
	}
	return state.token, nil
}

// UserInfoOptions configures [OAuthMethods.GetUserInfo].
type UserInfoOptions struct {
	UserInfoURL string // defaults to the CMP userinfo endpoint
}

// GetUserInfo fetches the OpenID userinfo claims for the current token.
func (o *OAuthMethods) GetUserInfo(ctx context.Context, opts UserInfoOptions) (*Response[map[string]any], error) {
	state, err := o.require()
	if err != nil {
		return nil, err
	}
	token, err := state.ensureToken(ctx, o.auth.httpClient)
	if err != nil {
		return nil, err
	}
	infoURL := opts.UserInfoURL
	if infoURL == "" {
		infoURL = defaultUserInfoURL
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, infoURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	spec := &RequestSpec{Method: http.MethodGet, URL: infoURL, Header: req.Header.Clone()}
	resp, err := o.auth.httpClient.Do(req)
	if err != nil {
		return nil, &HTTPError{Message: err.Error(), Request: spec}
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var data any
		if len(raw) > 0 && json.Unmarshal(raw, &data) != nil {
			data = string(raw)
		}
		return nil, &HTTPError{Status: resp.StatusCode, Message: resp.Status, Data: data, Header: resp.Header, Request: spec}
	}

	out := &Response[map[string]any]{Header: resp.Header, Status: resp.StatusCode, URL: infoURL, raw: raw}
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &out.Data); err != nil {
			return nil, err
		}
	}
	return out, nil
}
