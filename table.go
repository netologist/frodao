package frodao

import (
	"time"
)

type Record interface {
}

type Table struct {
	Record `db:"-" goqu:"skipinsert,skipupdate"`

	ID        ID        `db:"id" goqu:"skipinsert,skipupdate"`
	CreatedAt time.Time `db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt time.Time `db:"updated_at" goqu:"skipinsert"`
	Deleted   bool      `db:"deleted" goqu:"skipinsert,skipupdate"`
}

func (t *Table) SetID(id ID) {
	t.ID = id
}

func (t *Table) SetCreatedAt(createdAt time.Time) {
	t.CreatedAt = createdAt
}

func (t *Table) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t *Table) SetDeleted(deleted bool) {
	t.Deleted = deleted
}
