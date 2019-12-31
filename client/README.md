# client

## セットアップ方法
下記ドキュメントを参照
https://github.com/yusuke0701/gae-test-project/blob/master/README.md

### 開発用サーバーの立ち上げ方
```
npm run serve
```
vueの開発用サーバーが立ち上がります

([`vue-cli-service serve`](https://cli.vuejs.org/guide/cli-service.html#vue-cli-service-serve))
```
npm run mock
```
モックサーバーが立ち上がります

([`json-server`](https://github.com/typicode/json-server))
```
npm run serve:mock
```
上記のvueの開発用サーバーとモックサーバーが立ち上がります

### 注意点
* ローカル環境変数は `.env.dev` に置く
* `json-server` のデータは `mock`フォルダに置く