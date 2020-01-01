# GAEの動作確認用リポジトリ
[![CircleCI](https://circleci.com/gh/yusuke0701/gae-test-project.svg?style=svg)](https://circleci.com/gh/yusuke0701/gae-test-project)

### 主な構成
##### Server
GAE/Go(1.13) + Datastore
##### Front
Vue.js(2.6)

### 開発環境の構築
1. CloudSDKのインストール
    1. 公式URL: https://cloud.google.com/sdk/
1. golangのインストール
    1. 参考URL: https://golang.org/doc/install
1. npmとnodejsのインストール
    1. 参考URL: https://qiita.com/seibe/items/36cef7df85fe2cefa3ea
1. `scripts/setup.sh`の実行

### 注意点
1. push前に`scripts/prepush.sh`を実行

### 動作確認用URL一覧
URL: https://hoge-hoge-123456789.appspot.com/index.html