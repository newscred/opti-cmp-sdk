import type { APIClient, Options } from "../../client/types.js";
import type { Routes } from "../register-endpoints/types.js";

import routes from "./routes.json" with { type: "json" };

const typedRoutes = routes as Routes;

export function registerApiEndpointsPlugin(
  client: APIClient,
  _options: Options,
): void {
  client.registerEndpoints(typedRoutes);
}

export default registerApiEndpointsPlugin;
