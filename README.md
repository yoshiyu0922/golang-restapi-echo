# go-rest-api-base

## 事前準備

- [環境構築](./doc/SETUP.md)

## Project

### moduleを初期化

```
$ go mod init example.com/rest-api-base  
```

### アプリ起動

```
$ go run helloWorld.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

### 疎通確認

以下のコマンドを実行して「Hello, World!」が表示されれば問題なし

```
$ curl -w "\n" "http://localhost:1323"  
Hello, World!
```

### Deploy

- TODO: クロスコンパイル設定を確認
- ビルドしてバイナリファイルを作成する

```
$ go build

$ ./rest-api-base

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.17
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```
