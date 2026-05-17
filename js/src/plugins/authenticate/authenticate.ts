import type { AuthenticateOptions, AuthenticatePluginState } from "./types.js";

export function authenticate(
  state: AuthenticatePluginState,
  options?: AuthenticateOptions,
): void {
  if (!options) {
    state.auth = undefined;
    return;
  }

  if (options.type === "token" && !options.token) {
    throw new Error("Token authentication requires a token");
  }

  state.auth = options;
}
