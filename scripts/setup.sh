#! /bin/sh -eux

cd `dirname $0`

export GO111MODULE=off
go get -u github.com/codegangsta/gin
go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/lint/golint

cd ../client
npm install

cd ../server
export GO111MODULE=on
go mod download