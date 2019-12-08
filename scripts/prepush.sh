#! /bin/sh -eux

cd `dirname $0`
cd ../src

goimports -w ./..
golint ./..