#!/bin/sh -e

if [ -n "${HELM_LINTER_PLUGIN_NO_INSTALL_HOOK}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

# shellcheck disable=SC2002
version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing helm-file-utils v${version} ..."

url=""
if [ "$(uname)" = "Darwin" ]; then
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_darwin_amd64.tar.gz"
elif [ "$(uname)" = "Linux" ] ; then
    if [ "$(uname -m)" = "aarch64" ] || [ "$(uname -m)" = "arm64" ]; then
        url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_linux_arm64.tar.gz"
    else
        url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_linux_amd64.tar.gz"
    fi
else
    url="https://github.com/true-north-engineering/helm-file-utils/releases/download/v${version}/helm-file-utils_${version}_windows_amd64.tar.gz"
fi

echo "$url"

mkdir -p "bin"
mkdir -p "releases/v${version}"

# Download with curl if possible.
if [ -x "$(which curl 2>/dev/null)" ]; then
    curl -sSL "${url}" -o "releases/v${version}.tar.gz"
else
    wget -q "${url}" -O "releases/v${version}.tar.gz"
fi
tar xzf "releases/v${version}.tar.gz" -C "releases/v${version}"
mv "releases/v${version}/file-utils" "bin/file-utils" || \
    mv "releases/v${version}/file-utils.exe" "bin/file-utils"
mv "releases/v${version}/plugin.yaml" .
mv "releases/v${version}/README.md" .

echo "helm-file-utils ${version} is correctly installed."
echo
echo "See https://github.com/true-north-engineering/helm-file-utils for help getting started."
echo "Happy Helming day!"