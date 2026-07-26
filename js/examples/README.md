# Examples

Runnable examples for the [`opti-cmp`](../) SDK.

| File                                                           | Description                              |
| -------------------------------------------------------------- | ---------------------------------------- |
| [`token.ts`](./token.ts)                                       | Authenticate with a static access token  |
| [`oauth-client-credentials.ts`](./oauth-client-credentials.ts) | OAuth2 for server-to-server integrations |
| [`oauth-authorization-code.ts`](./oauth-authorization-code.ts) | OAuth2 for user-facing apps              |
| [`pagination.ts`](./pagination.ts)                             | Iterate through paginated list results   |

## Running

Clone the repository, install dependencies, then run any example from the `js`
directory:

```bash
pnpm install
cd js
CMP_TOKEN="<auth-token>" npx tsx examples/token.ts
```

Credentials are read from environment variables:

| Variable                             | Used by                     |
| ------------------------------------ | --------------------------- |
| `CMP_TOKEN`                          | `token.ts`, `pagination.ts` |
| `CMP_CLIENT_ID`, `CMP_CLIENT_SECRET` | the OAuth2 examples         |
