#!/usr/bin/env bash

# Publish one component's latest release. Everything except the "is it already
# published" probe and the publish command itself is the same for both, so the
# shared part lives here rather than in two copies that drift.

set -euo pipefail

COMPONENT="${1:-}"
case "$COMPONENT" in
js | py) ;;
*)
  echo "usage: $0 <js|py>" >&2
  exit 1
  ;;
esac

cd "$(dirname "$0")/.."

# Unscoped on purpose: release-please cuts the GitHub release for whichever
# release PRs have been merged, which is what we want regardless of which
# component we are about to push to a registry.
npx release-please github-release \
  --repo-url=newscred/opti-cmp-sdk \
  --token="$(gh auth token)" \
  --target-branch=main

REMOTE=$(git remote -v | grep 'newscred.*push' | awk '{print $1}')
git pull "$REMOTE" main --tags

LATEST_TAG=$(git describe --tags --match "$COMPONENT-*" --abbrev=0)
VERSION="${LATEST_TAG#"$COMPONENT"-}"
echo "Latest tag: $LATEST_TAG (version $VERSION)"

if [ "$COMPONENT" = js ]; then
  published=$([ -n "$(npm view "@optimizely/cmp-sdk@$VERSION" version 2>/dev/null)" ] && echo yes || echo no)
else
  published=$(curl -sf "https://pypi.org/pypi/opti-cmp/$VERSION/json" >/dev/null && echo yes || echo no)
fi

if [ "$published" = yes ]; then
  echo "opti-cmp-sdk@$VERSION already published, skipping publish"
else
  # Restore the branch even if the build or upload fails, so a failed run does
  # not strand the checkout on a detached HEAD.
  trap 'git checkout main' EXIT
  git checkout "$LATEST_TAG"
  if [ "$COMPONENT" = js ]; then
    pnpm --filter=@optimizely/cmp-sdk publish --no-git-checks
  else
    rm -rf py/dist
    uv build --project py
    uv publish py/dist/*
  fi
fi

git checkout main
