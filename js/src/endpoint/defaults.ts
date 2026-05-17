import type { EndpointDefaults } from "./types.js";

export const DEFAULTS: EndpointDefaults = {
  baseUrl: "https://api.cmp.optimizely.com/v3",
  headers: {
    accept: "application/json",
    "user-agent": `opti-cmp.js`,
  },
  method: "GET",
  url: "",
};
