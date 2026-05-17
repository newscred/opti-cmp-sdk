import type { EndpointDefaults, EndpointParams } from "./types.js";

import { lowerCaseHeaderFields } from "../utils/lower-case-header-fields.js";
import { deepmerge } from "../utils/merge.js";

export function merge(
  defaults: EndpointDefaults | null,
  route: EndpointParams | string,
  params?: EndpointParams,
): EndpointDefaults {
  let options: EndpointParams;

  if (typeof route === "string") {
    const [first, second] = route.split(" ");
    const method = second ? first : undefined;
    const url = second || first;

    options = { ...params };
    if (method) {
      options.method = method.toUpperCase() as EndpointDefaults["method"];
    }
    options.url = url;
  } else {
    options = route || {};
  }

  if (options.headers) {
    options.headers = lowerCaseHeaderFields(options.headers);
  }

  const mergedOptions = defaults ? deepmerge(defaults, options) : options;

  return mergedOptions as EndpointDefaults;
}
