import { factory } from "./client/factory.js";
import authPlugin from "./plugins/auth/index.js";
import paginationPlugin from "./plugins/pagination/index.js";
import registerApiEndpointsPlugin from "./plugins/register-api-endpoints/index.js";
import registerEndpointsPlugin from "./plugins/register-endpoints/index.js";

export type { APIClient, Plugin } from "./client/types.js";

export { HTTPError } from "./error/index.js";
export type { AuthToken } from "./plugins/auth/types.js";
export type { Request, Response } from "./request/types.js";
export type { APIEndpoints } from "./types/endpoints.js";
export type * as Params from "./types/params.js";

export type * as Schema from "./types/schema-public.js";

export const OptiCMP = factory([
  registerEndpointsPlugin,
  registerApiEndpointsPlugin,
  authPlugin,
  paginationPlugin,
]);
