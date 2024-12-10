package operator

import (
	"math"

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
	if operands[0].TypeName() != operands[0].TypeName() {
		panic("operands type not same")
	}
	switch operands[0].TypeName() {
	case "int8":
		n := typeSliceToNSlice[int8](operands)
		nr := n[0] | n[1]
		return types.NewInt8(nr)
	case "byte":
		n := typeSliceToNSlice[byte](operands)
		nr := n[0] | n[1]
		return types.NewByte(nr)
	case "int16":
		n := typeSliceToNSlice[int16](operands)
		nr := n[0] | n[1]
		return types.NewInt16(nr)
	case "int":
		n := typeSliceToNSlice[int32](operands)
		nr := n[0] | n[1]
		return types.NewInt(nr)
	case "int32":
		n := typeSliceToNSlice[int32](operands)
		nr := n[0] | n[1]
		return types.NewInt32(nr)
	case "int64":
		n := typeSliceToNSlice[int64](operands)
		nr := n[0] | n[1]
		return types.NewInt64(nr)
	case "uint16":
		n := typeSliceToNSlice[uint16](operands)
		nr := n[0] | n[1]
		return types.NewUint16(nr)
	case "uint":
		n := typeSliceToNSlice[uint32](operands)
		nr := n[0] | n[1]
		return types.NewUint(nr)
	case "uint32":
		n := typeSliceToNSlice[uint32](operands)
		nr := n[0] | n[1]
		return types.NewUint32(nr)
	case "uint64":
		n := typeSliceToNSlice[uint64](operands)
		nr := n[0] | n[1]
		return types.NewUint64(nr)
	case "float32":
		n := typeSliceToNSlice[float32](operands)
		nr := math.Float32bits(n[0]) | math.Float32bits(n[1])
		return types.NewFloat32(math.Float32frombits(nr))
	case "float64":
		n := typeSliceToNSlice[float64](operands)
		nr := math.Float64bits(n[0]) | math.Float64bits(n[1])
		return types.NewFloat64(math.Float64frombits(nr))
	case "number":
		n := typeSliceToNSlice[float64](operands)
		nr := int64(n[0]) | int64(n[1])
		return types.NewNumber(float64(nr))
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
	if operands[0].TypeName() != operands[0].TypeName() {
		panic("operands type not same")
	}
	switch operands[0].TypeName() {
	case "int8":
		n := typeSliceToNSlice[int8](operands)
		nr := n[0] & n[1]
		return types.NewInt8(nr)
	case "byte":
		n := typeSliceToNSlice[byte](operands)
		nr := n[0] & n[1]
		return types.NewByte(nr)
	case "int16":
		n := typeSliceToNSlice[int16](operands)
		nr := n[0] & n[1]
		return types.NewInt16(nr)
	case "int":
		n := typeSliceToNSlice[int32](operands)
		nr := n[0] & n[1]
		return types.NewInt(nr)
	case "int32":
		n := typeSliceToNSlice[int32](operands)
		nr := n[0] & n[1]
		return types.NewInt32(nr)
	case "int64":
		n := typeSliceToNSlice[int64](operands)
		nr := n[0] & n[1]
		return types.NewInt64(nr)
	case "uint16":
		n := typeSliceToNSlice[uint16](operands)
		nr := n[0] & n[1]
		return types.NewUint16(nr)
	case "uint":
		n := typeSliceToNSlice[uint32](operands)
		nr := n[0] & n[1]
		return types.NewUint(nr)
	case "uint32":
		n := typeSliceToNSlice[uint32](operands)
		nr := n[0] & n[1]
		return types.NewUint32(nr)
	case "uint64":
		n := typeSliceToNSlice[uint64](operands)
		nr := n[0] & n[1]
		return types.NewUint64(nr)
	case "float32":
		n := typeSliceToNSlice[float32](operands)
		nr := math.Float32bits(n[0]) & math.Float32bits(n[1])
		return types.NewFloat32(math.Float32frombits(nr))
	case "float64":
		n := typeSliceToNSlice[float64](operands)
		nr := math.Float64bits(n[0]) & math.Float64bits(n[1])
		return types.NewFloat64(math.Float64frombits(nr))
	case "number":
		n := typeSliceToNSlice[float64](operands)
		nr := int64(n[0]) & int64(n[1])
		return types.NewNumber(float64(nr))
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
	if operands[0].TypeName() != operands[0].TypeName() {
		panic("operands type not same")
	}
	switch operands[0].TypeName() {
	case "int8":
		n := typeSliceToNSlice[int8](operands)
		nr := n[0] ^ n[1]
		return types.NewInt8(nr)
	case "byte":
		n := typeSliceToNSlice[byte](operands)
		nr := n[0] ^ n[1]
		return types.NewByte(nr)
	case "int16":
		n := typeSliceToNSlice[int16](operands)
		nr := n[0] ^ n[1]
		return types.NewInt16(nr)
	case "int":
		n := typeSliceToNSlice[int32](operands)
		nr := n[0] ^ n[1]
		return types.NewInt(nr)
	case "int32":
		n := typeSliceToNSlice[int32](operands)
		nr := n[0] ^ n[1]
		return types.NewInt32(nr)
	case "int64":
		n := typeSliceToNSlice[int64](operands)
		nr := n[0] ^ n[1]
		return types.NewInt64(nr)
	case "uint16":
		n := typeSliceToNSlice[uint16](operands)
		nr := n[0] ^ n[1]
		return types.NewUint16(nr)
	case "uint":
		n := typeSliceToNSlice[uint32](operands)
		nr := n[0] ^ n[1]
		return types.NewUint(nr)
	case "uint32":
		n := typeSliceToNSlice[uint32](operands)
		nr := n[0] ^ n[1]
		return types.NewUint32(nr)
	case "uint64":
		n := typeSliceToNSlice[uint64](operands)
		nr := n[0] ^ n[1]
		return types.NewUint64(nr)
	case "float32":
		n := typeSliceToNSlice[float32](operands)
		nr := math.Float32bits(n[0]) ^ math.Float32bits(n[1])
		return types.NewFloat32(math.Float32frombits(nr))
	case "float64":
		n := typeSliceToNSlice[float64](operands)
		nr := math.Float64bits(n[0]) ^ math.Float64bits(n[1])
		return types.NewFloat64(math.Float64frombits(nr))
	case "number":
		n := typeSliceToNSlice[float64](operands)
		nr := int64(n[0]) ^ int64(n[1])
		return types.NewNumber(float64(nr))
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
	case "int":
		n := operands[1].Value().(int32)
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
	case "uint":
		n := operands[1].Value().(uint32)
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
	case "int":
		n := operands[0].Value().(int32)
		return types.NewInt(n >> shift)
	case "int32":
		n := operands[0].Value().(int32)
		return types.NewInt32(n >> shift)
	case "int64":
		n := operands[0].Value().(int64)
		return types.NewInt64(n >> shift)
	case "uint":
		n := operands[0].Value().(uint32)
		return types.NewUint(n >> shift)
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
	case "int":
		n := operands[1].Value().(int32)
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
	case "uint":
		n := operands[1].Value().(uint32)
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
	case "int":
		n := operands[0].Value().(int32)
		return types.NewInt(n << shift)
	case "int32":
		n := operands[0].Value().(int32)
		return types.NewInt32(n << shift)
	case "int64":
		n := operands[0].Value().(int64)
		return types.NewInt64(n << shift)
	case "uint":
		n := operands[0].Value().(uint32)
		return types.NewUint(n << shift)
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
