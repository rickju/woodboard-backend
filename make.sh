#!/bin/bash

dest_folder=./build
target=$dest_folder/bknd-svr
proto_path=../proto

rm -f ./build/*

protoc -I . --go_out=plugins=grpc:$dest_folder --proto_path=$proto_path $proto_path/vnc.proto


# build debug version, no optimization
# $ go build -gcflags=all="-N -l"


# build vnc2.pb.go 
go build -gcflags=all="-N -l" ./build 

#
go build  -o $target

# 
# go run .
