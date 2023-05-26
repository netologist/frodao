package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/netologist/frodao"
)

var SESSION *sql.DB

func SetSession(session *sql.DB) {
	SESSION = session
}

func Connect(dsn string, options ...frodao.Option) (err error) {
	if SESSION, err = sql.Open("postgres", dsn); err != nil {
		return err
	}
	frodao.SetOptions(SESSION, options)
	return nil
}

func Close() {
	_ = SESSION.Close()
}
