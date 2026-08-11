# Optimizely CMP SDK for TypeScript

TypeScript SDK for the [Optimizely CMP](https://www.optimizely.com/products/content-marketing/) (Content Marketing Platform) API.

## Install

```bash
npm install @optimizely/cmp-sdk
# or
pnpm add @optimizely/cmp-sdk
```

## Usage

```ts
import { OptiCMP } from "@optimizely/cmp-sdk";

const client = OptiCMP({
  auth: {
    token: "<auth-token>",
  },
});

// List campaigns
const campaigns = await client.campaign.listCampaigns({});

// Get a task
const task = await client.task.getTask({ id: "task-id" });
```

Endpoints are grouped into namespaces (`client.campaign`, `client.task`,
`client.asset`, `client.workflow`, …). Each method takes a single `params`
object and returns a `Promise<Response<T>>`.

## Authentication

The client supports a static token or full OAuth2.

### Static token

Use a pre-obtained access token:

```ts
const client = OptiCMP({
  auth: {
    token: "<auth-token>",
  },
});
```

### OAuth2 — authorization code

For user-facing apps. Redirect the user to the authorization URL, then exchange
the returned code for a token:

```ts
const client = OptiCMP({
  auth: {
    type: "oauth",
    grantType: "authorization_code",
    clientId: "<client-id>",
    clientSecret: "<client-secret>",
    // Called whenever the access token is refreshed — persist it here.
    onTokenRefresh: async (token) => {
      /* save token */
    },
  },
});

// 1. Send the user here to authorize.
const url = client.oauth.getAuthorizationUrl({
  redirectUri: "https://your.app/callback",
  scope: "<scopes>",
  state: "<csrf-state>",
});

// 2. In your callback handler, exchange the code for a token.
const token = await client.oauth.exchangeCode({
  code: "<code-from-callback>",
  redirectUri: "https://your.app/callback",
});

// 3. Refresh when needed.
const refreshed = await client.oauth.refreshToken();

// Fetch the authenticated user's info.
const userInfo = await client.oauth.getUserInfo();
```

### OAuth2 — client credentials

For server-to-server integrations:

```ts
const client = OptiCMP({
  auth: {
    type: "oauth",
    grantType: "client_credentials",
    clientId: "<client-id>",
    clientSecret: "<client-secret>",
    scope: "<scopes>",
  },
});

const token = await client.oauth.getClientCredentialsToken();
```

Tokens are refreshed automatically on expiry when a refresh token is available.

## Pagination

List endpoints support pagination. See [`examples/`](./examples) for a complete
walkthrough.

## Documentation

- [CMP API Reference](https://docs.developers.optimizely.com/content-marketing-platform/reference/api-reference)
- [Examples](./examples)

## License

Licensed under the MIT License. Check the [LICENSE](../LICENSE) file for details.
