import type {
  Endpoint,
  EndpointDefaults,
  EndpointParams,
  RequestOptions,
} from "./types.js";

import { merge } from "./merge.js";
import { parse } from "./parse.js";

export function withDefaults(
  oldDefaults: EndpointDefaults | null,
  newDefaults: EndpointParams,
): Endpoint {
  const DEFAULTS = merge(oldDefaults, newDefaults);

  const endpoint = endpointWithDefaults.bind(null, DEFAULTS) as Endpoint;

  endpoint.DEFAULTS = DEFAULTS;
  endpoint.defaults = withDefaults.bind(null, DEFAULTS);
  endpoint.merge = merge.bind(null, DEFAULTS);
  endpoint.parse = parse;

  return endpoint;
}

function endpointWithDefaults(
  defaults: EndpointDefaults,
  route: EndpointParams | string,
  params?: EndpointParams,
): RequestOptions {
  return parse(merge(defaults, route, params));
}
