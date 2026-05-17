import { describe, expect, it, vi } from "vitest";

import type { Plugin } from "./types.js";

import { factory } from "./factory.js";

describe("factory", () => {
  it("can be called with or without new keyword", () => {
    const Client = factory();
    const clientWithoutNew = Client();
    const clientWithNew = new Client();

    expect(clientWithoutNew.request).toBeDefined();
    expect(clientWithNew.request).toBeDefined();
  });

  it("passes options to client constructor", () => {
    const Client = factory();
    const client = Client({ baseUrl: "https://api.example.com" });

    expect(client.request.endpoint.DEFAULTS.baseUrl).toBe(
      "https://api.example.com",
    );
  });

  it("plugins method returns new factory with combined plugins", () => {
    const plugin1: Plugin = vi.fn();
    const plugin2: Plugin = vi.fn();

    const BaseClient = factory([plugin1]);
    const ExtendedClient = BaseClient.plugins([plugin2]);

    ExtendedClient();

    expect(plugin1).toHaveBeenCalled();
    expect(plugin2).toHaveBeenCalled();
  });

  it("registers plugins passed to factory", () => {
    const plugin: Plugin = vi.fn();
    const Client = factory([plugin]);

    Client();

    expect(plugin).toHaveBeenCalled();
  });
});
