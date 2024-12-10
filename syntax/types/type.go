package types

type Type interface {
	Value() any
	TypeName() string
	Assign(v Type)
	IsComparable() bool
}

type ArrayType interface {
	Value() []Type
	TypeName() string
	Assign(v []Type)
	Set(i uint64, v Type)
	Get(i uint64) Type
	Length() uint64
	IsComparable() bool
}
