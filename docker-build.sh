#!/usr/bin/env bash
set -euo pipefail

target="${1:-html}"
docker run --rm \
  -v "$(pwd):$(pwd)" \
  -w "$(pwd)" \
  ghcr.io/kastelo/sphinx-build:latest \
  make "$target"
