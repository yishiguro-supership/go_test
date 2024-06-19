# 3-1_no_di

# How to Use

## ローカルDB操作
- webアプリをrunする前に立ち上げておく必要がある
- dockerが入っている前提

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
