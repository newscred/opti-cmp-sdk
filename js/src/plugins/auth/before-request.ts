import type { Options } from "../../client/types.js";
import type { AuthState } from "./types.js";

import { ensureToken } from "./oauth/refresh.js";

export async function beforeRequest(
  auth: AuthState,
  requestOptions: Options,
): Promise<void> {
  if (auth.oauth) {
    const token = await ensureToken(auth.oauth, requestOptions.request);
    requestOptions.headers.authorization = `Bearer ${token.accessToken}`;
    return;
  }

  if (auth.token) {
    requestOptions.headers.authorization = `Bearer ${auth.token}`;
  }
}
