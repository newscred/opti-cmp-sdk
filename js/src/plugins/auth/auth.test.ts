import Hook from "before-after-hook";
import { describe, expect, it, vi } from "vitest";

import type { APIClient, Options, RequestHook } from "../../client/types.js";

import authPlugin from "./index.js";

describe("authPlugin", () => {
  function createMockClient(): APIClient {
    const requestHook: RequestHook = new Hook.Singular();
    return {
      request: vi.fn(),
      requestHook,
    } as unknown as APIClient;
  }

  it("adds Bearer authorization header when auth provided", async () => {
    const client = createMockClient();
    const options: Options = { auth: { token: "test-token" } };

    authPlugin(client, options);

    const requestOptions: Options = { headers: {} };
    await client.requestHook(
      async () => ({ data: null, headers: {}, status: 200, url: "" }),
      requestOptions as Parameters<typeof client.requestHook>[1],
    );

    expect(requestOptions.headers.authorization).toBe("Bearer test-token");
  });
});
