import type { RequestConfig } from "../../endpoint/types.js";
import type { Response } from "../../request/types.js";

export type AuthOAuthAuthorizationCodeOptions = AuthOAuthBaseOptions & {
  grantType: "authorization_code";
  token?: AuthToken;
};

export type AuthOAuthClientCredentialsOptions = AuthOAuthBaseOptions & {
  grantType: "client_credentials";
  scope?: string;
  token?: AuthToken;
};

export type AuthOAuthOptions =
  | AuthOAuthAuthorizationCodeOptions
  | AuthOAuthClientCredentialsOptions;

export type AuthOptions = AuthOAuthOptions | AuthTokenOptions;

export type AuthState = {
  oauth?: OAuthPluginState;
  request?: RequestConfig;
  token?: string;
};

export type AuthToken = {
  accessToken: string;
  expiresAt: number;
  refreshToken?: string;
};

export type AuthTokenOptions = {
  token: string;
  type?: "token";
};

export type AuthUserInfo = {
  [claim: string]: unknown;
};

export type OAuthMethods = {
  exchangeCode: (options: {
    code: string;
    redirectUri: string;
  }) => Promise<AuthToken>;
  getAuthorizationUrl: (options: {
    authorizationUrl?: string;
    redirectUri: string;
    scope?: string;
    state?: string;
  }) => string;
  getClientCredentialsToken: () => Promise<AuthToken>;
  getUserInfo: (options?: {
    userInfoUrl?: string;
  }) => Promise<Response<AuthUserInfo>>;
  refreshToken: () => Promise<AuthToken>;
};

export type OAuthPluginState = {
  oauth: AuthOAuthOptions;
  refreshPromise?: Promise<void>;
  token?: AuthToken;
};

type AuthOAuthBaseOptions = {
  clientId: string;
  clientSecret: string;
  onTokenRefresh?: (token: AuthToken) => Promise<void> | void;
  tokenUrl?: string;
  type: "oauth";
};
