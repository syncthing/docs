#!/bin/bash
set -euo pipefail

# Download pre-rendered / old versions of docs if this is a production (main
# branch) build.
if [[ ${CONTEXT:-} == "production" ]]; then
    mkdir -p _build
    git clone https://github.com/syncthing/docs-pre-rendered.git _build/html
    rm -rf _build/html/.git
fi

make html
make man
mv _build/man _build/html
