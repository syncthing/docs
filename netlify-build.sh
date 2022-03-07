#!/bin/bash
set -euo pipefail

# Fetch tags since Netlify doesn't
git fetch --tags https://github.com/syncthing/docs.git

# Download pre-rendered / old versions of docs.
mkdir -p _build
git clone --depth 1 https://github.com/syncthing/docs-pre-rendered.git _build/html
rm -rf _build/html/.git
go run _script/lsver.go _build/html > _build/html/versions.json

# Build the current set of docs.
make html
make man
mv _build/man _build/html
