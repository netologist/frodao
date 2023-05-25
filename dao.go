package frodao

import (
	"context"

	"github.com/netologist/frodao/tableid"
)

type DAO[T Record, I tableid.Constraint] interface {
	Create(ctx context.Context, record *T) (*T, error)
	Update(ctx context.Context, record *T) error
	Delete(ctx context.Context, id ID[I]) error
	FindByID(ctx context.Context, id ID[I]) (*T, error)
	GetTableName() string
}
