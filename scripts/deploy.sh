#! /bin/sh -eux

cd `dirname $0`

sh ./prepush.sh

cd ../client
npm run build

cd ../server
gcloud app deploy --quiet app.yaml backend.yaml modules/dispatch.yaml --version=from-pc