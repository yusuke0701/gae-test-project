#! /bin/sh -eux

# GAEへデプロイするためのシェルスクリプト
# (例) ./scripts/deploy.sh -a hoge-hoge-123456789 -v from-pc

cd `dirname $0`

sh ./prepush.sh

while getopts a:v: OPT
do
	case $OPT in
		"a" ) PROJECT_ID="$OPTARG";;
		"v" ) VERSION="$OPTARG";;
		* ) echo "Usage: $CMDNAME -a=<project_id> -v=<version>"
			exit 1 ;;
	esac
done

cd ../client
npm run build

cd ../server
gcloud app deploy app.yaml backend.yaml modules/dispatch.yaml --project=$PROJECT_ID --version=$VERSION