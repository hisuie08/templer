#!/usr/bin/env bash
set -euo pipefail

BIN="templer"
VERSION="${VERSION:-dev}"

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
    -ldflags="-s -w -X main.version=$VERSION" \
    -o "dist/${BIN}_${VERSION}_${GOOS}_${GOARCH}${ext}" \
    ./cmd/$BIN
done