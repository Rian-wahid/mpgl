package operator

import "github.com/Rian-wahid/mpgl/syntax/types"

type Operator interface {
	Operate(operands ...types.Type) types.Type
}
