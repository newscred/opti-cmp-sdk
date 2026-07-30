const scopes = Array.from(new Set(["js", "py"])).toSorted();

export default {
  extends: ["@commitlint/config-conventional"],
  rules: {
    "scope-enum": [2, "always", scopes],
  },
};
