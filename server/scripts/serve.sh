#! /bin/sh -eux

cd `dirname $0`

cd ../server/modules/backend
gin run main.go