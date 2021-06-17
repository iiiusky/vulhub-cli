#!/usr/bin/env bash

for FILE in dist/vulhub-cli*/*; do
    if [ ${FILE##*.} != "sha256" ];then
        du -sh ${FILE}
        upx ${FILE}
        du -sh ${FILE}
        echo "sum sha256"
        echo `shasum -a 256 ${FILE}`
        echo `shasum -a 256 ${FILE}` > ${FILE}*.sha256
    fi
done