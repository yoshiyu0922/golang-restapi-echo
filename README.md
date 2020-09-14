# go-rest-api-base

## 事前準備

- [環境構築](./doc/SETUP.md)

## プロジェクト準備

### ローカルDBを構築

```
$ cd go-rest-api-base/docker
$ docker-compose up
```

テーブルを作成＆データを登録する

```
$ cd go-rest-api-base/docker
$ MYSQL_PWD=root mysql -u root -h 127.0.0.1 --port=3306 go_sample < ddl/init.sql
$ mysql -u root -proot -h 127.0.0.1 --port=3306 -D go_sample
mysql> SHOW TABLES FROM go_sample;
-> テーブル一覧が出ればOK
```

#### 備忘録

- DBクライアントから接続するためにはユーザの認証プラグインを変更する必要があるが、my.cnfで設定済み

### ビルド

- プロジェクトルートへ行き、ビルドをする
- go.modに記載しているライブラリを全てダウンロードする

```
$ go build  
```

### .envの作成

- `.env.sample`をコピーして`.env`を作成する

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

### アプリ起動（Hot Reload）

- コード変更した場合に自動的に再起動してくれる

```
$ fresh
```

- うまくいかない場合はライブラリをインストールする

```
$ go get -u github.com/pilu/fresh
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
