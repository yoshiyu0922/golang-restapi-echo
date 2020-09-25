# golang-migrate

### 公式サイト

- https://github.com/golang-migrate/migrate

- [CLI](https://github.com/golang-migrate/migrate#cli-usage) を使って、マイグレーションを行う

- [インストール方法](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### SQLファイルの命名規則

- YYYYMMDDNN_(create/insert)\_table\_(テーブル名).(up/down).sql
    - up: テーブル作成 or データ登録する
    - down: up.sqlに対する取り消しクエリ

### マイグレーションの実行

- マイグレーションを実行するには以下のコマンドを実行する
- 以下のコマンドは全てのsqlを実行してマイグレーションを実行する
- 実行後に `schema_migrations` テーブルが作成され実行履歴が保持される
  - 一度実行したsqlはスキップされる
- 途中までマイグレーションをする場合は `up` の後に数値（ファイル数）を入れる
- 取り消しをしたい場合は `up` を `down` に変更する。履歴を元に戻したり、down.sqlを実行したする

```
migrate -source file:migrations/ -database 'mysql://gouser:password@tcp(localhost:3306)/go_sample' up
```

