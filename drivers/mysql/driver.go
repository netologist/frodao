package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/netologist/frodao"
)

var SESSION *sql.DB

func SetSession(session *sql.DB) {
	SESSION = session
}

func Connect(dsn string, options ...frodao.Option) (err error) {
	if SESSION, err = sql.Open("mysql", dsn); err != nil {
		return err
	}
	frodao.SetOptions(SESSION, options)
	return nil
}

func Close() error {
	return SESSION.Close()
}
