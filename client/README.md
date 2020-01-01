# client

## セットアップ方法
下記ドキュメントを参照
https://github.com/yusuke0701/gae-test-project/blob/master/README.md#%E9%96%8B%E7%99%BA%E7%92%B0%E5%A2%83%E3%81%AE%E6%A7%8B%E7%AF%89

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
* 本番環境変数は `.env` に置く
* `json-server` のデータは `mock`フォルダに置く