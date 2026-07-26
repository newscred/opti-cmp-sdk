/**
 * Iterate through every page of a paginated list endpoint.
 *
 *   CMP_TOKEN=<auth-token> npx tsx examples/pagination.ts
 */
import { OptiCMP } from "opti-cmp";

const token = process.env.CMP_TOKEN;
if (!token) throw new Error("Set CMP_TOKEN");

const client = OptiCMP({ auth: { token } });

let response = await client.campaign.listCampaigns({});
let page = 1;

while (true) {
  console.log(`Page ${page}:`, response.data);

  if (!client.hasNextPage(response)) break;

  response = await client.getNextPage(response);
  page += 1;
}

console.log(`Fetched ${page} page(s).`);
