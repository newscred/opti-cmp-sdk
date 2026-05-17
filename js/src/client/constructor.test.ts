import { describe, expect, it, vi } from "vitest";

import type { APIClient, Options, Plugin } from "./types.js";

import { constructor } from "./constructor.js";

describe("constructor", () => {
  it("registers plugins in order", () => {
    const callOrder: string[] = [];

    const plugin1: Plugin = () => callOrder.push("plugin1");
    const plugin2: Plugin = () => callOrder.push("plugin2");
    const plugin3: Plugin = () => callOrder.push("plugin3");

    constructor([plugin1, plugin2, plugin3]);

    expect(callOrder).toEqual(["plugin1", "plugin2", "plugin3"]);
  });

  it("passes client and options to each plugin", () => {
    const pluginSpy = vi.fn();
    const options: Options = { baseUrl: "https://api.example.com" };

    const client = constructor([pluginSpy], options);

    expect(pluginSpy).toHaveBeenCalledWith(client, options);
  });

  it("allows plugins to modify client", () => {
    const plugin: Plugin = (client: APIClient) => {
      client.customMethod = () => "custom";
    };

    const client = constructor([plugin]) as APIClient & {
      customMethod: () => string;
    };

    expect(client.customMethod()).toBe("custom");
  });
});
