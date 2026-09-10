// Command token authenticates with a static access token.
//
// Set OPTI_CMP_TOKEN to a valid access token, then run:
//
//	go run ./examples/token
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	opticmp "github.com/newscred/opti-cmp-sdk/go"
)

func main() {
	token := os.Getenv("OPTI_CMP_TOKEN")
	if token == "" {
		log.Fatal("set OPTI_CMP_TOKEN to run this example")
	}

	client := opticmp.New(opticmp.WithToken(token))
	ctx := context.Background()

	campaigns, err := client.Campaign.ListCampaigns(ctx, opticmp.ListCampaignsParams{})
	if err != nil {
		log.Fatalf("list campaigns: %v", err)
	}
	fmt.Printf("fetched %d campaigns\n", len(campaigns.Data.Data))
}
