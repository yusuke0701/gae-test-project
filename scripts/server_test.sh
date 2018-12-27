#! /bin/sh -eux

cd `dirname $0`

# import文の整理とフォーマッターの適用
goimports -w .

# コードの静的チェック
go vet ../...

# lintの実行
go list ../... | xargs golint -set_exit_status

# test
go list ../... | xargs goapp test