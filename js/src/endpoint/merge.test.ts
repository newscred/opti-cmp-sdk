import { describe, expect, it } from "vitest";

import type { EndpointDefaults } from "./types.js";

import { merge } from "./merge.js";

describe("merge", () => {
  const defaults: EndpointDefaults = {
    baseUrl: "https://api.example.com",
    headers: { accept: "application/json" },
    method: "GET",
    url: "/default",
  };

  describe("route string parsing", () => {
    it('should parse "METHOD /path" format', () => {
      const result = merge(defaults, "POST /users");
      expect(result.method).toBe("POST");
      expect(result.url).toBe("/users");
    });

    it("should use default method when only path provided", () => {
      const result = merge(defaults, "/users");
      expect(result.method).toBe("GET");
      expect(result.url).toBe("/users");
    });
  });

  describe("parameter merging", () => {
    it("should merge parameters into result", () => {
      const result = merge(defaults, "/users", { limit: 10, page: 1 });
      expect(result.page).toBe(1);
      expect(result.limit).toBe(10);
    });

    it("should merge headers", () => {
      const result = merge(defaults, "/users", {
        headers: { authorization: "Bearer token" },
      });
      expect(result.headers.accept).toBe("application/json");
      expect(result.headers.authorization).toBe("Bearer token");
    });

    it("should lowercase header keys", () => {
      const result = merge(defaults, "/users", {
        headers: {
          Authorization: "Bearer token",
          "Content-Type": "application/json",
        },
      });
      expect(result.headers.authorization).toBe("Bearer token");
      expect(result.headers["content-type"]).toBe("application/json");
    });

    it("should override defaults with params", () => {
      const result = merge(defaults, "/users", {
        baseUrl: "https://other.com",
        headers: { accept: "text/plain" },
      });
      expect(result.baseUrl).toBe("https://other.com");
      expect(result.headers.accept).toBe("text/plain");
    });
  });

  describe("object options", () => {
    it("should accept object with method and url", () => {
      const result = merge(defaults, { method: "POST", url: "/items" });
      expect(result.method).toBe("POST");
      expect(result.url).toBe("/items");
    });

    it("should merge additional properties from object", () => {
      const result = merge(defaults, {
        headers: { "x-custom": "value" },
        method: "GET",
        url: "/users",
      });
      expect(result.headers["x-custom"]).toBe("value");
    });
  });

  describe("baseUrl handling", () => {
    it("should preserve baseUrl from defaults", () => {
      const result = merge(defaults, "/users");
      expect(result.baseUrl).toBe("https://api.example.com");
    });

    it("should allow overriding baseUrl", () => {
      const result = merge(defaults, "/users", {
        baseUrl: "https://other.com",
      });
      expect(result.baseUrl).toBe("https://other.com");
    });
  });
});
