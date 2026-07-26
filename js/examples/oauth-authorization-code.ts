/**
 * OAuth2 authorization code grant, for user-facing apps.
 *
 * Running this prints the URL to send a user to. Completing the flow needs an
 * HTTP server to receive the redirect, so the second half is shown as the
 * `handleCallback` function you would call from your callback route.
 *
 *   CMP_CLIENT_ID=<client-id> CMP_CLIENT_SECRET=<client-secret> \
 *     npx tsx examples/oauth-authorization-code.ts
 */
import { OptiCMP } from "opti-cmp";

const clientId = process.env.CMP_CLIENT_ID;
const clientSecret = process.env.CMP_CLIENT_SECRET;
if (!clientId || !clientSecret) {
  throw new Error("Set CMP_CLIENT_ID and CMP_CLIENT_SECRET");
}

const redirectUri = "https://your.app/callback";

const client = OptiCMP({
  auth: {
    clientId,
    clientSecret,
    grantType: "authorization_code",
    // Called whenever the access token is refreshed. Persist the token here so
    // it survives a restart.
    onTokenRefresh: async (token) => {
      console.log(
        "Token refreshed, expires at:",
        new Date(token.expiresAt).toISOString(),
      );
    },
    type: "oauth",
  },
});

// 1. Send the user here to authorize your app.
const authorizationUrl = client.oauth.getAuthorizationUrl({
  redirectUri,
  scope: "<scopes>",
  state: "<csrf-state>",
});
console.log("Authorize at:", authorizationUrl);

// 2. The user is redirected back to `redirectUri` with a `code` query
//    parameter. Exchange it for a token, then use the client as usual.
export async function handleCallback(code: string) {
  const token = await client.oauth.exchangeCode({ code, redirectUri });
  console.log("Token expires at:", new Date(token.expiresAt).toISOString());

  const userInfo = await client.oauth.getUserInfo();
  console.log("Signed in as:", userInfo.data);
}
