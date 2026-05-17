import type { Options } from "../../client/types.js";
import type { AuthPluginState } from "./types.js";

export function beforeRequest(
  state: AuthPluginState,
  requestOptions: Options,
): void {
  requestOptions.headers.authorization = `Bearer ${state.auth.token}`;
}
