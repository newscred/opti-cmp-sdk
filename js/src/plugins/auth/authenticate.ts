import type { AuthOptions, AuthState } from "./types.js";

export function authenticate(auth: AuthState, options?: AuthOptions): void {
  if (!options) {
    auth.oauth = undefined;
    auth.token = undefined;
    return;
  }

  if (options.type === "oauth") {
    auth.oauth = {
      oauth: options,
      token: options.token,
    };
    auth.token = undefined;
    return;
  }

  if (!options.token) {
    throw new Error("Token authentication requires a token");
  }

  auth.token = options.token;
  auth.oauth = undefined;
}
