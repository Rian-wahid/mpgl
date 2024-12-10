package types

type Null struct{}

func (n *Null) TypeName() string {
	return "null"
}
func (n *Null) Value() any {
	//for string concat & print value
	return "null"
}

func (n *Null) Assign(v Type) {
	_ = v
}

func (n *Null) IsComparable() bool {
	//because for string concat need a comparable type
	return true
}

func NewNull() Type {
	return &Null{}
}
