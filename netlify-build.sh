#!/bin/bash
set -euo pipefail

# Fetch tags since Netlify doesn't
git fetch --tags https://github.com/syncthing/docs.git

# Install RClone
mkdir -p ~/bin
pushd ~/bin
curl -O https://downloads.rclone.org/rclone-current-linux-amd64.zip
unzip rclone-current-linux-amd64.zip
mv rclone-*-linux-amd64/rclone .
popd

# Download pre-rendered / old versions of docs.
mkdir -p _build
~/bin/rclone sync spaces:syncthing-docs-rendered/ _build/html/
go run _script/lsver.go _build/html > _build/html/versions.json

# Build the current set of docs.
make html
make man
mv _build/man _build/html

