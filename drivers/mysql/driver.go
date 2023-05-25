package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var SESSION *sql.DB

func SetSession(session *sql.DB) {
	SESSION = session
}

func Connect(dsn string) (err error) {
	if SESSION, err = sql.Open("mysql", dsn); err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = SESSION.Close()
}
