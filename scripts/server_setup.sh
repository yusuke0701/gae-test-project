#! /bin/sh -eux

go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/lint/golint
go get -u github.com/golang/dep/cmd/dep

dep ensure