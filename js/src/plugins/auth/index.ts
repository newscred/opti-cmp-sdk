import type { APIClient, Options } from "../../client/types.js";
import type { AuthState } from "./types.js";

import { authenticate } from "./authenticate.js";
import { beforeRequest } from "./before-request.js";
import { createOAuthMethods } from "./oauth/index.js";

export default function authPlugin(client: APIClient, options: Options): void {
  const auth: AuthState = {
    request: options.request,
  };

  client.authenticate = authenticate.bind(null, auth);
  client.oauth = createOAuthMethods(auth);

  if (options.auth) {
    authenticate(auth, options.auth);
  }

  client.requestHook.before(beforeRequest.bind(null, auth));
}
