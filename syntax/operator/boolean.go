package operator

import "github.com/Rian-wahid/mpgl/syntax/types"

type Not struct{}

func (n *Not) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 1 || len(operands) > 1 {
		panic("need 1 operand")
	}
	if operands[0].TypeName() != "boolean" {
		panic("invalid type")
	}
	b := operands[0].Value().(bool)

	return types.NewBoolean(!b)
}

type Or struct{}

func (o *Or) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if operands[0].TypeName() != "boolean" || operands[1].TypeName() != "boolean" {
		panic("invalid type")
	}
	b1 := operands[0].Value().(bool)
	b2 := operands[1].Value().(bool)

	return types.NewBoolean(b1 || b2)
}

type And struct{}

func (a *And) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if operands[0].TypeName() != "boolean" || operands[1].TypeName() != "boolean" {
		panic("invalid type")
	}
	b1 := operands[0].Value().(bool)
	b2 := operands[1].Value().(bool)
	return types.NewBoolean(b1 && b2)
}
