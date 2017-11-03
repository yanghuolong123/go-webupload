#!/bin/bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
echo "当前路径："$CURDIR
echo "go原库路径："$OLDGOPATH

export GOPATH="$CURDIR:$OLDGOPATH"
echo "go新库路径:"$GOPATH

bee api webupload 

export GOPATH="$OLDGOPATH"

echo "go恢复后库路径:"$GOPATH
echo "finished"
