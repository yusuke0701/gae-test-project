#! /bin/sh -eux

cd `dirname $0`
cd ../server

goimports -w ./..
golint ./..