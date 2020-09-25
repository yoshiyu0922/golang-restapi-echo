package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(
	user string,
	password string,
	host string,
	port string,
	dbName string,
) *SqlHandler {
	// TODO: GCPに接続する場合は少しURLが異なるかもしれないので、その場合はENVで切り分ける
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	conn, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		panic(err)
	}
	if err := conn.Ping(); err != nil {
		panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
