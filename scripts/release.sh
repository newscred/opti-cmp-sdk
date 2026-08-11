#!/usr/bin/env bash

# One command for both components: `release-please release-pr` reads the whole
# manifest, so it opens release PRs for every component at once.

set -euo pipefail

cd "$(dirname "$0")/.."

npx release-please release-pr \
  --repo-url=newscred/opti-cmp-sdk \
  --token="$(gh auth token)" \
  --target-branch=main
