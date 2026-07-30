# Contributing

Thanks for your interest in contributing to `opti-cmp-sdk`!

## Development setup

This is a [pnpm](https://pnpm.io/) monorepo. You'll need Node.js `>=20` and pnpm.
The Python SDK in `py/` additionally needs [uv](https://docs.astral.sh/uv/) and
Python `>=3.10`.

```bash
pnpm install
uv sync --directory py
```

Common commands (run from the repo root):

| Command                        | Description               |
| ------------------------------ | ------------------------- |
| `pnpm build:js`                | Build the JS SDK          |
| `pnpm test:js`                 | Run the JS test suite     |
| `pnpm build:py`                | Build the Python SDK      |
| `pnpm test:py`                 | Run the Python test suite |
| `pnpm typecheck:py`            | Type-check the Python SDK |
| `pnpm lint:py`                 | Lint the Python SDK       |
| `pnpm exec eslint .`           | Lint                      |
| `pnpm exec prettier --check .` | Check formatting          |

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
Generated files carry an `Auto-generated - DO NOT EDIT` header and must not be
hand-edited. These include:

- `specification/*.json` — the split/sorted OpenAPI spec snapshot
- `js/src/types/*` — generated TypeScript types and endpoint definitions
- `js/src/plugins/register-api-endpoints/routes.json` — the runtime routing table
- `py/src/opti_cmp/_generated/*` — generated Python schema types, namespaces and
  the routing table

Both SDKs are generated from the same artifacts in `specification/`, so they
expose the same namespaces, endpoints and schema type names.

### Regenerating

```bash
# Fetch the latest spec and regenerate endpoint names, then regenerate each SDK:
pnpm generate:js
pnpm generate:py
```

Broken down:

- `pnpm generate` — fetches the remote OpenAPI spec into `specification/` and
  derives endpoint names.
- `pnpm generate:js` — runs `pnpm generate`, then regenerates the JS types,
  schema exports, and endpoints.
- `pnpm generate:py` — runs `pnpm generate`, then regenerates the Python schema
  types and endpoints.

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
3. Ensure the build, tests, lint, and format checks pass for whichever SDKs you
   touched.
4. Open a PR with a clear description. CI runs across Node 20, 22, and 24, and
   Python 3.10 through 3.13.

## Releases

Releases are managed by [release-please](https://github.com/googleapis/release-please)
and published to the public npm and PyPI registries. `pnpm release` opens a
release PR covering every component at once; `pnpm publish:js` and
`pnpm publish:py` push the corresponding tag to its registry. Maintainers: see
`scripts/release.sh` and `scripts/publish.sh` for details.
