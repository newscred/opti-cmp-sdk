# Optimizely CMP SDK for Go

Go SDK for the [Optimizely CMP](https://www.optimizely.com/products/content-marketing/) (Content Marketing Platform) API.

## Install

```bash
go get github.com/newscred/opti-cmp-sdk/go
```

The package is named `opticmp`. The import path ends in `go`, so alias it for
clarity:

```go
import opticmp "github.com/newscred/opti-cmp-sdk/go"
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	opticmp "github.com/newscred/opti-cmp-sdk/go"
)

func main() {
	client := opticmp.New(opticmp.WithToken("<auth-token>"))
	ctx := context.Background()

	// List campaigns.
	campaigns, err := client.Campaign.ListCampaigns(ctx, opticmp.ListCampaignsParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(campaigns.Data.Data))

	// Get a task.
	task, err := client.Task.GetTask(ctx, opticmp.GetTaskParams{ID: "task-id"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(task.Data.Title)
}
```

Endpoints are grouped into namespaces (`client.Campaign`, `client.Task`,
`client.Library`, `client.Workflow`, and more). Each method takes a `context.Context`
and a params struct, and returns a typed `*opticmp.Response[T]`. Response and request
types live in the `github.com/newscred/opti-cmp-sdk/go/schema` package.

### Optional fields

Optional scalar fields are pointers, so `nil` means "unset" and is distinct from
the zero value. Set them with the generic pointer helper `opticmp.Ptr`:

```go
resp, err := client.Library.ListAssets(ctx, opticmp.ListAssetsParams{
	PageSize: opticmp.Ptr(20),
	Offset:   opticmp.Ptr(40),
})
```

## Authentication

The client supports a static token or full OAuth2.

### Static token

```go
client := opticmp.New(opticmp.WithToken("<auth-token>"))
```

### OAuth2 — authorization code

For user-facing apps. Redirect the user to the authorization URL, then exchange
the returned code for a token.

```go
client := opticmp.New(opticmp.WithOAuth(opticmp.OAuthConfig{
	GrantType:    opticmp.GrantAuthorizationCode,
	ClientID:     "<client-id>",
	ClientSecret: "<client-secret>",
	// Called after each token refresh. Persist the token here.
	OnTokenRefresh: func(ctx context.Context, token *opticmp.Token) error {
		return nil
	},
}))

// 1. Send the user here to authorize.
url, _ := client.OAuth.GetAuthorizationURL(opticmp.AuthorizationURLOptions{
	RedirectURI: "https://your.app/callback",
	Scope:       "<scopes>",
	State:       "<csrf-state>",
})

// 2. In your callback handler, exchange the code for a token.
token, _ := client.OAuth.ExchangeCode(ctx, opticmp.ExchangeCodeOptions{
	Code:        "<code-from-callback>",
	RedirectURI: "https://your.app/callback",
})

// 3. Refresh when needed.
refreshed, _ := client.OAuth.RefreshToken(ctx)

// Fetch the authenticated user's info.
userInfo, _ := client.OAuth.GetUserInfo(ctx, opticmp.UserInfoOptions{})
```

### OAuth2 — client credentials

For server-to-server integrations.

```go
client := opticmp.New(opticmp.WithOAuth(opticmp.OAuthConfig{
	GrantType:    opticmp.GrantClientCredentials,
	ClientID:     "<client-id>",
	ClientSecret: "<client-secret>",
	Scope:        "<scopes>",
}))

token, _ := client.OAuth.GetClientCredentialsToken(ctx)
```

Tokens are refreshed automatically on expiry when a refresh token is available.

## Pagination

List endpoints support pagination. See [`examples/`](./examples) for a complete
walkthrough.

## Errors

A failed request returns an `*opticmp.HTTPError` with the status, decoded body,
and response headers:

```go
var httpErr *opticmp.HTTPError
if errors.As(err, &httpErr) {
	fmt.Println(httpErr.Status, httpErr.Data)
}
```

## Documentation

- [CMP API Reference](https://docs.developers.optimizely.com/content-marketing-platform/reference/api-reference)
- [Examples](./examples)

## License

Licensed under the MIT License. Check the [LICENSE](../LICENSE) file for details.
