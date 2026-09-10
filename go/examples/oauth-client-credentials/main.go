// Command oauth-client-credentials shows the OAuth2 client-credentials grant,
// for server-to-server integrations.
//
// Set OPTI_CMP_CLIENT_ID and OPTI_CMP_CLIENT_SECRET, then run:
//
//	go run ./examples/oauth-client-credentials
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	opticmp "github.com/newscred/opti-cmp-sdk/go"
)

func main() {
	clientID := os.Getenv("OPTI_CMP_CLIENT_ID")
	clientSecret := os.Getenv("OPTI_CMP_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("set OPTI_CMP_CLIENT_ID and OPTI_CMP_CLIENT_SECRET to run this example")
	}

	client := opticmp.New(opticmp.WithOAuth(opticmp.OAuthConfig{
		GrantType:    opticmp.GrantClientCredentials,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        "<scopes>",
	}))
	ctx := context.Background()

	token, err := client.OAuth.GetClientCredentialsToken(ctx)
	if err != nil {
		log.Fatalf("get token: %v", err)
	}
	fmt.Println("token expires at:", token.ExpiresAt)

	// The token is attached to requests automatically, and refreshed on expiry.
	campaigns, err := client.Campaign.ListCampaigns(ctx, opticmp.ListCampaignsParams{})
	if err != nil {
		log.Fatalf("list campaigns: %v", err)
	}
	fmt.Printf("fetched %d campaigns\n", len(campaigns.Data.Data))
}
