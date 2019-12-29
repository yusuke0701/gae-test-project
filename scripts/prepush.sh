#! /bin/sh -eux

cd `dirname $0`

cd ../client
npm run lint

cd ../server
goimports -w ./..
golint ./..
go mod tidy