import type { RequestConfig } from "../../../endpoint/types.js";
import type { AuthToken, OAuthPluginState } from "../types.js";

import { requestTokenForGrant } from "./endpoints.js";

const EXPIRY_BUFFER_MS = 60_000;

export async function ensureToken(
  state: OAuthPluginState,
  request?: RequestConfig,
): Promise<AuthToken> {
  if (!state.token || Date.now() >= state.token.expiresAt - EXPIRY_BUFFER_MS) {
    await performRefresh(state, request);
  }

  return state.token as AuthToken;
}

export function performRefresh(
  state: OAuthPluginState,
  request?: RequestConfig,
): Promise<void> {
  state.refreshPromise ??= (async () => {
    try {
      const previous = state.token;
      const token = await requestTokenForGrant(
        state.oauth,
        previous?.refreshToken,
        request,
      );

      state.token = token.refreshToken
        ? token
        : { ...token, refreshToken: previous?.refreshToken };

      await state.oauth.onTokenRefresh?.(state.token);
    } finally {
      state.refreshPromise = undefined;
    }
  })();

  return state.refreshPromise;
}
