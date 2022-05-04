package dal

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasanozgan/frodao/drivers/postgres"
)

func NewUserDAO() *UserDAO {
	return &UserDAO{
		DAO: postgres.DAO[UserTable]{
			TableName: "users",
		},
	}
}

type UserDAO struct {
	postgres.DAO[UserTable]
}

func (d *UserDAO) FindByUsername(ctx context.Context, username string) (*UserTable, error) {
	return d.FindByQuery(ctx, d.SelectQuery().Where(goqu.Ex{"username": username}).Limit(1))
}
