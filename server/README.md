# server

## セットアップ方法
下記ドキュメントを参照
https://github.com/yusuke0701/gae-test-project/blob/master/README.md

### 開発用サーバーの立ち上げ方
```
cd ../client
npm run serve

cd ../server
./scripts/serve.sh
```
TODO: ローカルではDatastore,GCSへ接続できないので動かない。
下記ドキュメントを参考に作りたい
https://cloud.google.com/datastore/docs/tools/datastore-emulator?hl=ja