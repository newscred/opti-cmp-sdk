import js from "@eslint/js";
import perfectionist from "eslint-plugin-perfectionist";
import { defineConfig, globalIgnores } from "eslint/config";
import globals from "globals";
import tseslint from "typescript-eslint";

export default defineConfig([
  globalIgnores(["js/lib", "**/node_modules", "py"]),
  {
    extends: [js.configs.recommended, tseslint.configs.recommended],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
      parser: tseslint.parser,
      parserOptions: {
        tsconfigRootDir: import.meta.dirname,
      },
    },
    rules: {
      "@typescript-eslint/no-unused-vars": [
        "error",
        {
          argsIgnorePattern: "^_",
        },
      ],
    },
  },
  {
    extends: [perfectionist.configs["recommended-natural"]],
    ignores: [
      "js/src/types/endpoints.ts",
      "js/src/types/openapi.d.ts",
      "js/src/types/params.ts",
      "js/src/types/schema.ts",
      "js/src/types/schema-public.ts",
    ],
  },
]);
