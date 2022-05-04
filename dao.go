package frodao

import (
	"context"
)

type DAO[T Record] interface {
	Create(ctx context.Context, record *T) (*T, error)
	Update(ctx context.Context, record *T) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*T, error)
	GetTableName() string
}
