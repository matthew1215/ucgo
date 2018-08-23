#!/usr/bin/env bash

OLDGOPATH=${GOPATH}
export GOPATH=$(dirname "$(dirname "$PWD")")
# export GOPATH=${GOPATH//'?'/' '}
sh ./build.sh
if [ $? != 0 ]; then
    exit 1
fi
sh ./service.sh
export GOPATH=${OLDGOPATH}

