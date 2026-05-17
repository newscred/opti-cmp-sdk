import type { APIClientFactory, Plugin } from "./types.js";

import { factory } from "./factory.js";

export function registerPlugins(
  existingPlugins: Plugin[] = [],
  newPlugins: Plugin[] = [],
): APIClientFactory {
  const seen = new Set(existingPlugins);
  const plugins = [...existingPlugins];

  for (const plugin of newPlugins) {
    if (!seen.has(plugin)) {
      seen.add(plugin);
      plugins.push(plugin);
    }
  }

  return factory(plugins);
}
