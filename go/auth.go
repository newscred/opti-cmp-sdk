package opticmp

import (
	"context"
	"net/http"
	"sync"
)

// authProvider injects the Authorization header. It holds either a static token
// or OAuth state. mu guards token and oauth so [Client.Authenticate] can swap
// the credentials while requests run on other goroutines.
type authProvider struct {
	mu         sync.RWMutex
	token      string
	oauth      *oauthState
	httpClient *http.Client
}

func (a *authProvider) apply(ctx context.Context, req *http.Request) error {
	a.mu.RLock()
	oauth, token := a.oauth, a.token
	a.mu.RUnlock()

	if oauth != nil {
		t, err := oauth.ensureToken(ctx, a.httpClient)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+t.AccessToken)
		return nil
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return nil
}

// currentOAuth returns the OAuth state under the read lock.
func (a *authProvider) currentOAuth() *oauthState {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.oauth
}
