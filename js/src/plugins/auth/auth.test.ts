import Hook from "before-after-hook";
import { beforeEach, describe, expect, it, vi } from "vitest";

import type { APIClient, Options, RequestHook } from "../../client/types.js";
import type { AuthOAuthOptions } from "./types.js";

import { HTTPError } from "../../error/index.js";
import authPlugin from "./index.js";

function createMockClient(): { client: APIClient; requestHook: RequestHook } {
  const requestHook: RequestHook = new Hook.Singular();
  const client = {
    request: vi.fn(),
    requestHook,
  } as unknown as APIClient;
  return { client, requestHook };
}

async function runRequest(
  requestHook: RequestHook,
  fetch?: typeof globalThis.fetch,
): Promise<Options> {
  const requestOptions: Options = {
    headers: {},
    request: fetch ? { fetch } : undefined,
  };
  await requestHook(
    async () => ({ data: null, headers: {}, status: 200, url: "" }),
    requestOptions as Parameters<typeof requestHook>[1],
  );
  return requestOptions;
}

function tokenResponse(
  body: Record<string, unknown>,
  status = 200,
): globalThis.Response {
  return new Response(JSON.stringify(body), {
    headers: { "content-type": "application/json" },
    status,
  });
}

const baseOAuth: AuthOAuthOptions = {
  clientId: "id",
  clientSecret: "secret",
  grantType: "client_credentials",
  type: "oauth",
};

