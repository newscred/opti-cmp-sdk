// Command oauth-authorization-code shows the OAuth2 authorization-code grant,
// for user-facing apps.
//
// Running this prints the URL to send a user to. Completing the flow needs an
// HTTP server to receive the redirect, so the second half is shown as the
// handleCallback function you would call from your callback route.
//
// Set OPTI_CMP_CLIENT_ID and OPTI_CMP_CLIENT_SECRET, then run:
//
//	go run ./examples/oauth-authorization-code
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	opticmp "github.com/newscred/opti-cmp-sdk/go"
)

const redirectURI = "https://your.app/callback"

var client *opticmp.Client

func main() {
	clientID := os.Getenv("OPTI_CMP_CLIENT_ID")
	clientSecret := os.Getenv("OPTI_CMP_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("set OPTI_CMP_CLIENT_ID and OPTI_CMP_CLIENT_SECRET to run this example")
	}

	client = opticmp.New(opticmp.WithOAuth(opticmp.OAuthConfig{
		GrantType:    opticmp.GrantAuthorizationCode,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		// Called after each token refresh. Persist the token here so it
		// survives a restart.
		OnTokenRefresh: func(_ context.Context, token *opticmp.Token) error {
			fmt.Println("token refreshed, expires at:", token.ExpiresAt)
			return nil
		},
	}))

	// 1. Send the user here to authorize your app.
	authorizationURL, err := client.OAuth.GetAuthorizationURL(opticmp.AuthorizationURLOptions{
		RedirectURI: redirectURI,
		Scope:       "<scopes>",
		State:       "<csrf-state>",
	})
	if err != nil {
		log.Fatalf("build authorization URL: %v", err)
	}
	fmt.Println("authorize at:", authorizationURL)
}

// handleCallback runs in your callback route. The user is redirected back to
// redirectURI with a code query parameter. Exchange it for a token, then use
// the client as usual.
func handleCallback(ctx context.Context, code string) error {
	token, err := client.OAuth.ExchangeCode(ctx, opticmp.ExchangeCodeOptions{
		Code:        code,
		RedirectURI: redirectURI,
	})
	if err != nil {
		return err
	}
	fmt.Println("token expires at:", token.ExpiresAt)

	userInfo, err := client.OAuth.GetUserInfo(ctx, opticmp.UserInfoOptions{})
	if err != nil {
		return err
	}
	fmt.Println("signed in as:", userInfo.Data)
	return nil
}
