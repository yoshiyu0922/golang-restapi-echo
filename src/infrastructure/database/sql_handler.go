package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
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

/**
トランザクションを実行する
*/
func (sqlHandler *SqlHandler) Transaction(txFunc func(*sql.Tx) error) error {
	tx, err := sqlHandler.Conn.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	err = txFunc(tx)
	defer func() {
		if p := recover(); p != nil {
			// panic が発生したとき
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	return err
}
