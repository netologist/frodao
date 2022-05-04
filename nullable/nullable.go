package nullable

import "database/sql/driver"

type Type[T any] struct {
	TypeValue T
	Exist     bool
}

func (n *Type[T]) Scan(value interface{}) error {
	n.TypeValue, n.Exist = value.(T)
	return nil
}

func (n Type[T]) Value() (driver.Value, error) {
	if !n.Exist {
		return nil, nil
	}
	return n.TypeValue, nil
}

func New[T any](value T) Type[T] {
	return Type[T]{
		TypeValue: value,
		Exist:     true,
	}
}
