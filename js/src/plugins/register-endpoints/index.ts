import type { APIClient, Options } from "../../client/types.js";
import type { Routes } from "./types.js";

export function registerEndpointsPlugin(
  client: APIClient,
  options: Options,
): void {
  client.registerEndpoints = function registerEndpoints(routes: Routes): void {
    const makeMethod = options.adapter?.request
      ? (namespace: string, operation: string) => {
          return (params: Record<string, unknown>) =>
            options.adapter!.request!({ namespace, operation, params });
        }
      : (namespace: string, operation: string) => {
          const route = routes[namespace][operation];
          return client.request.defaults({
            method: route.method.toUpperCase(),
            url: route.url,
          });
        };

    for (const namespaceName of Object.keys(routes)) {
      if (!client[namespaceName]) {
        client[namespaceName] = {};
      }
      for (const methodName of Object.keys(routes[namespaceName])) {
        client[namespaceName][methodName] = makeMethod(
          namespaceName,
          methodName,
        );
      }
    }
  };
}

export default registerEndpointsPlugin;
