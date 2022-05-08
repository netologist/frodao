package tableid

type Int = int64
type String = string
type UUID = [16]byte

type Constraint interface {
	~Int | ~String | ~UUID
}

type IntConstraint interface {
	~int | ~int32 | ~int64
}
