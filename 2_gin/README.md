# 2_gin

## Description
- golangのwebフレームワークであるginの基本的な使い方のテストコード
- サーバを起動してcurlコマンドを叩くことで得られるレスポンスとコードを見比べてginの理解を深めることを目的としている

## How to Use

### build and execute
```
make build
./myproject
```

```
curl http://localhost:3000/
```

### clean
```
make clean
```

### run
```
make run
```

### curlでのリクエスト例
```
curl "http://localhost:3000/"
curl "http://localhost:3000/hello"
curl "http://localhost:3000/hello/Alice"
curl "http://localhost:3000/hello?name=Bob"
curl "http://localhost:3000/hello?name=Cathy&message=HowAreYou"
```
