#!/bin/bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
echo "当前路径："$CURDIR
echo "go原库路径："$OLDGOPATH

export GOPATH="$CURDIR:$OLDGOPATH:/var/work/work_golang/mylib"
echo "go新库路径:"$GOPATH

cd ./src/webupload
bee run 

export GOPATH="$OLDGOPATH"

echo "go恢复后库路径:"$GOPATH
echo "finished"
