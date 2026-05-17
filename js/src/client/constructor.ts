import Hook from "before-after-hook";

import type { APIClient, Options, Plugin, RequestHook } from "./types.js";

import { request } from "../request/index.js";
import { getEndpointOptions } from "./get-endpoint-options.js";

export function constructor(
  plugins: Plugin[],
  options: Options = {},
): APIClient {
  const requestHook: RequestHook = new Hook.Singular();

  const client = {
    request: request.defaults(getEndpointOptions(options, requestHook)),
    requestHook,
  } as APIClient;

  for (const plugin of plugins) {
    plugin(client, options);
  }

  return client;
}
