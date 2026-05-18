#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.."

npx release-please release-pr \
  --repo-url=newscred/opti-cmp-sdk \
  --token="$(gh auth token)" \
  --target-branch=main
