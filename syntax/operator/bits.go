package operator

import (
	"github.com/Rian-wahid/mpgl/syntax/types"
)

type BitOr struct{}

func (b *BitOr) Operate(operands ...types.Type) types.Type {
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
			return types.NewInt8(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}

	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}
	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 | n2))
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 | n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) | n2))
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := int64(operands[0].Value().(float32))
		if operands[1].TypeName() == "float32" {
			n2 := int64(operands[1].Value().(float32))
			return types.NewFloat32(float32(n1 | n2))
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 | n2))
		}
	case "float64":
		n1 := int64(operands[0].Value().(float64))
		if operands[1].TypeName() == "float64" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewFloat64(float64(n1 | n2))
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 | n2))
		}
	case "number":
		n1 := int64(operands[0].Value().(float64))
		var n2 int64
		if operands[1].TypeName() == "int8" {
			n2 = int64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = int64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = int64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = int64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else if operands[1].TypeName() == "uint16" {
			n2 = int64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = int64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = int64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = int64(operands[1].Value().(float32))
		} else {
			n2 = int64(operands[1].Value().(float64))
		}
		return types.NewNumber(float64(n1 | n2))
	}
	return nil
}

type BitAnd struct{}

func (b *BitAnd) Operate(operands ...types.Type) types.Type {
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
			return types.NewInt8(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}

	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}
	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 & n2))
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 & n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) & n2))
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := int64(operands[0].Value().(float32))
		if operands[1].TypeName() == "float32" {
			n2 := int64(operands[1].Value().(float32))
			return types.NewFloat32(float32(n1 & n2))
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 & n2))
		}
	case "float64":
		n1 := int64(operands[0].Value().(float64))
		if operands[1].TypeName() == "float64" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewFloat64(float64(n1 & n2))
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 & n2))
		}
	case "number":
		n1 := int64(operands[0].Value().(float64))
		var n2 int64
		if operands[1].TypeName() == "int8" {
			n2 = int64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = int64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = int64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = int64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else if operands[1].TypeName() == "uint16" {
			n2 = int64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = int64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = int64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = int64(operands[1].Value().(float32))
		} else {
			n2 = int64(operands[1].Value().(float64))
		}
		return types.NewNumber(float64(n1 & n2))
	}
	return nil
}

type BitXor struct{}

func (b *BitXor) Operate(operands ...types.Type) types.Type {
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
			return types.NewInt8(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}
	case "byte":
		n1 := operands[0].Value().(byte)
		if operands[1].TypeName() == "byte" {
			n2 := operands[1].Value().(byte)
			return types.NewByte(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}
	case "int16":
		n1 := operands[0].Value().(int16)
		if operands[1].TypeName() == "int16" {
			n2 := operands[1].Value().(int16)
			return types.NewInt16(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}

	case "int32":
		n1 := operands[0].Value().(int32)
		if operands[1].TypeName() == "int32" {
			n2 := operands[1].Value().(int32)
			return types.NewInt32(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}
	case "int64":
		n1 := operands[0].Value().(int64)
		if operands[1].TypeName() == "int64" {
			n2 := operands[1].Value().(int64)
			return types.NewInt64(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 ^ n2))
		} else {
			panic("invalid type")
		}
	case "uint16":
		n1 := operands[0].Value().(uint16)
		if operands[1].TypeName() == "uint16" {
			n2 := operands[1].Value().(uint16)
			return types.NewUint16(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}

	case "uint32":
		n1 := operands[0].Value().(uint32)
		if operands[1].TypeName() == "uint32" {
			n2 := operands[1].Value().(uint32)
			return types.NewUint32(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}
	case "uint64":
		n1 := operands[0].Value().(uint64)
		if operands[1].TypeName() == "uint64" {
			n2 := operands[1].Value().(uint64)
			return types.NewUint64(n1 ^ n2)
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(int64(n1) ^ n2))
		} else {
			panic("invalid type")
		}
	case "float32":
		n1 := int64(operands[0].Value().(float32))
		if operands[1].TypeName() == "float32" {
			n2 := int64(operands[1].Value().(float32))
			return types.NewFloat32(float32(n1 ^ n2))
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 ^ n2))
		}
	case "float64":
		n1 := int64(operands[0].Value().(float64))
		if operands[1].TypeName() == "float64" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewFloat64(float64(n1 ^ n2))
		} else if operands[1].TypeName() == "number" {
			n2 := int64(operands[1].Value().(float64))
			return types.NewNumber(float64(n1 ^ n2))
		}
	case "number":
		n1 := int64(operands[0].Value().(float64))
		var n2 int64
		if operands[1].TypeName() == "int8" {
			n2 = int64(operands[1].Value().(int8))
		} else if operands[1].TypeName() == "byte" {
			n2 = int64(operands[1].Value().(byte))
		} else if operands[1].TypeName() == "int16" {
			n2 = int64(operands[1].Value().(int16))
		} else if operands[1].TypeName() == "int32" {
			n2 = int64(operands[1].Value().(int32))
		} else if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else if operands[1].TypeName() == "uint16" {
			n2 = int64(operands[1].Value().(uint16))
		} else if operands[1].TypeName() == "uint32" {
			n2 = int64(operands[1].Value().(uint32))
		} else if operands[1].TypeName() == "uint64" {
			n2 = int64(operands[1].Value().(uint64))
		} else if operands[1].TypeName() == "float32" {
			n2 = int64(operands[1].Value().(float32))
		} else {
			n2 = int64(operands[1].Value().(float64))
		}
		return types.NewNumber(float64(n1 ^ n2))
	}
	return nil
}

