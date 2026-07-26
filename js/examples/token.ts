/**
 * Authenticate with a static access token.
 *
 *   CMP_TOKEN=<auth-token> npx tsx examples/token.ts
 */
import { OptiCMP } from "opti-cmp";

const token = process.env.CMP_TOKEN;
if (!token) throw new Error("Set CMP_TOKEN");

const client = OptiCMP({ auth: { token } });

const campaigns = await client.campaign.listCampaigns({});
console.log(campaigns.data);
