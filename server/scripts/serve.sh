#! /bin/sh -eux

cd `dirname $0`

# データストアエミュレータの設定
# `gcloud beta emulators datastore env-init` の実行結果を貼る
# https://cloud.google.com/sdk/gcloud/reference/beta/emulators/datastore/env-init?hl=ja
export DATASTORE_DATASET=hoge-hoge-123456789
export DATASTORE_EMULATOR_HOST=localhost:8081
export DATASTORE_EMULATOR_HOST_PATH=localhost:8081/datastore
export DATASTORE_HOST=http://localhost:8081
export DATASTORE_PROJECT_ID=hoge-hoge-123456789

cd ./../modules/backend
# TODO: gin を使って live-reloading したいが下記issueと同様のエラーが出る
# https://github.com/codegangsta/gin/issues/152
go run main.go