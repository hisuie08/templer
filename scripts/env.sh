#!/usr/bin/env bash
set -euo pipefail

export GOFLAGS="-trimpath"
export CGO_ENABLED=1
export VERSION="$(git tag --sort=-v:refname -l "v*" | head -1)"