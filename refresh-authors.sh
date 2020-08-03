#!/bin/sh
set -euo pipefail

# "git contributors" is https://github.com/calmh/git-contributors

git contributors --read-authors=AUTHORS --authors > authors-new
grep \# AUTHORS > authors-hdr
cat authors-hdr authors-new > AUTHORS
rm authors-hdr authors-new

git contributors --read-authors=AUTHORS --names --geekrank | sed 's/_/\\_/g' > thanks.txt
