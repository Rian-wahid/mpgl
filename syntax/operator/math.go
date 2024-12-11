package operator

import (
	"math"

	"github.com/Rian-wahid/mpgl/syntax/types"
)

func isNumber(v types.Type) bool {
	switch v.TypeName() {
	case "byte":
		return true
	case "int8":
		return true
	case "int16":
		return true
	case "int32":
		return true
	case "int64":
		return true
	case "uint16":
		return true
	case "uint32":
		return true
	case "uint64":
		return true
	case "float32":
		return true
	case "float64":
		return true
	case "number":
		return true
	default:
		return false
	}

}

type Add struct{}

func (a *Add) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}

	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operand must be a number")
	}

	switch operands[0].TypeName() {
	case "int8":
		n1 := operands[0].Value().(int8)
		if operands[1].TypeName() == "int8" {
			n2 := operands[1].Value().(int8)
			return types.NewInt8(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}
	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}

	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := operands[0].Value().(float32)
		if operands[1].TypeName() == "float32" {
			n2 := operands[1].Value().(float32)
			return types.NewFloat32(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) + n2)
		} else {
			panic("invalid type")
		}

	case "float64":
		n1 := operands[0].Value().(float64)
		if operands[1].TypeName() == "float64" {
			n2 := operands[1].Value().(float64)
			return types.NewFloat64(n1 + n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(n1 + n2)
		} else {
			panic("invalid type")
		}
	case "number":
		n1 := operands[0].Value().(float64)
		var n2 float64
		if operands[1].TypeName() == "int8" {
			n2 = float64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = float64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = float64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = float64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = float64(operands[1].Value().(int64))
		} else if operands[1].TypeName() == "uint16" {
			n2 = float64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = float64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = float64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = float64(operands[1].Value().(float32))
		} else {
			n2 = operands[1].Value().(float64)
		}
		return types.NewNumber(n1 + n2)

	}

	return nil
}

type Sub struct{}

func (s *Sub) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}

	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}

	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operand must be a number")
	}

	switch operands[0].TypeName() {
	case "int8":
		n1 := operands[0].Value().(int8)
		if operands[1].TypeName() == "int8" {
			n2 := operands[1].Value().(int8)
			return types.NewInt8(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}
	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}

	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := operands[0].Value().(float32)
		if operands[1].TypeName() == "float32" {
			n2 := operands[1].Value().(float32)
			return types.NewFloat32(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) - n2)
		} else {
			panic("invalid type")
		}

	case "float64":
		n1 := operands[0].Value().(float64)
		if operands[1].TypeName() == "float64" {
			n2 := operands[1].Value().(float64)
			return types.NewFloat64(n1 - n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(n1 - n2)
		} else {
			panic("invalid type")
		}
	case "number":
		n1 := operands[0].Value().(float64)
		var n2 float64
		if operands[1].TypeName() == "int8" {
			n2 = float64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = float64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = float64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = float64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = float64(operands[1].Value().(int64))
		} else if operands[1].TypeName() == "uint16" {
			n2 = float64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = float64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = float64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = float64(operands[1].Value().(float32))
		} else {
			n2 = operands[1].Value().(float64)
		}
		return types.NewNumber(n1 - n2)
	}

	return nil
}

type Mul struct{}

func (m *Mul) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}

	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}

	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operand must be a number")
	}
	switch operands[0].TypeName() {
	case "int8":
		n1 := operands[0].Value().(int8)
		if operands[1].TypeName() == "int8" {
			n2 := operands[1].Value().(int8)
			return types.NewInt8(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}
	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}

	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := operands[0].Value().(float32)
		if operands[1].TypeName() == "float32" {
			n2 := operands[1].Value().(float32)
			return types.NewFloat32(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) * n2)
		} else {
			panic("invalid type")
		}

	case "float64":
		n1 := operands[0].Value().(float64)
		if operands[1].TypeName() == "float64" {
			n2 := operands[1].Value().(float64)
			return types.NewFloat64(n1 * n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(n1 * n2)
		} else {
			panic("invalid type")
		}
	case "number":
		n1 := operands[0].Value().(float64)
		var n2 float64
		if operands[1].TypeName() == "int8" {
			n2 = float64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = float64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = float64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = float64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = float64(operands[1].Value().(int64))
		} else if operands[1].TypeName() == "uint16" {
			n2 = float64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = float64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = float64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = float64(operands[1].Value().(float32))
		} else {
			n2 = operands[1].Value().(float64)
		}
		return types.NewNumber(n1 * n2)

	}

	return nil
}

type Div struct{}

func (d *Div) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}

	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}

	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operand must be a number")
	}
	switch operands[0].TypeName() {
	case "int8":
		n1 := operands[0].Value().(int8)
		if operands[1].TypeName() == "int8" {
			n2 := operands[1].Value().(int8)
			return types.NewInt8(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}
	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}

	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := operands[0].Value().(float32)
		if operands[1].TypeName() == "float32" {
			n2 := operands[1].Value().(float32)
			return types.NewFloat32(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(float64(n1) / n2)
		} else {
			panic("invalid type")
		}

	case "float64":
		n1 := operands[0].Value().(float64)
		if operands[1].TypeName() == "float64" {
			n2 := operands[1].Value().(float64)
			return types.NewFloat64(n1 / n2)
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(n1 / n2)
		} else {
			panic("invalid type")
		}
	case "number":
		n1 := operands[0].Value().(float64)
		var n2 float64
		if operands[1].TypeName() == "int8" {
			n2 = float64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = float64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = float64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = float64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = float64(operands[1].Value().(int64))
		} else if operands[1].TypeName() == "uint16" {
			n2 = float64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = float64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = float64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = float64(operands[1].Value().(float32))
		} else {
			n2 = operands[1].Value().(float64)
		}
		return types.NewNumber(n1 / n2)

	}

	return nil
}

type Exp struct{}

func (e *Exp) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}

	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}

	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operand must be a number")
	}
	switch operands[0].TypeName() {
	case "int8":
		n1 := operands[0].Value().(int8)
		if operands[1].TypeName() == "int8" {
			n2 := operands[1].Value().(int8)
			return types.NewInt8(int8(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(byte(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(int16(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}
	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(int32(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}

	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(int64(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(uint16(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(uint32(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(uint64(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := operands[0].Value().(float32)
		if operands[1].TypeName() == "float32" {
			n2 := operands[1].Value().(float32)
			return types.NewFloat32(float32(math.Pow(float64(n1), float64(n2))))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(float64(n1), n2))
		} else {
			panic("invalid type")
		}

	case "float64":
		n1 := operands[0].Value().(float64)
		if operands[1].TypeName() == "float64" {
			n2 := operands[1].Value().(float64)
			return types.NewFloat64(math.Pow(n1, n2))
		} else if operands[1].TypeName() == "number" {
			n2 := operands[1].Value().(float64)
			return types.NewNumber(math.Pow(n1, n2))
		} else {
			panic("invalid type")
		}
	case "number":
		n1 := operands[0].Value().(float64)
		var n2 float64
		if operands[1].TypeName() == "int8" {
			n2 = float64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = float64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = float64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = float64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = float64(operands[1].Value().(int64))
		} else if operands[1].TypeName() == "uint16" {
			n2 = float64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = float64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = float64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = float64(operands[1].Value().(float32))
		} else {
			n2 = operands[1].Value().(float64)
		}
		return types.NewNumber(math.Pow(n1, n2))
	}

	return nil
}
