package operator

import "github.com/Rian-wahid/mpgl/syntax/types"

func isNumber(v types.Type) bool {
	switch v.TypeName() {
	case "int":
		return true
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
	case "uint":
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
	return false
}

type number interface {
	int8 | byte | int16 | int32 | int64 | uint16 | uint32 | uint64 | float32 | float64
}

func typeSliceToNSlice[N number](arr []types.Type) []N {
	var result []N
	for i := range arr {
		n := arr[i].Value().(N)
		result = append(result, n)
	}
	return result
}

func add[N number](n []N) N {
	var sum N = 0
	for i := range n {
		sum += n[i]
	}
	return sum
}

func sub[N number](n []N) N {
	var result N = n[0]
	for i := 1; i < len(n); i++ {
		result -= n[i]
	}
	return result
}

func mul[N number](n []N) N {
	var result N = 1
	for i := range n {
		result *= n[i]
	}
	return result
}

func div[N number](n []N) N {
	var result N = n[0]
	for i := 1; i < len(n); i++ {
		result /= n[i]
	}
	return result
}

func exp[N number](n []N) N {
	var result N = n[0]
	for i := 1; i < len(n); i++ {
		var j N = 0
		for ; j < n[i]; i++ {
			result *= result
		}
	}

	return result
}

type Add struct{}

func (a *Add) Operate(operands ...types.Type) types.Type {
	if len(operands) < 2 {
		panic("need min 2 operand")
	}

	if !isNumber(operands[0]) {
		panic("operand must be a number")
	}

	for i := 1; i < len(operands); i++ {
		if operands[i].TypeName() != operands[0].TypeName() {
			panic("operand must be same type")
		}
	}
	switch operands[0].TypeName() {
	case "int8":
		n := add(typeSliceToNSlice[int8](operands))
		return types.NewInt8(n)
	case "byte":
		n := add(typeSliceToNSlice[byte](operands))
		return types.NewByte(n)
	case "int16":
		n := add(typeSliceToNSlice[int16](operands))
		return types.NewInt16(n)
	case "int32":
		n := add(typeSliceToNSlice[int32](operands))
		return types.NewInt32(n)
	case "int":
		n := add(typeSliceToNSlice[int32](operands))
		return types.NewInt(n)
	case "int64":
		n := add(typeSliceToNSlice[int64](operands))
		return types.NewInt64(n)
	case "uint16":
		n := add(typeSliceToNSlice[uint16](operands))
		return types.NewUint16(n)
	case "uint":
		n := add(typeSliceToNSlice[uint32](operands))
		return types.NewUint(n)
	case "uint32":
		n := add(typeSliceToNSlice[uint32](operands))
		return types.NewUint32(n)
	case "uint64":
		n := add(typeSliceToNSlice[uint64](operands))
		return types.NewUint64(n)
	case "float32":
		f := add(typeSliceToNSlice[float32](operands))
		return types.NewFloat32(f)
	case "float64":
		f := add(typeSliceToNSlice[float64](operands))
		return types.NewFloat64(f)
	case "number":
		n := add(typeSliceToNSlice[float64](operands))
		return types.NewNumber(n)

	}

	return nil
}

type Sub struct{}

func (s *Sub) Operate(operands ...types.Type) types.Type {
	if len(operands) < 2 {
		panic("need min 2 operand")
	}

	if !isNumber(operands[0]) {
		panic("operand must be a number")
	}

	for i := 1; i < len(operands); i++ {
		if operands[i].TypeName() != operands[0].TypeName() {
			panic("operand must be same type")
		}
	}
	switch operands[0].TypeName() {
	case "int8":
		n := sub(typeSliceToNSlice[int8](operands))
		return types.NewInt8(n)
	case "byte":
		n := sub(typeSliceToNSlice[byte](operands))
		return types.NewByte(n)
	case "int16":
		n := sub(typeSliceToNSlice[int16](operands))
		return types.NewInt16(n)
	case "int32":
		n := sub(typeSliceToNSlice[int32](operands))
		return types.NewInt32(n)
	case "int":
		n := sub(typeSliceToNSlice[int32](operands))
		return types.NewInt(n)
	case "int64":
		n := sub(typeSliceToNSlice[int64](operands))
		return types.NewInt64(n)
	case "uint16":
		n := sub(typeSliceToNSlice[uint16](operands))
		return types.NewUint16(n)
	case "uint":
		n := sub(typeSliceToNSlice[uint32](operands))
		return types.NewUint(n)
	case "uint32":
		n := sub(typeSliceToNSlice[uint32](operands))
		return types.NewUint32(n)
	case "uint64":
		n := sub(typeSliceToNSlice[uint64](operands))
		return types.NewUint64(n)
	case "float32":
		f := sub(typeSliceToNSlice[float32](operands))
		return types.NewFloat32(f)
	case "float64":
		f := sub(typeSliceToNSlice[float64](operands))
		return types.NewFloat64(f)
	case "number":
		n := sub(typeSliceToNSlice[float64](operands))
		return types.NewNumber(n)

	}

	return nil
}

type Mul struct{}

func (m *Mul) Operate(operands ...types.Type) types.Type {
	if len(operands) < 2 {
		panic("need min 2 operand")
	}

	if !isNumber(operands[0]) {
		panic("operand must be a number")
	}

	for i := 1; i < len(operands); i++ {
		if operands[i].TypeName() != operands[0].TypeName() {
			panic("operand must be same type")
		}
	}
	switch operands[0].TypeName() {
	case "int8":
		n := mul(typeSliceToNSlice[int8](operands))
		return types.NewInt8(n)
	case "byte":
		n := mul(typeSliceToNSlice[byte](operands))
		return types.NewByte(n)
	case "int16":
		n := mul(typeSliceToNSlice[int16](operands))
		return types.NewInt16(n)
	case "int32":
		n := mul(typeSliceToNSlice[int32](operands))
		return types.NewInt32(n)
	case "int":
		n := mul(typeSliceToNSlice[int32](operands))
		return types.NewInt(n)
	case "int64":
		n := mul(typeSliceToNSlice[int64](operands))
		return types.NewInt64(n)
	case "uint16":
		n := mul(typeSliceToNSlice[uint16](operands))
		return types.NewUint16(n)
	case "uint":
		n := mul(typeSliceToNSlice[uint32](operands))
		return types.NewUint(n)
	case "uint32":
		n := mul(typeSliceToNSlice[uint32](operands))
		return types.NewUint32(n)
	case "uint64":
		n := mul(typeSliceToNSlice[uint64](operands))
		return types.NewUint64(n)
	case "float32":
		f := mul(typeSliceToNSlice[float32](operands))
		return types.NewFloat32(f)
	case "float64":
		f := mul(typeSliceToNSlice[float64](operands))
		return types.NewFloat64(f)
	case "number":
		n := mul(typeSliceToNSlice[float64](operands))
		return types.NewNumber(n)

	}

	return nil
}

type Div struct{}

func (d *Div) Operate(operands ...types.Type) types.Type {
	if len(operands) < 2 {
		panic("need min 2 operand")
	}

	if !isNumber(operands[0]) {
		panic("operand must be a number")
	}

	for i := 1; i < len(operands); i++ {
		if operands[i].TypeName() != operands[0].TypeName() {
			panic("operand must be same type")
		}
	}
	switch operands[0].TypeName() {
	case "int8":
		n := div(typeSliceToNSlice[int8](operands))
		return types.NewInt8(n)
	case "byte":
		n := div(typeSliceToNSlice[byte](operands))
		return types.NewByte(n)
	case "int16":
		n := div(typeSliceToNSlice[int16](operands))
		return types.NewInt16(n)
	case "int32":
		n := div(typeSliceToNSlice[int32](operands))
		return types.NewInt32(n)
	case "int":
		n := div(typeSliceToNSlice[int32](operands))
		return types.NewInt(n)
	case "int64":
		n := div(typeSliceToNSlice[int64](operands))
		return types.NewInt64(n)
	case "uint16":
		n := div(typeSliceToNSlice[uint16](operands))
		return types.NewUint16(n)
	case "uint":
		n := div(typeSliceToNSlice[uint32](operands))
		return types.NewUint(n)
	case "uint32":
		n := div(typeSliceToNSlice[uint32](operands))
		return types.NewUint32(n)
	case "uint64":
		n := div(typeSliceToNSlice[uint64](operands))
		return types.NewUint64(n)
	case "float32":
		f := div(typeSliceToNSlice[float32](operands))
		return types.NewFloat32(f)
	case "float64":
		f := div(typeSliceToNSlice[float64](operands))
		return types.NewFloat64(f)
	case "number":
		n := div(typeSliceToNSlice[float64](operands))
		return types.NewNumber(n)

	}

	return nil
}

type Exp struct{}

func (e *Exp) Operate(operands ...types.Type) types.Type {
	if len(operands) < 2 {
		panic("need min 2 operand")
	}

	if !isNumber(operands[0]) {
		panic("operand must be a number")
	}

	for i := 1; i < len(operands); i++ {
		if operands[i].TypeName() != operands[0].TypeName() {
			panic("operand must be same type")
		}
	}
	switch operands[0].TypeName() {
	case "int8":
		n := exp(typeSliceToNSlice[int8](operands))
		return types.NewInt8(n)
	case "byte":
		n := exp(typeSliceToNSlice[byte](operands))
		return types.NewByte(n)
	case "int16":
		n := exp(typeSliceToNSlice[int16](operands))
		return types.NewInt16(n)
	case "int32":
		n := exp(typeSliceToNSlice[int32](operands))
		return types.NewInt32(n)
	case "int":
		n := exp(typeSliceToNSlice[int32](operands))
		return types.NewInt(n)
	case "int64":
		n := exp(typeSliceToNSlice[int64](operands))
		return types.NewInt64(n)
	case "uint16":
		n := exp(typeSliceToNSlice[uint16](operands))
		return types.NewUint16(n)
	case "uint":
		n := exp(typeSliceToNSlice[uint32](operands))
		return types.NewUint(n)
	case "uint32":
		n := exp(typeSliceToNSlice[uint32](operands))
		return types.NewUint32(n)
	case "uint64":
		n := exp(typeSliceToNSlice[uint64](operands))
		return types.NewUint64(n)
	case "float32":
		f := exp(typeSliceToNSlice[float32](operands))
		return types.NewFloat32(f)
	case "float64":
		f := exp(typeSliceToNSlice[float64](operands))
		return types.NewFloat64(f)
	case "number":
		n := exp(typeSliceToNSlice[float64](operands))
		return types.NewNumber(n)

	}

	return nil
}
