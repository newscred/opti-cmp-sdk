import type {
  Endpoint,
  EndpointDefaults,
  EndpointOptions,
  EndpointParams,
} from "../endpoint/types.js";
import type { Request, Response } from "./types.js";

import { fetchWrapper } from "./fetch-wrapper.js";

export function withDefaults(
  oldEndpoint: Endpoint,
  newDefaults: EndpointParams,
): Request {
  const endpoint = oldEndpoint.defaults(newDefaults);

  function request<T>(
    route: EndpointOptions | string,
    params?: EndpointParams,
  ): Promise<Response<T>> {
    const endpointOptions = endpoint.merge(route as string, params);

    if (!endpointOptions.request?.hook) {
      return fetchWrapper<T>(endpoint.parse(endpointOptions));
    }

    return endpointOptions.request.hook(
      (options: EndpointDefaults) => fetchWrapper<T>(endpoint.parse(options)),
      endpointOptions,
    );
  }

  request.defaults = withDefaults.bind(null, endpoint);
  request.endpoint = endpoint;

  return request as Request;
}
