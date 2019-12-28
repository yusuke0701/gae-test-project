#! /bin/sh -eux

cd `dirname $0`

cd ../client
npm run build

cd ../server
gin run main.go