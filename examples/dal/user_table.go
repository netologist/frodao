package dal

import (
	"github.com/hasanozgan/frodao"
	"github.com/hasanozgan/frodao/nullable"
)

type UserTable struct {
	frodao.Table

	Username string                `db:"username"`
	Password string                `db:"password"`
	Address  nullable.Type[string] `db:"address"`
}
