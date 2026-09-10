# Examples

Runnable examples for the [`opti-cmp-sdk/go`](../) SDK.

| Directory                                                | Description                              |
| -------------------------------------------------------- | ---------------------------------------- |
| [`token`](./token)                                       | Authenticate with a static access token  |
| [`oauth-client-credentials`](./oauth-client-credentials) | OAuth2 for server-to-server integrations |
| [`oauth-authorization-code`](./oauth-authorization-code) | OAuth2 for user-facing apps              |
| [`pagination`](./pagination)                             | Iterate through paginated list results   |

## Running

Clone the repository, then run any example from the `go` directory:

```bash
cd go
OPTI_CMP_TOKEN="<auth-token>" go run ./examples/token
```

Credentials are read from environment variables:

| Variable                                       | Used by               |
| ---------------------------------------------- | --------------------- |
| `OPTI_CMP_TOKEN`                               | `token`, `pagination` |
| `OPTI_CMP_CLIENT_ID`, `OPTI_CMP_CLIENT_SECRET` | the OAuth2 examples   |
