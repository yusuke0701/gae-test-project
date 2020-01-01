# server

### 主な言語/ライブラリ
##### DB
[Cloud Datastore](https://cloud.google.com/datastore/docs/concepts/overview?hl=ja)
##### httpサーバー
[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
##### live-reloader
[github.com/codegangsta/gin](https://github.com/codegangsta/gin)

### セットアップ方法
下記ドキュメントを参照
https://github.com/yusuke0701/gae-test-project/blob/master/README.md#%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%81%AE%E6%A7%8B%E7%AF%89

### 開発用サーバーの立ち上げ方
```
cd ../client
npm run serve

cd ../server
./scripts/run_datastore.sh
./scripts/serve.sh
```