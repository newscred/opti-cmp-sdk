import type { APIClient, Options } from "../../client/types.js";
import type { Routes } from "./types.js";

export function registerEndpointsPlugin(
  client: APIClient,
  _options: Options,
): void {
  client.registerEndpoints = function registerEndpoints(routes: Routes): void {
    for (const namespaceName of Object.keys(routes)) {
      if (!client[namespaceName]) {
        client[namespaceName] = {};
      }

      for (const [methodName, route] of Object.entries(routes[namespaceName])) {
        client[namespaceName][methodName] = client.request.defaults({
          method: route.method.toUpperCase(),
          url: route.url,
        });
      }
    }
  };
}

export default registerEndpointsPlugin;
