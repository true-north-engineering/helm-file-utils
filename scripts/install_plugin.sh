#!/bin/sh -e

# Copied w/ love from the excellent hypnoglow/helm-s3

if [ -n "${HELM_OUTDATED_DEPENDENCIES_PLUGIN_NO_INSTALL_HOOK}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

version="$(cat plugin.yaml | grep "version:" | cut -d '"' -f 2)"
echo "Downloading and installing helm-charts-plugin v${version} ..."

project_name="$(cat .goreleaser.yml | grep "project_name:" | cut -d ' ' -f 2)"

url=""
if [ "$(uname)" = "Darwin" ]; then
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/${project_name}_darwin_amd64.tar.gz"
elif [ "$(uname)" = "Linux" ] ; then
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/${project_name}_linux_amd64.tar.gz"
else
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/${project_name}_windows_amd64.tar.gz"
fi

echo $url

mkdir -p "bin"
mkdir -p "releases/v${version}"

# Download with curl if possible.
if [ -x "$(which curl 2>/dev/null)" ]; then
    curl -sSL "${url}" -o "releases/v${version}.tar.gz"
else
    wget -q "${url}" -O "releases/v${version}.tar.gz"
fi
tar xzf "releases/v${version}.tar.gz" -C "releases/v${version}"
mv "releases/v${version}/bin/${project_name}" "bin/${project_name}" || \
    mv "releases/v${version}/bin/${project_name}.exe" "bin/${project_name}"