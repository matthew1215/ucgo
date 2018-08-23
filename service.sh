#!/usr/bin/env bash

# project_path=$(cd `dirname $0`; pwd)
# project_path=$BUILD_ROOT
project_path=${GOPATH}/bin
processStr=${project_path}/ucgo

#if [ ! -n "ls ${processStr} >/dev/null 2>&1" ];then
if [ ! -f ${processStr} ];then
    printf "${processStr} is not found, please run build.sh first\n"
    exit 1
fi

while true
do
    process=`ps aux | grep ${processStr} | grep -v grep`;
    if [ "$process" == "" ]; then
        nohup ${project_path}/ucgo > /dev/null 2>&1 &
        if [ $? != 0 ]; then
            printf "process start failed\n"
            exit 2
        else
            printf "process start success\n"
            exit 0
        fi
        break;
    else
        pid=`ps aux | grep ${processStr} | grep -v grep | awk '{print $2}'`
        kill -2 ${pid}
        if [ $? != 0 ]; then
            printf "kill ${pid} failed\n"
            exit 3
        fi
        nohup ${project_path}/ucgo > /dev/null 2>&1 &
          if [ $? != 0 ]; then
            printf "process reload failed\n"
            exit 4
          else
            printf "process reload success\n"
            exit 0
          fi
        break;
    fi
done