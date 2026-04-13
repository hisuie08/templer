#!/usr/bin/env bash
set -euo pipefail
source scripts/env.sh

mkdir -p build

echo "==> build templer"
go build \
  -buildvcs=false \
  -ldflags="-s -w -X templer/cmd.Version=${VERSION}" \
  -o build/templer
chmod +x build/templer
echo "✔ build/templer"