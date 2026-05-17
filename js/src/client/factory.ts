import type { APIClientFactory, Plugin } from "./types.js";

import { constructor } from "./constructor.js";
import { registerPlugins } from "./register-plugins.js";

export function factory(plugins: Plugin[] = []): APIClientFactory {
  const OptiCMP = constructor.bind(null, plugins) as APIClientFactory;
  OptiCMP.plugins = registerPlugins.bind(null, plugins);
  return OptiCMP;
}
