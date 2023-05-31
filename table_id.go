package frodao

import (
	"database/sql/driver"
	"reflect"
	"strconv"

	"github.com/netologist/frodao/tableid"
)

type ID[T tableid.Constraint] struct {
	IDValue T
	Exist   bool
}

func (n *ID[T]) Scan(value interface{}) error {
	typeOfID := reflect.TypeOf(n.IDValue).String()
	typeOfValue := reflect.TypeOf(value).String()

	if typeOfID == "int64" && typeOfValue == "int32" {
		n.convertToInt64(int64(value.(int32)))
	} else if typeOfID == "int64" && typeOfValue == "int16" {
		n.convertToInt64(int64(value.(int16)))
	} else if typeOfID == "int64" && typeOfValue == "int" {
		n.convertToInt64(int64(value.(int)))
	} else if typeOfID == "int64" && typeOfValue == "[]uint8" {
		intValue, err := strconv.Atoi(string(value.([]byte)))
		n.convertToInt64(int64(intValue))
		n.Exist = (err == nil)
	} else if typeOfID == "string" && typeOfValue == "[]uint8" {
		rv := reflect.ValueOf(string(value.([]byte)))
		n.IDValue, n.Exist = rv.Interface().(T)
	} else {
		n.IDValue, n.Exist = value.(T)
	}

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

func (n *ID[T]) convertToInt64(i64 int64) {
	rv := reflect.ValueOf(i64)
	n.IDValue, n.Exist = rv.Interface().(T)
}

func TableID[T tableid.Constraint](id T) ID[T] {
	return ID[T]{IDValue: id, Exist: true}
}

func TableIDFromInt[T tableid.IntConstraint](value T) ID[tableid.Int] {
	return TableID(tableid.Int(value))
}
