import type { EndpointPath } from "../types/routes.js";

export interface Endpoint {
  (route: string, params?: EndpointParams): RequestOptions;
  (options: EndpointOptions): RequestOptions;
  DEFAULTS: EndpointDefaults;
  defaults(params: EndpointParams): Endpoint;
  merge(route: string, params?: EndpointParams): EndpointDefaults;
  merge(options: EndpointParams): EndpointDefaults;
  parse(options: EndpointDefaults): RequestOptions;
}

export type EndpointDefaults = EndpointOptions & {
  baseUrl: string;
  headers: Headers;
};

export type EndpointOptions = EndpointParams & {
  method: RequestMethod;
  url: EndpointPath | (string & {});
};

export type EndpointParams = {
  // URL template parameters and other endpoint-specific params
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [param: string]: any;
  baseUrl?: string;
  headers?: Headers;
  request?: RequestConfig;
};

export type Headers = Record<string, string>;

export interface RequestConfig {
  [option: string]: unknown;
  fetch?: typeof fetch;
  // Hook from before-after-hook library - typed loosely for compatibility
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  hook?: any;
  signal?: AbortSignal;
  timeout?: number;
}

export type RequestMethod = "DELETE" | "GET" | "PATCH" | "POST" | "PUT";

export type RequestOptions = {
  body?: unknown;
  headers: Headers;
  method: EndpointOptions["method"];
  request?: RequestConfig;
  url: EndpointOptions["url"];
};
