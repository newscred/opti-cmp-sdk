/**
 * OAuth2 client credentials grant, for server-to-server integrations.
 *
 *   CMP_CLIENT_ID=<client-id> CMP_CLIENT_SECRET=<client-secret> \
 *     npx tsx examples/oauth-client-credentials.ts
 */
import { OptiCMP } from "opti-cmp";

const clientId = process.env.CMP_CLIENT_ID;
const clientSecret = process.env.CMP_CLIENT_SECRET;
if (!clientId || !clientSecret) {
  throw new Error("Set CMP_CLIENT_ID and CMP_CLIENT_SECRET");
}

const client = OptiCMP({
  auth: {
    clientId,
    clientSecret,
    grantType: "client_credentials",
    scope: "<scopes>",
    type: "oauth",
  },
});

const token = await client.oauth.getClientCredentialsToken();
console.log("Token expires at:", new Date(token.expiresAt).toISOString());

// The token is attached to requests automatically, and refreshed on expiry.
const campaigns = await client.campaign.listCampaigns({});
console.log(campaigns.data);
