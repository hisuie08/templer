#!/usr/bin/env bash
set -euo pipefail

BIN="templer"

mkdir -p dist

# 対応プラットフォーム
platforms=(
  "linux amd64"
  "linux arm64"
  "windows amd64"
)

for platform in "${platforms[@]}"; do
  read -r GOOS GOARCH <<< "$platform"

  ext=""
  if [ "$GOOS" = "windows" ]; then
    ext=".exe"
  fi

  echo "Building $GOOS/$GOARCH..."

  CGO_ENABLED=0 \
  GOOS=$GOOS GOARCH=$GOARCH \
  go build \
    -ldflags="-s -w -X templer/cmd.version=${VERSION}" \
    -o "dist/${BIN}_${GOOS}_${GOARCH}${ext}"
done