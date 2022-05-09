# Frodao

The golang DAO library. Your best friend on the road like Frodo Baggins. 
- Soft delete support.
- Postgresql support.

## Example Schema

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY NOT NULL,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL, 
  address VARCHAR(100) NULL, 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted BOOLEAN NOT NULL DEFAULT FALSE
);
```


## DAO Interface
```go
package frodao

import (
	"context"

	"github.com/hasanozgan/frodao/tableid"
)

type DAO[T Record, I tableid.Constraint] interface {
	Create(ctx context.Context, record *T) (*T, error)
	Update(ctx context.Context, record *T) error
	Delete(ctx context.Context, id ID[I]) error
	FindByID(ctx context.Context, id ID[I]) (*T, error)
	GetTableName() string
}
```

## Custom Table and DAO Definition

```go
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
```

```go
import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/hasanozgan/frodao/drivers/postgres"
	"github.com/hasanozgan/frodao/tableid"
)

func NewUserDAO() *UserDAO {
	return &UserDAO{
		DAO: postgres.DAO[UserTable, tableid.Int]{
			TableName: "users",
		},
	}
}

type UserDAO struct {
	postgres.DAO[UserTable]
}

func (d *UserDAO) FindByUsername(ctx context.Context, username string) ([]*UserTable, error) {
	return d.FindByQuery(ctx, d.SelectQuery().Where(goqu.Ex{"username": username}).Limit(1))
}
```
