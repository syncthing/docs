#!/bin/sh
set -euo pipefail

# "git contributors" is https://github.com/calmh/git-contributors

git contributors --read-authors=AUTHORS --authors > authors-new
grep \# AUTHORS > authors-hdr
cat authors-hdr authors-new > AUTHORS
rm authors-hdr authors-new

git contributors --read-authors=AUTHORS --names --geekrank > thanks.txt

# Underscores need be escaped in RST.
sed -i '/_/ s//\\_/g' thanks.txt
