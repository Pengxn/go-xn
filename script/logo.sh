#!/usr/bin/env bash

set -e

CURRENT_DIR="$(cd "$(dirname $0)"; pwd)"
ROOT_DIR=$CURRENT_DIR/..

cd $ROOT_DIR

function _sips() {
    sips -z 16 16     $1 --out $2/icon_16x16.png
    sips -z 32 32     $1 --out $2/icon_16x16@2x.png
    sips -z 32 32     $1 --out $2/icon_32x32.png
    sips -z 64 64     $1 --out $2/icon_32x32@2x.png
    sips -z 64 64     $1 --out $2/icon_64x64.png
    sips -z 128 128   $1 --out $2/icon_64x64@2x.png
    sips -z 128 128   $1 --out $2/icon_128x128.png
    sips -z 256 256   $1 --out $2/icon_128x128@2x.png
    sips -z 256 256   $1 --out $2/icon_256x256.png
    sips -z 512 512   $1 --out $2/icon_256x256@2x.png
    sips -z 512 512   $1 --out $2/icon_512x512.png
    sips -z 1024 1024 $1 --out $2/icon_512x512@2x.png
}

function _magick() {
    magick $1 -resize 16x16     $2/icon_16x16.png
    magick $1 -resize 32x32     $2/icon_16x16@2x.png
    magick $1 -resize 32x32     $2/icon_32x32.png
    magick $1 -resize 64x64     $2/icon_32x32@2x.png
    magick $1 -resize 64x64     $2/icon_64x64.png
    magick $1 -resize 128x128   $2/icon_64x64@2x.png
    magick $1 -resize 128x128   $2/icon_128x128.png
    magick $1 -resize 256x256   $2/icon_128x128@2x.png
    magick $1 -resize 256x256   $2/icon_256x256.png
    magick $1 -resize 512x512   $2/icon_256x256@2x.png
    magick $1 -resize 512x512   $2/icon_512x512.png
    magick $1 -resize 1024x1024 $2/icon_512x512@2x.png
}

# build_icons build icons by `sips` or `magick` command
# $1: original image name
# $2: destination directory name
function build_icons() {
    oriImage=script/assets/$1.png
    desDir=build/icons/$2
    # create destination directory if not exist
    if [ ! -d "$desDir" ]; then
        mkdir -p $desDir
    fi

    if [ $(uname) == "Darwin" ]; then
        _sips $oriImage $desDir # resize image by sips for macOS
    else
        _magick $oriImage $desDir # resize image by magick
    fi

    # convert image png to ico file
    convert $desDir/icon_16x16.png $desDir/icon_32x32.png $desDir/icon_64x64.png build/icons/$1.ico

    macosIconsetDir=build/icons/macos
    # create macos icons directory if not exist
    if [ ! -d "$macosIconsetDir" ]; then
        mkdir -p $macosIconsetDir
    fi
    mv $desDir/ $macosIconsetDir/icons.iconset
    # convert to icns file
    if [ $(uname) == "Darwin" ]; then
        iconutil -c icns $macosIconsetDir/icons.iconset -o build/icons/$1.icns
    fi

    rm -rf $macosIconsetDir $desDir
}

# detect if necessary command is installed
if test ! $(which convert); then
    echo 'ImageMagic tool (convert) is not installed'
    exit 1
else
    build_icons logo-white white
    build_icons logo-transparent transparent
fi
