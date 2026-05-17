import { EndpointParams } from "../endpoint/types.js";
import { deepmerge } from "../utils/merge.js";
import { pick } from "../utils/pick.js";
import { Options, RequestHook } from "./types.js";

export function getEndpointOptions(
  clientOptions: Options,
  requestHook: RequestHook,
): EndpointParams {
  const endpointOptions = deepmerge(
    { headers: {}, request: {} },
    pick(clientOptions, ["baseUrl", "headers", "request"]),
  );

  endpointOptions.request.hook = requestHook;

  return endpointOptions;
}
