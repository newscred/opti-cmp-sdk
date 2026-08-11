# Contributing

Thanks for your interest in contributing to `opti-cmp-sdk`!

## Development setup

This is a [pnpm](https://pnpm.io/) monorepo. You'll need Node.js `>=20` and pnpm.

```bash
pnpm install
```

Common commands (run from the repo root):

| Command                        | Description           |
| ------------------------------ | --------------------- |
| `pnpm build:js`                | Build the JS SDK      |
| `pnpm test:js`                 | Run the JS test suite |
| `pnpm exec eslint .`           | Lint                  |
| `pnpm exec prettier --check .` | Check formatting      |

## Commit messages

This repo uses [Conventional Commits](https://www.conventionalcommits.org/),
enforced by commitlint via a git hook. Examples:

```
feat(js): add pagination helper
fix(js): correct token refresh timing
chore: update dependencies
```

Husky + lint-staged run formatting and commit-message checks automatically on
commit.

## Code generation

Much of the SDK is **generated from Optimizely's hosted OpenAPI specification**.
Generated files carry a `// Auto-generated - DO NOT EDIT` header and must not be
hand-edited. These include:

- `specification/*.json` — the split/sorted OpenAPI spec snapshot
- `js/src/types/*` — generated TypeScript types and endpoint definitions
- `js/src/plugins/register-api-endpoints/routes.json` — the runtime routing table

### Regenerating

```bash
# Fetch the latest spec and regenerate endpoint names, then regenerate all JS artifacts:
pnpm generate:js
```

Broken down:

- `pnpm generate` — fetches the remote OpenAPI spec into `specification/` and
  derives endpoint names.
- `pnpm generate:js` — runs `pnpm generate`, then regenerates the JS types,
  schema exports, and endpoints.

The spec source URL is defined in `scripts/generate-specification.ts` and can be
overridden with the `OPENAPI_SPEC_URI` environment variable.

### Namespace names

Method namespaces (e.g. `client.campaign`, `client.task`) are derived from the
OpenAPI tags via a hand-maintained `TAG_TO_NAMESPACE` map in
[`scripts/generate-endpoint-names.ts`](./scripts/generate-endpoint-names.ts).
If the API adds a new tag, add a corresponding entry to this map — generation
will throw with `Namespace not found for tag: <tag>` otherwise. This is the one
part of the generation pipeline you are expected to edit by hand.

## Pull requests

1. Fork and create a feature branch.
2. Make your changes, adding or updating tests as appropriate.
3. Ensure `pnpm build:js`, `pnpm test:js`, lint, and format checks pass.
4. Open a PR with a clear description. CI runs across Node 20, 22, and 24.

## Releases

Releases are automated with [release-please](https://github.com/googleapis/release-please).
Merging its release PR tags the release and publishes `@optimizely/cmp-sdk` to npm.
