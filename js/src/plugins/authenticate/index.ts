import type { APIClient, Options } from "../../client/types.js";
import type { AuthenticatePluginState } from "./types.js";

import { authenticate } from "./authenticate.js";
import { beforeRequest } from "./before-request.js";

export default function authenticatePlugin(
  client: APIClient,
  options: Options,
): void {
  if (options.auth) return;

  const state: AuthenticatePluginState = {
    auth: undefined,
    client,
  };

  client.authenticate = authenticate.bind(null, state);
  client.requestHook.before(beforeRequest.bind(null, state));
}
