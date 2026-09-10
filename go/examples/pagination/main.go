// Command pagination iterates through every page of a paginated list endpoint.
//
// Set OPTI_CMP_TOKEN to a valid access token, then run:
//
//	go run ./examples/pagination
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

	resp, err := client.Campaign.ListCampaigns(ctx, opticmp.ListCampaignsParams{})
	if err != nil {
		log.Fatalf("list campaigns: %v", err)
	}

	page := 1
	for {
		fmt.Printf("page %d: %d campaigns\n", page, len(resp.Data.Data))

		if !opticmp.HasNextPage(resp) {
			break
		}
		resp, err = opticmp.GetNextPage(ctx, resp)
		if err != nil {
			log.Fatalf("get next page: %v", err)
		}
		page++
	}

	fmt.Printf("fetched %d page(s)\n", page)
}
