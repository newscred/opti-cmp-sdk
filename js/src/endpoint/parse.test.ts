import { describe, expect, it } from "vitest";

import type { EndpointDefaults } from "./types.js";

import { parse } from "./parse.js";

describe("parse", () => {
  const baseDefaults: EndpointDefaults = {
    baseUrl: "https://api.example.com",
    headers: { accept: "application/json" },
    method: "GET",
    url: "/users",
  };

  describe("URL construction", () => {
    it("should prepend baseUrl to relative URLs", () => {
      const result = parse(baseDefaults);
      expect(result.url).toBe("https://api.example.com/users");
    });

    it("should not prepend baseUrl to absolute URLs", () => {
      const result = parse({ ...baseDefaults, url: "https://other.com/path" });
      expect(result.url).toBe("https://other.com/path");
    });

    it("should expand URL template variables", () => {
      const result = parse({
        ...baseDefaults,
        id: "123",
        url: "/users/{id}",
      });
      expect(result.url).toBe("https://api.example.com/users/123");
    });

    it("should expand multiple URL template variables", () => {
      const result = parse({
        ...baseDefaults,
        postId: "post-2",
        url: "/users/{userId}/posts/{postId}",
        userId: "user-1",
      });
      expect(result.url).toBe(
        "https://api.example.com/users/user-1/posts/post-2",
      );
    });
  });

  describe("query parameters", () => {
    it("should add query parameters for GET requests", () => {
      const result = parse({
        ...baseDefaults,
        limit: "10",
        method: "GET",
        page: "1",
      });
      expect(result.url).toContain("page=1");
      expect(result.url).toContain("limit=10");
    });

    it("should handle array query parameters", () => {
      const result = parse({
        ...baseDefaults,
        method: "GET",
        tags: ["a", "b", "c"],
      });
      expect(result.url).toContain("tags=a");
      expect(result.url).toContain("tags=b");
      expect(result.url).toContain("tags=c");
    });

    it("should URL-encode query parameter values", () => {
      const result = parse({
        ...baseDefaults,
        method: "GET",
        query: "hello world",
      });
      expect(result.url).toContain("query=hello%20world");
    });
  });

  describe("request body", () => {
    it("should set body for POST requests", () => {
      const result = parse({
        ...baseDefaults,
        email: "john@example.com",
        method: "POST",
        name: "John",
      });
      expect(result.body).toBe(
        JSON.stringify({ email: "john@example.com", name: "John" }),
      );
      expect(result.headers["content-type"]).toBe(
        "application/json; charset=utf-8",
      );
    });

    it("should use explicit body if provided", () => {
      const body = { custom: "data" };
      const result = parse({
        ...baseDefaults,
        body,
        method: "POST",
      });
      expect(result.body).toBe(JSON.stringify(body));
    });

    it("should not include URL variables in body", () => {
      const result = parse({
        ...baseDefaults,
        id: "123",
        method: "POST",
        name: "John",
        url: "/users/{id}",
      });
      expect(result.url).toBe("https://api.example.com/users/123");
      expect(result.body).toBe(JSON.stringify({ name: "John" }));
    });
  });
});
