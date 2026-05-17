import { describe, expect, it, vi } from "vitest";

import type { APIClient } from "../../client/types.js";
import type { Response } from "../../request/types.js";
import type { Pagination } from "../../types/schema.js";

import paginationPlugin from "./index.js";

type PaginatedResponse = Response<{ pagination: Pagination }>;

describe("paginationPlugin", () => {
  function createMockClient(): APIClient {
    return {
      request: vi.fn(),
    } as unknown as APIClient;
  }

  describe("hasNextPage", () => {
    it("returns true when pagination.next exists", () => {
      const client = createMockClient();
      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: {
            next: "https://api.example.com/items?page=2",
            previous: null,
          },
        },
        headers: {},
        status: 200,
        url: "",
      };

      expect(client.hasNextPage(response)).toBe(true);
    });

    it("returns false when pagination.next is null", () => {
      const client = createMockClient();
      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: { next: null, previous: null },
        },
        headers: {},
        status: 200,
        url: "",
      };

      expect(client.hasNextPage(response)).toBe(false);
    });
  });

  describe("hasPreviousPage", () => {
    it("returns true when pagination.previous exists", () => {
      const client = createMockClient();
      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: {
            next: null,
            previous: "https://api.example.com/items?page=1",
          },
        },
        headers: {},
        status: 200,
        url: "",
      };

      expect(client.hasPreviousPage(response)).toBe(true);
    });

    it("returns false when pagination.previous is null", () => {
      const client = createMockClient();
      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: { next: null, previous: null },
        },
        headers: {},
        status: 200,
        url: "",
      };

      expect(client.hasPreviousPage(response)).toBe(false);
    });
  });

  describe("getNextPage", () => {
    it("calls client.request with next URL", async () => {
      const client = createMockClient();
      const mockResponse: PaginatedResponse = {
        data: { pagination: { next: null, previous: null } },
        headers: {},
        status: 200,
        url: "https://api.example.com/items?page=2",
      };
      vi.mocked(client.request).mockResolvedValue(mockResponse);

      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: {
            next: "https://api.example.com/items?page=2",
            previous: null,
          },
        },
        headers: {},
        status: 200,
        url: "",
      };

      await client.getNextPage(response);

      expect(client.request).toHaveBeenCalledWith(
        "https://api.example.com/items?page=2",
      );
    });

    it("throws HTTPError when no next page", async () => {
      const client = createMockClient();
      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: { next: null, previous: null },
        },
        headers: {},
        status: 200,
        url: "",
      };

      await expect(client.getNextPage(response)).rejects.toMatchObject({
        message: "No next page available",
        name: "HTTPError",
      });
    });
  });

  describe("getPreviousPage", () => {
    it("calls client.request with previous URL", async () => {
      const client = createMockClient();
      const mockResponse: PaginatedResponse = {
        data: { pagination: { next: null, previous: null } },
        headers: {},
        status: 200,
        url: "https://api.example.com/items?page=1",
      };
      vi.mocked(client.request).mockResolvedValue(mockResponse);

      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: {
            next: null,
            previous: "https://api.example.com/items?page=1",
          },
        },
        headers: {},
        status: 200,
        url: "",
      };

      await client.getPreviousPage(response);

      expect(client.request).toHaveBeenCalledWith(
        "https://api.example.com/items?page=1",
      );
    });

    it("throws HTTPError when no previous page", async () => {
      const client = createMockClient();
      paginationPlugin(client);

      const response: PaginatedResponse = {
        data: {
          pagination: { next: null, previous: null },
        },
        headers: {},
        status: 200,
        url: "",
      };

      await expect(client.getPreviousPage(response)).rejects.toMatchObject({
        message: "No previous page available",
        name: "HTTPError",
      });
    });
  });
});
