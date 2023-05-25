package dal

import (
	"github.com/netologist/frodao"
	"github.com/netologist/frodao/nullable"
	"github.com/netologist/frodao/tableid"
)

type UserTable struct {
	frodao.Table[tableid.Int]

	Username string                `db:"username"`
	Password string                `db:"password"`
	Address  nullable.Type[string] `db:"address"`
}
