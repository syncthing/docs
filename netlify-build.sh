#!/bin/bash
set -euo pipefail

# Fetch tags since Netlify doesn't
git fetch --tags https://github.com/syncthing/docs.git

# Download pre-rendered / old versions of docs.
mkdir -p _build
curl https://rclone.org/install.sh | sudo bash
rclone sync spaces:syncthing-docs-rendered/ _build/html/
go run _script/lsver.go _build/html > _build/html/versions.json

# Build the current set of docs.
make html
make man
mv _build/man _build/html

