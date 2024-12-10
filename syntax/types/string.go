package types

type String struct {
	value string
}

func (s *String) TypeName() string {
	return "string"
}

func (s *String) Value() any {
	return s.value
}

func (s *String) Assign(v Type) {
	if v == nil {
		panic("v must not nil")
	}
	if v.TypeName() != "string" {
		panic("invalid type")
	}
	ss := v.Value().(string)
	s.value = ss

}

func (s *String) IsComparable() bool {
	return true
}

func NewString(v string) Type {

	return &String{
		value: v,
	}
}
