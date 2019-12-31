#! /bin/sh -eux

cd `dirname $0`

# データストアエミュレータの起動
# https://cloud.google.com/datastore/docs/tools/datastore-emulator?hl=ja
gcloud beta emulators datastore start