#!/bin/bash

dest_folder=./build
target=$dest_folder/pseudo-clnt

rm -f ./build/*

# to build debug version, no optimization
# use: $ go build -gcflags=all="-N -l"
go build -gcflags=all="-N -l" -o $target

