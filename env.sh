#!/bin/bash

OUTPUT=$1

if test -z "${OUTPUT}"; then
    echo "Usage: $0 <output file>"
    exit 1
fi

if [ ! -d "./bin" ]; then
    $(mkdir bin)
fi

PWD=$(pwd)
echo GOPATH=${GOPATH}:${PWD} >> ${OUTPUT}
