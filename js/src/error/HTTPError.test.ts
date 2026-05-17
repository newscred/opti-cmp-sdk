import { describe, expect, it } from "vitest";

import { HTTPError } from "./index.js";

describe("HTTPError", () => {
  describe("constructor", () => {
    it("should create error with message and status", () => {
      const error = new HTTPError("Not Found", 404, {});
      expect(error.message).toBe("Not Found");
      expect(error.status).toBe(404);
      expect(error.name).toBe("HTTPError");
    });

    it("should store response data", () => {
      const data = { code: "NOT_FOUND", error: "Resource not found" };
      const error = new HTTPError("Not Found", 404, { data });
      expect(error.data).toEqual(data);
    });

    it("should store response headers", () => {
      const headers = {
        "content-type": "application/json",
        "x-request-id": "abc123",
      };
      const error = new HTTPError("Error", 500, { headers });
      expect(error.headers).toEqual(headers);
    });

    it("should store request information", () => {
      const request = {
        headers: { accept: "application/json" },
        method: "GET",
        url: "https://api.example.com/users/123",
      };
      const error = new HTTPError("Not Found", 404, { request });
      expect(error.request).toEqual(request);
    });

    it("should be instanceof Error", () => {
      const error = new HTTPError("Error", 500, {});
      expect(error).toBeInstanceOf(Error);
      expect(error).toBeInstanceOf(HTTPError);
    });
  });
});
