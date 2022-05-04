package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var SESSION *sql.DB

func SetSession(session *sql.DB) {
	SESSION = session
}

func Connect(dsn string) (err error) {
	if SESSION, err = sql.Open("postgres", dsn); err != nil {
		return err
	}
	return nil
}

func Close() {
	_ = SESSION.Close()
}
