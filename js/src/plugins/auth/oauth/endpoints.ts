import type { RequestConfig } from "../../../endpoint/types.js";
import type { Response } from "../../../request/types.js";
import type { AuthOAuthOptions, AuthToken, AuthUserInfo } from "../types.js";

import { HTTPError } from "../../../error/index.js";
import { fetchWrapper } from "../../../request/fetch-wrapper.js";
import {
  DEFAULT_AUTHORIZATION_URL,
  DEFAULT_TOKEN_URL,
  DEFAULT_USERINFO_URL,
} from "./constants.js";

export type ExchangeAuthorizationCodeOptions = {
  clientId: string;
  clientSecret: string;
  code: string;
  fetch?: typeof fetch;
  redirectUri: string;
  tokenUrl?: string;
};

export type GetAuthorizationUrlOptions = {
  authorizationUrl?: string;
  clientId: string;
  redirectUri: string;
  scope?: string;
  state?: string;
};

export type RequestClientCredentialsTokenOptions = {
  clientId: string;
  clientSecret: string;
  fetch?: typeof fetch;
  scope?: string;
  tokenUrl?: string;
};

export function exchangeAuthorizationCode(
  options: ExchangeAuthorizationCodeOptions,
): Promise<AuthToken> {
  return requestToken(
    options.tokenUrl ?? DEFAULT_TOKEN_URL,
    {
      client_id: options.clientId,
      client_secret: options.clientSecret,
      code: options.code,
      grant_type: "authorization_code",
      redirect_uri: options.redirectUri,
    },
    options.fetch,
  );
}

export function fetchUserInfo(
  accessToken: string,
  url: string = DEFAULT_USERINFO_URL,
  fetchImpl?: typeof fetch,
): Promise<Response<AuthUserInfo>> {
  return fetchWrapper<AuthUserInfo>({
    headers: {
      accept: "application/json",
      authorization: `Bearer ${accessToken}`,
    },
    method: "GET",
    request: fetchImpl ? { fetch: fetchImpl } : undefined,
    url,
  });
}

export function getAuthorizationUrl(
  options: GetAuthorizationUrlOptions,
): string {
  const url = new URL(options.authorizationUrl ?? DEFAULT_AUTHORIZATION_URL);

  url.searchParams.set("client_id", options.clientId);
  url.searchParams.set("redirect_uri", options.redirectUri);
  url.searchParams.set("response_type", "code");
  if (options.scope) {
    url.searchParams.set("scope", options.scope);
  }
  if (options.state) {
    url.searchParams.set("state", options.state);
  }

  return url.toString();
}

export function requestClientCredentialsToken(
  options: RequestClientCredentialsTokenOptions,
): Promise<AuthToken> {
  const params: Record<string, string> = {
    client_id: options.clientId,
    client_secret: options.clientSecret,
    grant_type: "client_credentials",
  };
  if (options.scope) {
    params.scope = options.scope;
  }

  return requestToken(
    options.tokenUrl ?? DEFAULT_TOKEN_URL,
    params,
    options.fetch,
  );
}

export async function requestToken(
  url: string,
  params: Record<string, string>,
  fetchImpl?: typeof fetch,
): Promise<AuthToken> {
  const response = await fetchWrapper<{
    access_token: string;
    expires_in: number;
    refresh_token?: string;
    token_type: "Bearer";
  }>({
    body: new URLSearchParams(params),
    headers: {
      accept: "application/json",
      "content-type": "application/x-www-form-urlencoded",
    },
    method: "POST",
    request: fetchImpl ? { fetch: fetchImpl } : undefined,
    url,
  }).catch((error: unknown) => {
    // The request body carries client_secret / refresh_token / code. Drop it
    // from the error so secrets are not leaked to callers that log it.
    if (error instanceof HTTPError && error.request) {
      error.request.body = undefined;
    }
    throw error;
  });

  const data = response.data;

  return {
    accessToken: data.access_token,
    expiresAt: Date.now() + data.expires_in * 1000,
    refreshToken: data.refresh_token,
  };
}

export function requestTokenForGrant(
  oauth: AuthOAuthOptions,
  refreshToken: string | undefined,
  request?: RequestConfig,
): Promise<AuthToken> {
  if (oauth.grantType === "authorization_code") {
    if (!refreshToken) {
      throw new Error(
        "Cannot refresh the access token without a refresh token",
      );
    }

    return requestToken(
      oauth.tokenUrl ?? DEFAULT_TOKEN_URL,
      {
        client_id: oauth.clientId,
        client_secret: oauth.clientSecret,
        grant_type: "refresh_token",
        refresh_token: refreshToken,
      },
      request?.fetch,
    );
  }

  return requestClientCredentialsToken({
    clientId: oauth.clientId,
    clientSecret: oauth.clientSecret,
    fetch: request?.fetch,
    scope: oauth.scope,
    tokenUrl: oauth.tokenUrl,
  });
}
