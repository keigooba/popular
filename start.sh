#!/bin/sh
go mod tidy
golangci-lint run ./...
GIT_VER=`git describe --tags`
go build --ldflags "-X main.version=${GIT_VER}"