type BitRightShift struct{}

func (b *BitRightShift) Operate(operands ...types.Type) types.Type {

	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}
	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operands must be number")
	}
	var shift uint64
	switch operands[1].TypeName() {
	case "byte":
		n := operands[1].Value().(byte)
		shift = uint64(n)
	case "int8":
		n := operands[1].Value().(int8)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "int16":
		n := operands[1].Value().(int16)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)

	case "int32":
		n := operands[1].Value().(int32)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "int64":
		n := operands[1].Value().(int64)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "uint16":
		n := operands[1].Value().(uint16)
		shift = uint64(n)

	case "uint32":
		n := operands[1].Value().(uint32)
		shift = uint64(n)
	case "uint64":
		shift = operands[1].Value().(uint64)
	case "float32":
		n := operands[1].Value().(float32)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "float64":
		n := operands[1].Value().(float64)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "number":
		n := operands[1].Value().(float64)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	}

	switch operands[0].TypeName() {
	case "int8":
		n := operands[0].Value().(int8)
		return types.NewInt8(n >> shift)
	case "byte":
		n := operands[0].Value().(byte)
		return types.NewByte(n >> shift)
	case "int16":
		n := operands[0].Value().(int16)
		return types.NewInt16(n >> shift)

	case "int32":
		n := operands[0].Value().(int32)
		return types.NewInt32(n >> shift)
	case "int64":
		n := operands[0].Value().(int64)
		return types.NewInt64(n >> shift)

	case "uint16":
		n := operands[0].Value().(uint16)
		return types.NewUint16(n >> shift)
	case "uint32":
		n := operands[0].Value().(uint32)
		return types.NewUint32(n >> shift)
	case "uint64":
		n := operands[0].Value().(uint64)
		return types.NewUint64(n >> shift)
	case "float32":
		n := operands[0].Value().(float32)
		return types.NewFloat32(float32(int32(n) >> shift))
	case "float64":
		n := operands[0].Value().(float64)
		return types.NewFloat64(float64(int64(n) >> shift))
	case "number":
		n := operands[0].Value().(float64)
		return types.NewFloat64(float64(int64(n) >> shift))

	}
	return nil
}

type BitLeftShift struct{}

func (b *BitLeftShift) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operand")
	}
	if !isNumber(operands[0]) || !isNumber(operands[1]) {
		panic("operands must be number")
	}
	var shift uint64
	switch operands[1].TypeName() {
	case "byte":
		n := operands[1].Value().(byte)
		shift = uint64(n)
	case "int8":
		n := operands[1].Value().(int8)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "int16":
		n := operands[1].Value().(int16)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)

	case "int32":
		n := operands[1].Value().(int32)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "int64":
		n := operands[1].Value().(int64)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "uint16":
		n := operands[1].Value().(uint16)
		shift = uint64(n)

	case "uint32":
		n := operands[1].Value().(uint32)
		shift = uint64(n)
	case "uint64":
		shift = operands[1].Value().(uint64)
	case "float32":
		n := operands[1].Value().(float32)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "float64":
		n := operands[1].Value().(float64)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	case "number":
		n := operands[1].Value().(float64)
		if n < 0 {
			panic("cannot use negatif number")
		}
		shift = uint64(n)
	}

	switch operands[0].TypeName() {
	case "int8":
		n := operands[0].Value().(int8)
		return types.NewInt8(n << shift)
	case "byte":
		n := operands[0].Value().(byte)
		return types.NewByte(n << shift)
	case "int16":
		n := operands[0].Value().(int16)
		return types.NewInt16(n << shift)

	case "int32":
		n := operands[0].Value().(int32)
		return types.NewInt32(n << shift)
	case "int64":
		n := operands[0].Value().(int64)
		return types.NewInt64(n << shift)

	case "uint16":
		n := operands[0].Value().(uint16)
		return types.NewUint16(n << shift)
	case "uint32":
		n := operands[0].Value().(uint32)
		return types.NewUint32(n << shift)
	case "uint64":
		n := operands[0].Value().(uint64)
		return types.NewUint64(n << shift)
	case "float32":
		n := operands[0].Value().(float32)
		return types.NewFloat32(float32(int32(n) << shift))
	case "float64":
		n := operands[0].Value().(float64)
		return types.NewFloat64(float64(int64(n) << shift))
	case "number":
		n := operands[0].Value().(float64)
		return types.NewFloat64(float64(int64(n) << shift))

	}
	return nil
}
