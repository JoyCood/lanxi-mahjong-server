#!/bin/bash
export GOPATH=`pwd`
export GOARCH=amd64
export GOOS=linux
cd bin

go build -o server -ldflags "-w -s" ../src/server.go



