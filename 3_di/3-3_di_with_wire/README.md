# 3-3_di_with_wire

# How to Use

## ローカルDB操作
- webアプリをrunする前に立ち上げておく必要がある
- dockerが入っている前提

## wireのインストール
- make wireを実施するためにwireコマンドを使えるようにする
  - 書き込み先の.zshrcはお使いのシェルに合わせて適宜修正
```
go install github.com/google/wire/cmd/wire@latest

echo 'export GOPATH=$(go env GOPATH)' >>~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >>~/.zshrc
```

## wire_gen.goの出力
- 依存関係をまとめたファイル(`./injector/wire.go`)を元にwire_gen.goを出力する
  - 本テストコードではすでに出力済み
  - 試す場合、一度wire_gen.goを削除してから以下を実行
```
wire ./injector/
# Makefileにコマンドを作ってあるので以下でも可
make wire
```

### dynamodb start
- DB立ち上げ〜データの挿入
- 永続化はしていないため、stopするたびリセットされる
```
make dynamodb-start
```

### dynamodb stop
- DB終了
```
make dynamodb-stop
```

## webサーバ操作
- エンドポイントはlocalhost:3000
- 終了時はctrl+Cで
### run
```
make run
```

## curlでのリクエスト例
```
curl -L http://localhost:3000/          # healthcheck
curl -L http://localhost:3000/user/101  # uid=101のユーザ情報取得
curl -L http://localhost:3000/user/105  # uid=105のユーザ情報取得
curl -L http://localhost:3000/user/999  # uid=999のユーザ情報取得(存在しないため空)
```
