package dal

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/netologist/frodao/drivers/postgres"
	"github.com/netologist/frodao/tableid"
)

func NewUserDAO() *UserDAO {
	return &UserDAO{
		DAO: postgres.DAO[UserTable, tableid.Int]{
			TableName: "users",
		},
	}
}

type UserDAO struct {
	postgres.DAO[UserTable, tableid.Int]
}

func (d *UserDAO) FindByUsername(ctx context.Context, username string) (*UserTable, error) {
	return d.FirstRow(
		d.FindByQuery(ctx, d.SelectQuery().Where(goqu.Ex{"username": username}).Limit(1)),
	)
}
