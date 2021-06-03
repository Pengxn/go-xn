#!/usr/bin/env bash

set -e

# Make sure the `rcedit` is installed
# https://github.com/electron/rcedit

GIT_VERSION=$(git describe --tags `git rev-list --tags --max-count=1`)

VERSION=$(echo ${GIT_VERSION#*v})

rcedit ./build/go-xn.exe \
        --set-icon "./web/icons/xn-02f.ico" \
        --set-version-string CompanyName "xn-02f Lab" \
        --set-version-string ProductName "xn-02f Lab" \
        --set-version-string FileDescription "✍ The platform for publishing and running your blog." \
        --set-version-string OriginalFilename "go-xn.exe" \
        --set-version-string InternalName "go-xn" \
        --set-version-string LegalCopyright "(c) All Go-xn Contributors. Zlib LICENSE." \
        --set-file-version ${VERSION} \
        --set-product-version ${VERSION} \
        --application-manifest ./script/windows/manifest.xml
