import type { APIClient, Options } from "../../client/types.js";
import type { AuthPluginState } from "./types.js";

import { beforeRequest } from "./before-request.js";

export default function authPlugin(client: APIClient, options: Options): void {
  if (!options.auth) return;

  const state: AuthPluginState = {
    auth: options.auth,
    client,
  };

  client.requestHook.before(beforeRequest.bind(null, state));
}
