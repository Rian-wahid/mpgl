package types

type Boolean struct {
	value bool
}

func (b *Boolean) Value() any {

	return b.value
}

func (b *Boolean) TypeName() string {
	return "boolean"
}

func (b *Boolean) Assign(v Type) {
	if v.TypeName() != "boolean" {
		panic("invalid type")
	}
	bl := v.Value().(bool)
	b.value = bl
}

func (b *Boolean) IsComparable() bool {
	return true
}

func NewBoolean(b bool) *Boolean {

	return &Boolean{b}
}
