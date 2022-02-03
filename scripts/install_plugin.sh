#!/bin/sh -e

# Copied w/ love from the excellent hypnoglow/helm-s3

if [ -n "${HELM_OUTDATED_DEPENDENCIES_PLUGIN_NO_INSTALL_HOOK}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing helm-charts-plugin v${version} ..."

url=""
if [ "$(uname)" = "Darwin" ]; then
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_darwin_amd64.tar.gz"
elif [ "$(uname)" = "Linux" ] ; then
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_linux_amd64.tar.gz"
else
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_windows_amd64.tar.gz"
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
mv "releases/v${version}/bin/helm-file-utils" "bin/helm-file-utils" || \
    mv "releases/v${version}/bin/helm-file-utils.exe" "bin/helm-file-utils"