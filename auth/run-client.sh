#!/bin/bash
set -x
go build -o ../bin/client ./
if [ $? -eq 0 ]; then
  ../bin/client;
fi;
