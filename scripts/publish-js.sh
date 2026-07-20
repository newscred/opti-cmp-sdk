#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.."

npx release-please github-release \
  --repo-url=newscred/opti-cmp-sdk \
  --token="$(gh auth token)" \
  --target-branch=main

REMOTE=$(git remote -v | grep 'newscred.*push' | awk '{print $1}')
git pull "$REMOTE" main --tags

LATEST_TAG=$(git describe --tags --match "js-*" --abbrev=0)
VERSION="${LATEST_TAG#js-}"
echo "Latest tag: $LATEST_TAG (version $VERSION)"

if [ -n "$(npm view "opti-cmp@$VERSION" version 2>/dev/null)" ]; then
  echo "opti-cmp@$VERSION already published, skipping publish"
else
  git checkout "$LATEST_TAG"
  pnpm --filter=opti-cmp publish --no-git-checks
fi

git checkout main
