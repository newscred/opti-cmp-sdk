import Hook from "before-after-hook";
import { describe, expect, it } from "vitest";

import type { RequestHook } from "./types.js";

import { getEndpointOptions } from "./get-endpoint-options.js";

describe("getEndpointOptions", () => {
  const createHook = (): RequestHook => new Hook.Singular();

  it("extracts baseUrl from options", () => {
    const hook = createHook();
    const result = getEndpointOptions(
      { baseUrl: "https://api.example.com" },
      hook,
    );

    expect(result.baseUrl).toBe("https://api.example.com");
  });

  it("extracts headers from options", () => {
    const hook = createHook();
    const result = getEndpointOptions(
      { headers: { "x-custom": "value" } },
      hook,
    );

    expect(result.headers).toEqual({ "x-custom": "value" });
  });

  it("extracts request options", () => {
    const hook = createHook();
    const customFetch = () => Promise.resolve(new Response());
    const result = getEndpointOptions(
      { request: { fetch: customFetch } },
      hook,
    );

    expect(result.request!.fetch).toBe(customFetch);
  });

  it("attaches hook to request options", () => {
    const hook = createHook();
    const result = getEndpointOptions({}, hook);

    expect(result.request!.hook).toBe(hook);
  });

  it("ignores auth option", () => {
    const hook = createHook();
    const result = getEndpointOptions({ auth: { token: "secret" } }, hook);

    expect(result).not.toHaveProperty("auth");
  });

  it("provides default empty headers", () => {
    const hook = createHook();
    const result = getEndpointOptions({}, hook);

    expect(result.headers).toEqual({});
  });

  it("merges provided headers with defaults", () => {
    const hook = createHook();
    const result = getEndpointOptions(
      { headers: { authorization: "Bearer token" } },
      hook,
    );

    expect(result.headers!.authorization).toBe("Bearer token");
  });
});
