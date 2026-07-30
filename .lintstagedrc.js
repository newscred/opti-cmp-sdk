export default {
  "*.{js,jsx,ts,tsx}": ["prettier --write", "eslint --cache --fix"],
  "*.{json,md,yml}": "prettier --write",
  "*.{ts,tsx}": () => [
    "tsc --noEmit",
    "tsc --noEmit -p js",
    "tsc --noEmit -p js/tsconfig.examples.json",
  ],
  "py/**/*.py": (files) => {
    // Only mypy needs whole-project scope; ruff writes, so it gets the staged
    // files only — `.` would reformat unstaged files and leave the rewrites out
    // of the commit.
    const paths = files.map((file) => JSON.stringify(file)).join(" ");
    return [
      `uv run --directory py ruff format ${paths}`,
      `uv run --directory py ruff check --fix ${paths}`,
      "uv run --directory py mypy",
    ];
  },
};
