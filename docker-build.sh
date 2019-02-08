#!/usr/bin/env bash
set -euo pipefail

target="${1:-html}"
docker run --rm \
  -v "$(pwd):$(pwd)" \
  -w "$(pwd)" \
  kastelo/sphinx-build:latest \
  sphinx-versioning build . _build/html

