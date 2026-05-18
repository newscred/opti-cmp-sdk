import { defineConfig } from "tsdown";

export default defineConfig({
  dts: true,
  entry: ["./src/index.ts"],
  exports: {
    devExports: true,
  },
  format: {
    cjs: {
      deps: {
        alwaysBundle: () => true,
        neverBundle: ["is-mergeable-object"],
        skipNodeModulesBundle: false,
      },
    },
    esm: {},
    umd: {
      deps: {
        alwaysBundle: () => true,
        neverBundle: ["is-mergeable-object"],
        skipNodeModulesBundle: false,
      },
      outputOptions: {
        name: "OptiCMP",
      },
    },
  },
  outDir: "./lib",
  platform: "neutral",
});
