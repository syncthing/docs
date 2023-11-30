#!/bin/sh
set -euo pipefail

go install github.com/calmh/git-contributors@latest
git-contributors --read-authors=AUTHORS --authors > authors-new
grep \# AUTHORS > authors-hdr
cat authors-hdr authors-new > AUTHORS
rm authors-hdr authors-new

git-contributors --read-authors=AUTHORS --names --geekrank | sed 's/_/\\_/g' > thanks.txt
