import type { AuthenticatePluginState } from "./types.js";

import { Options } from "../../client/types.js";

export function beforeRequest(
  state: AuthenticatePluginState,
  requestOptions: Options,
): void {
  if (!state.auth) return;

  requestOptions.headers.authorization = `Bearer ${state.auth.token}`;
}
