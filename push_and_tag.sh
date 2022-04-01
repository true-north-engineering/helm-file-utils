#!/bin/sh -e

# shellcheck disable=SC2002
version="$(cat plugin.yaml | grep "version:" | cut -d '"' -f 2)"
echo "-----> Tag version is: v${version} ..."

echo "Commit message?"
read message
git add .
git commit -m "$message"
git push
git tag "v$version"
git push --tags