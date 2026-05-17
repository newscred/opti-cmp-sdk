import type { APIClient } from "../../client/types.js";

export type AuthenticateOptions = {
  token: string;
  type: "token";
};

export type AuthenticatePluginState = {
  auth?: AuthenticateOptions;
  client: APIClient;
};
