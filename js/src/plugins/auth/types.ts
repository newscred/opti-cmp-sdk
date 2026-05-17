import type { APIClient } from "../../client/types.js";

export type AuthOptions = {
  token: string;
};

export type AuthPluginState = {
  auth: AuthOptions;
  client: APIClient;
};
