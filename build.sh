#!/usr/bin/env bash

# project_path=$(cd `dirname $0`; pwd)
# project_path=$BUILD_ROOT/bin
project_path=${GOPATH}/bin

go build -o "${project_path}/ucgo"

if [ $? != 0 ]; then
    printf "Compiling project failed\n"
    printf "GOPATH: "${GOPATH}"\n"
    exit 1
else
    printf "Compiling project successed\n"
fi

# mv ${project_path}/ucgo $BUILD_ROOT/bin

exit 0