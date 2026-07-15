import type {
  AuthState,
  AuthToken,
  OAuthMethods,
  OAuthPluginState,
} from "../types.js";

import {
  exchangeAuthorizationCode,
  fetchUserInfo,
  getAuthorizationUrl,
  requestClientCredentialsToken,
} from "./endpoints.js";
import { ensureToken, performRefresh } from "./refresh.js";

export function createOAuthMethods(auth: AuthState): OAuthMethods {
  function requireOAuth(): OAuthPluginState {
    if (!auth.oauth) {
      throw new Error(
        'OAuth is not configured. Provide auth: { type: "oauth", ... } when creating the client, or call client.authenticate({ type: "oauth", ... }).',
      );
    }
    return auth.oauth;
  }

  return {
    exchangeCode: async (options) => {
      const state = requireOAuth();
      const token = await exchangeAuthorizationCode({
        clientId: state.oauth.clientId,
        clientSecret: state.oauth.clientSecret,
        code: options.code,
        fetch: auth.request?.fetch,
        redirectUri: options.redirectUri,
        tokenUrl: state.oauth.tokenUrl,
      });

      return storeToken(state, token);
    },

    getAuthorizationUrl: (options) => {
      const state = requireOAuth();
      return getAuthorizationUrl({
        authorizationUrl: options.authorizationUrl,
        clientId: state.oauth.clientId,
        redirectUri: options.redirectUri,
        scope: options.scope,
        state: options.state,
      });
    },

    getClientCredentialsToken: async () => {
      const state = requireOAuth();
      const token = await requestClientCredentialsToken({
        clientId: state.oauth.clientId,
        clientSecret: state.oauth.clientSecret,
        fetch: auth.request?.fetch,
        scope:
          state.oauth.grantType === "client_credentials"
            ? state.oauth.scope
            : undefined,
        tokenUrl: state.oauth.tokenUrl,
      });

      return storeToken(state, token);
    },

    getUserInfo: async (options) => {
      const state = requireOAuth();
      const token = await ensureToken(state, auth.request);
      return fetchUserInfo(
        token.accessToken,
        options?.userInfoUrl,
        auth.request?.fetch,
      );
    },

    refreshToken: async (): Promise<AuthToken> => {
      const state = requireOAuth();
      await performRefresh(state, auth.request);
      return state.token as AuthToken;
    },
  };
}

async function storeToken(
  state: OAuthPluginState,
  token: AuthToken,
): Promise<AuthToken> {
  state.token = token;
  await state.oauth.onTokenRefresh?.(token);
  return token;
}
