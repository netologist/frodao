package frodao

import (
	"time"

	"github.com/netologist/frodao/tableid"
)

type Record interface {
}

type Table[T tableid.Constraint] struct {
	Record `db:"-" goqu:"skipinsert,skipupdate"`

	ID        ID[T]     `db:"id" goqu:"skipinsert,skipupdate"`
	CreatedAt time.Time `db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt time.Time `db:"updated_at" goqu:"skipinsert"`
	Deleted   bool      `db:"deleted" goqu:"skipinsert,skipupdate"`
}

func (t *Table[T]) SetID(id ID[T]) {
	t.ID = id
}

func (t *Table[T]) SetCreatedAt(createdAt time.Time) {
	t.CreatedAt = createdAt
}

func (t *Table[T]) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t *Table[T]) SetDeleted(deleted bool) {
	t.Deleted = deleted
}
