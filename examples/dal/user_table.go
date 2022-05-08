package dal

import (
	"github.com/hasanozgan/frodao"
	"github.com/hasanozgan/frodao/nullable"
	"github.com/hasanozgan/frodao/tableid"
)

type UserTable struct {
	frodao.Table[tableid.Int]

	Username string                `db:"username"`
	Password string                `db:"password"`
	Address  nullable.Type[string] `db:"address"`
}