describe("authPlugin", () => {
  beforeEach(() => {
    vi.useRealTimers();
  });

  it("leaves requests unauthenticated when auth is not provided", async () => {
    const { client, requestHook } = createMockClient();

    authPlugin(client, {});
    const requestOptions = await runRequest(requestHook);

    expect(requestOptions.headers.authorization).toBeUndefined();
  });

  describe("deferred authentication", () => {
    it("configures token auth later via client.authenticate()", async () => {
      const { client, requestHook } = createMockClient();

      authPlugin(client, {});
      client.authenticate({ token: "late-token", type: "token" });
      const requestOptions = await runRequest(requestHook);

      expect(requestOptions.headers.authorization).toBe("Bearer late-token");
    });

    it("configures oauth later via client.authenticate()", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, {});
      client.authenticate({
        clientId: "id",
        clientSecret: "secret",
        grantType: "client_credentials",
        type: "oauth",
      });
      const requestOptions = await runRequest(requestHook, fetch);

      expect(fetch).toHaveBeenCalledTimes(1);
      expect(requestOptions.headers.authorization).toBe("Bearer access-1");
    });

    it("client.oauth methods throw until oauth is configured", async () => {
      const { client } = createMockClient();

      authPlugin(client, {});

      expect(() =>
        client.oauth.getAuthorizationUrl({
          redirectUri: "https://app.test/cb",
        }),
      ).toThrow(/OAuth is not configured/);
      await expect(client.oauth.getUserInfo()).rejects.toThrow(
        /OAuth is not configured/,
      );
    });
  });

  describe("token auth", () => {
    it("adds Bearer authorization header when a token is provided", async () => {
      const { client, requestHook } = createMockClient();

      authPlugin(client, { auth: { token: "test-token" } });
      const requestOptions = await runRequest(requestHook);

      expect(requestOptions.headers.authorization).toBe("Bearer test-token");
    });
  });

  describe("oauth (client credentials)", () => {
    it("fetches a token and sets the Bearer header on first request", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth });
      const requestOptions = await runRequest(requestHook, fetch);

      expect(fetch).toHaveBeenCalledTimes(1);
      const [url, init] = (fetch as unknown as ReturnType<typeof vi.fn>).mock
        .calls[0];
      expect(url).toBe("https://accounts.cmp.optimizely.com/o/oauth2/v1/token");
      const body = init.body as URLSearchParams;
      expect(body.get("client_id")).toBe("id");
      expect(body.get("client_secret")).toBe("secret");
      expect(body.get("grant_type")).toBe("client_credentials");
      expect(requestOptions.headers.authorization).toBe("Bearer access-1");
    });

    it("includes scope and honors a custom tokenUrl", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, {
        auth: {
          ...baseOAuth,
          scope: "openid profile",
          tokenUrl: "https://example.test/token",
        },
      });
      await runRequest(requestHook, fetch);

      const [url, init] = (fetch as unknown as ReturnType<typeof vi.fn>).mock
        .calls[0];
      expect(url).toBe("https://example.test/token");
      expect((init.body as URLSearchParams).get("scope")).toBe(
        "openid profile",
      );
    });

    it("reuses a valid token without fetching again", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth });
      await runRequest(requestHook, fetch);
      const second = await runRequest(requestHook, fetch);

      expect(fetch).toHaveBeenCalledTimes(1);
      expect(second.headers.authorization).toBe("Bearer access-1");
    });

    it("refreshes when the token has expired", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi
        .fn()
        .mockResolvedValueOnce(
          tokenResponse({
            access_token: "access-1",
            expires_in: 60, // < buffer, so immediately considered expired
            token_type: "Bearer",
          }),
        )
        .mockResolvedValueOnce(
          tokenResponse({
            access_token: "access-2",
            expires_in: 3600,
            token_type: "Bearer",
          }),
        ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth });
      await runRequest(requestHook, fetch);
      const second = await runRequest(requestHook, fetch);

      expect(fetch).toHaveBeenCalledTimes(2);
      expect(second.headers.authorization).toBe("Bearer access-2");
    });

    it("invokes onTokenRefresh with issued tokens", async () => {
      const { client, requestHook } = createMockClient();
      const onTokenRefresh = vi.fn();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: { ...baseOAuth, onTokenRefresh } });
      await runRequest(requestHook, fetch);

      expect(onTokenRefresh).toHaveBeenCalledWith({
        accessToken: "access-1",
        expiresAt: expect.any(Number),
      });
    });

    it("uses a seeded token without fetching", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn() as unknown as typeof globalThis.fetch;

      authPlugin(client, {
        auth: {
          ...baseOAuth,
          token: {
            accessToken: "seeded",
            expiresAt: Date.now() + 3_600_000,
          },
        },
      });
      const requestOptions = await runRequest(requestHook, fetch);

      expect(fetch).not.toHaveBeenCalled();
      expect(requestOptions.headers.authorization).toBe("Bearer seeded");
    });

    it("refetches when the seeded token is expired", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, {
        auth: {
          ...baseOAuth,
          token: {
            accessToken: "seeded",
            expiresAt: Date.now() - 1000,
          },
        },
      });
      const requestOptions = await runRequest(requestHook, fetch);

      expect(fetch).toHaveBeenCalledTimes(1);
      expect(requestOptions.headers.authorization).toBe("Bearer access-1");
    });

    it("de-dupes concurrent refreshes into a single token fetch", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth });
      await Promise.all([
        runRequest(requestHook, fetch),
        runRequest(requestHook, fetch),
      ]);

      expect(fetch).toHaveBeenCalledTimes(1);
    });

    it("throws an HTTPError when the token endpoint fails", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({ error: "invalid_client" }, 401),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth });

      await expect(runRequest(requestHook, fetch)).rejects.toBeInstanceOf(
        HTTPError,
      );
    });

    it("clears the in-flight refresh so a later request retries after a failure", async () => {
      const { client, requestHook } = createMockClient();
      const fetch = vi
        .fn()
        .mockResolvedValueOnce(tokenResponse({ error: "invalid_client" }, 401))
        .mockResolvedValueOnce(
          tokenResponse({
            access_token: "access-1",
            expires_in: 3600,
            token_type: "Bearer",
          }),
        ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth });

      await expect(runRequest(requestHook, fetch)).rejects.toBeInstanceOf(
        HTTPError,
      );

      const requestOptions = await runRequest(requestHook, fetch);

      expect(fetch).toHaveBeenCalledTimes(2);
      expect(requestOptions.headers.authorization).toBe("Bearer access-1");
    });
  });

  describe("oauth (authorization code)", () => {
    function authCodeOptions(
      overrides?: Partial<{ onTokenRefresh: () => void }>,
    ): Options {
      return {
        auth: {
          clientId: "id",
          clientSecret: "secret",
          grantType: "authorization_code",
          token: {
            accessToken: "seeded",
            expiresAt: Date.now() - 1000, // expired → refresh on first use
            refreshToken: "refresh-1",
          },
          type: "oauth",
          ...overrides,
        },
      };
    }

    it("refreshes via the refresh_token grant and rotates the refresh token", async () => {
      const { client, requestHook } = createMockClient();
      const onTokenRefresh = vi.fn();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-2",
          expires_in: 3600,
          refresh_token: "refresh-2",
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, authCodeOptions({ onTokenRefresh }));
      const requestOptions = await runRequest(requestHook, fetch);

      const [url, init] = (fetch as unknown as ReturnType<typeof vi.fn>).mock
        .calls[0];
      expect(url).toBe("https://accounts.cmp.optimizely.com/o/oauth2/v1/token");
      const body = init.body as URLSearchParams;
      expect(body.get("grant_type")).toBe("refresh_token");
      expect(body.get("refresh_token")).toBe("refresh-1");
      expect(requestOptions.headers.authorization).toBe("Bearer access-2");
      expect(onTokenRefresh).toHaveBeenCalledWith({
        accessToken: "access-2",
        expiresAt: expect.any(Number),
        refreshToken: "refresh-2",
      });
    });

    it("keeps the previous refresh token when the response omits a new one", async () => {
      const { client, requestHook } = createMockClient();
      const onTokenRefresh = vi.fn();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-2",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, authCodeOptions({ onTokenRefresh }));
      await runRequest(requestHook, fetch);

      expect(onTokenRefresh).toHaveBeenCalledWith(
        expect.objectContaining({ refreshToken: "refresh-1" }),
      );
    });
  });

  describe("client.oauth", () => {
    it("getClientCredentialsToken() fetches, stores, and returns the token", async () => {
      const { client, requestHook } = createMockClient();
      const onTokenRefresh = vi.fn();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, {
        auth: { ...baseOAuth, onTokenRefresh, scope: "openid" },
        request: { fetch },
      });
      const token = await client.oauth.getClientCredentialsToken();

      const [url, init] = (fetch as unknown as ReturnType<typeof vi.fn>).mock
        .calls[0];
      expect(url).toBe("https://accounts.cmp.optimizely.com/o/oauth2/v1/token");
      const body = init.body as URLSearchParams;
      expect(body.get("grant_type")).toBe("client_credentials");
      expect(body.get("scope")).toBe("openid");
      expect(token.accessToken).toBe("access-1");
      expect(onTokenRefresh).toHaveBeenCalledWith(token);

      // Stored on the client — a later request reuses it without re-fetching.
      const requestOptions = await runRequest(requestHook, fetch);
      expect(fetch).toHaveBeenCalledTimes(1);
      expect(requestOptions.headers.authorization).toBe("Bearer access-1");
    });

    it("refreshToken() refreshes on demand and returns the new token", async () => {
      const { client } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth, request: { fetch } });
      const token = await client.oauth.refreshToken();

      expect(fetch).toHaveBeenCalledTimes(1);
      expect(token.accessToken).toBe("access-1");
    });

    it("getAuthorizationUrl() builds the URL, defaulting clientId from config", () => {
      const { client } = createMockClient();

      authPlugin(client, {
        auth: {
          clientId: "id",
          clientSecret: "secret",
          grantType: "authorization_code",
          type: "oauth",
        },
      });

      const parsed = new URL(
        client.oauth.getAuthorizationUrl({
          redirectUri: "https://app.test/callback",
          scope: "openid profile",
          state: "xyz",
        }),
      );

      expect(parsed.origin + parsed.pathname).toBe(
        "https://accounts.cmp.optimizely.com/o/oauth2/v1/auth",
      );
      expect(parsed.searchParams.get("client_id")).toBe("id");
      expect(parsed.searchParams.get("redirect_uri")).toBe(
        "https://app.test/callback",
      );
      expect(parsed.searchParams.get("response_type")).toBe("code");
      expect(parsed.searchParams.get("scope")).toBe("openid profile");
      expect(parsed.searchParams.get("state")).toBe("xyz");
    });

    it("exchangeCode() stores the token and authorizes later requests", async () => {
      const { client, requestHook } = createMockClient();
      const onTokenRefresh = vi.fn();
      const fetch = vi.fn(async () =>
        tokenResponse({
          access_token: "access-1",
          expires_in: 3600,
          refresh_token: "refresh-1",
          token_type: "Bearer",
        }),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, {
        auth: {
          clientId: "id",
          clientSecret: "secret",
          grantType: "authorization_code",
          onTokenRefresh,
          type: "oauth",
        },
        request: { fetch },
      });

      const token = await client.oauth.exchangeCode({
        code: "the-code",
        redirectUri: "https://app.test/callback",
      });

      const [url, init] = (fetch as unknown as ReturnType<typeof vi.fn>).mock
        .calls[0];
      expect(url).toBe("https://accounts.cmp.optimizely.com/o/oauth2/v1/token");
      const body = init.body as URLSearchParams;
      expect(body.get("grant_type")).toBe("authorization_code");
      expect(body.get("client_id")).toBe("id");
      expect(body.get("code")).toBe("the-code");
      expect(body.get("redirect_uri")).toBe("https://app.test/callback");
      expect(token.accessToken).toBe("access-1");
      expect(onTokenRefresh).toHaveBeenCalledWith(token);

      // The stored token authorizes a subsequent request without re-fetching.
      const requestOptions = await runRequest(requestHook, fetch);
      expect(fetch).toHaveBeenCalledTimes(1);
      expect(requestOptions.headers.authorization).toBe("Bearer access-1");
    });

    it("getUserInfo() ensures a token then returns the claims", async () => {
      const { client } = createMockClient();
      const fetch = vi
        .fn()
        .mockResolvedValueOnce(
          tokenResponse({
            access_token: "access-1",
            expires_in: 3600,
            token_type: "Bearer",
          }),
        )
        .mockResolvedValueOnce(
          tokenResponse({ email: "u@test.dev", sub: "user-1" }),
        ) as unknown as typeof globalThis.fetch;

      authPlugin(client, { auth: baseOAuth, request: { fetch } });
      const info = await client.oauth.getUserInfo();

      expect(fetch).toHaveBeenCalledTimes(2);
      const [url, init] = (fetch as unknown as ReturnType<typeof vi.fn>).mock
        .calls[1];
      expect(url).toBe(
        "https://accounts.cmp.optimizely.com/o/oauth2/v1/userinfo",
      );
      expect(init.headers.authorization).toBe("Bearer access-1");
      expect(info.data).toEqual({ email: "u@test.dev", sub: "user-1" });
    });

    it("getUserInfo() throws an HTTPError when the token is invalid", async () => {
      const { client } = createMockClient();
      const fetch = vi.fn(async () =>
        tokenResponse({ error: "invalid_token" }, 401),
      ) as unknown as typeof globalThis.fetch;

      authPlugin(client, {
        auth: {
          ...baseOAuth,
          token: {
            accessToken: "access-1",
            expiresAt: Date.now() + 3_600_000,
          },
        },
        request: { fetch },
      });

      await expect(client.oauth.getUserInfo()).rejects.toBeInstanceOf(
        HTTPError,
      );
    });
  });
});
