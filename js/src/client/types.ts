import type { HookSingular } from "before-after-hook";

import type { RequestOptions } from "../endpoint/types.js";
import type { HTTPError } from "../error/index.js";
import type { AuthOptions, OAuthMethods } from "../plugins/auth/types.js";
import type { Routes } from "../plugins/register-endpoints/types.js";
import type { Request, RequestAdapter, Response } from "../request/types.js";
import type { APIEndpoints } from "../types/endpoints.js";
import type { Pagination } from "../types/schema.js";

export interface APIClient extends APIEndpoints {
  // Dynamic properties added by plugins
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;

  authenticate: (options?: AuthOptions) => void;

  getNextPage: <T extends { pagination: Pagination }>(
    response: Response<T>,
  ) => Promise<Response<T>>;
  getPreviousPage: <T extends { pagination: Pagination }>(
    response: Response<T>,
  ) => Promise<Response<T>>;
  hasNextPage: <T extends { pagination: Pagination }>(
    response: Response<T>,
  ) => boolean;
  hasPreviousPage: <T extends { pagination: Pagination }>(
    response: Response<T>,
  ) => boolean;

  oauth: OAuthMethods;

  registerEndpoints: (routes: Routes) => void;

  request: Request;
  requestHook: RequestHook;
}

export interface APIClientFactory {
  new (options?: Options): APIClient;
  (options?: Options): APIClient;

  plugins(plugins: Plugin[]): APIClientFactory;
}

export interface Options {
  // Allow additional options for extensibility
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [option: string]: any;
  adapter?: {
    request?: RequestAdapter;
  };
  auth?: AuthOptions;
  baseUrl?: string;
  request?: RequestOptions["request"];
}

export type Plugin = (client: APIClient, options: Options) => void;

export type RequestHook = HookSingular<
  RequestOptions,
  Response<unknown>,
  HTTPError
>;
