#!/usr/bin/env bash
set -euo pipefail

git describe --tags --long --always > RELEASE || true
git describe --tags --exact-match > TAG || true

target="${1:-html}"
docker run --rm \
  -v "$(pwd):$(pwd)" \
  -w "$(pwd)" \
  docker.io/sphinxdoc/sphinx-latexpdf:latest \
  make "$target"
