package frodao

import (
	"database/sql/driver"

	"github.com/hasanozgan/frodao/tableid"
)

type ID[T tableid.Constraint] struct {
	IDValue T
	Exist   bool
}

func (n *ID[T]) Scan(value interface{}) error {

	n.IDValue, n.Exist = value.(T)

	return nil
}

func (n ID[T]) Value() (driver.Value, error) {
	if !n.Exist {
		return nil, nil
	}
	return n.IDValue, nil
}

func (n ID[T]) Get() T {
	if !n.Exist {
		panic("value not found")
	}
	return n.IDValue
}

func TableID[T tableid.Constraint](id T) ID[T] {
	return ID[T]{IDValue: id, Exist: true}
}

func TableIDFromInt[T tableid.IntConstraint](value T) ID[tableid.Int] {
	return TableID(tableid.Int(value))
}
