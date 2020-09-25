# go-rest-api-base

## 環境

- go lang: v1.15.0

## ライブラリ

- [echo](https://github.com/labstack/echo) (v4.1.17) : REST APIフレームワーク
- [squirrel](https://github.com/Masterminds/squirrel) (v1.4.0) : SQLクエリビルダー
- [mysql driver](https://github.com/go-sql-driver/mysql) (v1.5.0) : MySQLドライバー
- [fresh](https://github.com/gravityblast/fresh) (v0.0.0) : アプリの起動（HotReload）ライブラリ
- [echo-swagger](https://github.com/swaggo/echo-swagger) (v1.0.0) : Echo用Swaggerライブラリ
- [koazee](https://github.com/wesovilabs/koazee) (v0.0.5) : コレクションライブラリ

## ツール

- [golang-migrate](https://github.com/golang-migrate/migrate)
  - [README](./migrations/README.md)

## URLs

- API: http://localhost:1323/
- Swagger UI: http://localhost:1323/swagger/index.html
- phpMyAdmin: http://localhost:18080/


## 事前準備

- [環境構築](readme/SETUP.md)
- [実装方針](readme/IMPLEMENT.md)

### .envの作成

- `.env.sample`をコピーして`.env`を作成する

### ビルド

- go.modに記載しているライブラリを全てダウンロードする
- `git clone`をした後にビルドをする必要がある
- `git pull`する時に新しいライブラリが追加されていた場合もビルドする必要がある

```
$ go build  
```

## アプリ起動

- 一般的な起動方法
```
$ go run main.go
```

- 上記の起動方法はコード変更時に再起動する必要がある
- HotReload対応ライブラリを使って起動する方法
- 以下のコマンドを実行する

```
$ fresh
```

- うまくいかない場合はライブラリをインストールする

```
$ go get -u github.com/pilu/fresh
```

以下のように出力されれば、アプリ起動は問題なし
```
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

- Swagger UIを開く: http://localhost:1323/swagger/index.html
- Message API（GET /message）を実行してデータが帰ってくれば問題なし

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

## Swagger UIのAPIの定義手順

- 参考URL: https://github.com/swaggo/echo-swagger
- controllerでgoコメントで以下を参考にAPIを定義する

```
// search messages.
// @Summary search messages
// @Description search messages
// @Accept  json
// @Produce  json
// @Param message_id query int false "メッセージID"
// @Param user_id query int false "ユーザID"
// @Param title query string false "タイトル"
// @Param message query string false "メッセージ"
// @Success 200 {array} models.Message
// @Failure 500 {object} error_handling.APIError
// @Router /message [get]
func NewMessageController(sqlHandler *database.SqlHandler) *MessageController {
  ...
}
```

- Swaggerを更新する

```
$ swag i
```

うまくいかない場合はライブラリをインストールする

```
$ go get -u github.com/swaggo/swag/cmd/swag
```
