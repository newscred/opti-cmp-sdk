import { beforeEach, describe, expect, it, vi } from "vitest";

import type { RequestOptions } from "../endpoint/types.js";

import { HTTPError } from "../error/index.js";
import { fetchWrapper } from "./fetch-wrapper.js";

describe("fetchWrapper", () => {
  const baseOptions: RequestOptions = {
    headers: { accept: "application/json" },
    method: "GET",
    url: "https://api.example.com/users",
  };

  beforeEach(() => {
    vi.restoreAllMocks();
  });

  describe("successful responses", () => {
    it("should return JSON data for application/json response", async () => {
      const mockData = { id: 1, name: "John" };
      const mockFetch = vi.fn().mockResolvedValue({
        headers: new Headers({ "content-type": "application/json" }),
        json: () => Promise.resolve(mockData),
        ok: true,
        status: 200,
        url: "https://api.example.com/users",
      });

      const result = await fetchWrapper({
        ...baseOptions,
        request: { fetch: mockFetch },
      });

      expect(result.data).toEqual(mockData);
      expect(result.status).toBe(200);
      expect(result.url).toBe("https://api.example.com/users");
    });

    it("should return text data for text content types", async () => {
      const mockFetch = vi.fn().mockResolvedValue({
        headers: new Headers({ "content-type": "text/html" }),
        ok: true,
        status: 200,
        text: () => Promise.resolve("<html></html>"),
        url: "https://api.example.com/page",
      });

      const result = await fetchWrapper({
        ...baseOptions,
        request: { fetch: mockFetch },
      });

      expect(result.data).toBe("<html></html>");
    });

    it("should return undefined for 204 No Content", async () => {
      const mockFetch = vi.fn().mockResolvedValue({
        headers: new Headers(),
        ok: true,
        status: 204,
        url: "https://api.example.com/users/1",
      });

      const result = await fetchWrapper({
        ...baseOptions,
        method: "DELETE",
        request: { fetch: mockFetch },
      });

      expect(result.data).toBeUndefined();
      expect(result.status).toBe(204);
    });
  });

  describe("error handling", () => {
    it("should throw HTTPError with response data for non-ok responses", async () => {
      const errorData = { error: "Validation failed", fields: ["email"] };
      const mockFetch = vi.fn().mockResolvedValue({
        headers: new Headers({ "content-type": "application/json" }),
        json: () => Promise.resolve(errorData),
        ok: false,
        status: 400,
        statusText: "Bad Request",
        url: "https://api.example.com/users",
      });

      try {
        await fetchWrapper({
          ...baseOptions,
          method: "POST",
          request: { fetch: mockFetch },
        });
        expect.fail("Should have thrown");
      } catch (error) {
        expect(error).toBeInstanceOf(HTTPError);
        expect((error as HTTPError).status).toBe(400);
        expect((error as HTTPError).data).toEqual(errorData);
      }
    });
  });

  describe("response headers", () => {
    it("should capture response headers", async () => {
      const mockFetch = vi.fn().mockResolvedValue({
        headers: new Headers({
          "content-type": "application/json",
          "x-request-id": "abc123",
        }),
        json: () => Promise.resolve({}),
        ok: true,
        status: 200,
        url: "https://api.example.com/users",
      });

      const result = await fetchWrapper({
        ...baseOptions,
        request: { fetch: mockFetch },
      });

      expect(result.headers["content-type"]).toBe("application/json");
      expect(result.headers["x-request-id"]).toBe("abc123");
    });
  });
});
