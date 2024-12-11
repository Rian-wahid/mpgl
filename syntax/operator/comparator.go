package operator

import "github.com/Rian-wahid/mpgl/syntax/types"

type Equal struct{}

func (e *Equal) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot compare")
	}
	switch operands[0].TypeName() {
	case "null":
		return types.NewBoolean(operands[1].TypeName() == "null")
	case "string":
		if operands[1].TypeName() != "string" {
			return types.NewBoolean(false)
		}
		s1 := operands[0].Value().(string)
		s2 := operands[1].Value().(string)
		return types.NewBoolean(s1 == s2)
	case "boolean":
		if operands[1].TypeName() != "boolean" {
			return types.NewBoolean(false)
		}
		b1 := operands[0].Value().(bool)
		b2 := operands[1].Value().(bool)
		return types.NewBoolean(b1 == b2)
	case "int8":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int8" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int8)
		var n2 int8
		if operands[1].TypeName() == "int8" {
			n2 = operands[1].Value().(int8)
		} else {
			n := operands[1].Value().(float64)
			n2 = int8(n)
		}
		return types.NewBoolean(n1 == n2)
	case "byte":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "byte" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(byte)
		var n2 byte
		if operands[1].TypeName() == "byte" {
			n2 = operands[1].Value().(byte)
		} else {
			n := operands[1].Value().(float64)
			n2 = byte(n)
		}
		return types.NewBoolean(n1 == n2)
	case "int16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int16)
		var n2 int16
		if operands[1].TypeName() == "int16" {
			n2 = operands[1].Value().(int16)
		} else {
			n := operands[1].Value().(float64)
			n2 = int16(n)
		}
		return types.NewBoolean(n1 == n2)

	case "int32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int32)
		var n2 int32
		if operands[1].TypeName() == "int32" {
			n2 = operands[1].Value().(int32)
		} else {
			n := operands[1].Value().(float64)
			n2 = int32(n)
		}
		return types.NewBoolean(n1 == n2)
	case "int64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int64)
		var n2 int64
		if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else {
			n := operands[1].Value().(float64)
			n2 = int64(n)
		}
		return types.NewBoolean(n1 == n2)
	case "uint16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint16)
		var n2 uint16
		if operands[1].TypeName() == "uint16" {
			n2 = operands[1].Value().(uint16)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint16(n)
		}
		return types.NewBoolean(n1 == n2)
	case "uint32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint32)
		var n2 uint32
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint32)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint32(n)
		}
		return types.NewBoolean(n1 == n2)
	case "uint64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint64)
		var n2 uint64
		if operands[1].TypeName() == "uint64" {
			n2 = operands[1].Value().(uint64)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint64(n)
		}
		return types.NewBoolean(n1 == n2)
	case "float32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float32)
		var n2 float32
		if operands[1].TypeName() == "float32" {
			n2 = operands[1].Value().(float32)
		} else {
			n := operands[1].Value().(float64)
			n2 = float32(n)
		}
		return types.NewBoolean(n1 == n2)
	case "float64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float64)
		n2 := operands[1].Value().(float64)
		return types.NewBoolean(n1 == n2)
	case "number":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
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
		return types.NewBoolean(n1 == n2)
	}
	return nil
}

type NotEqual struct{}

func (ne *NotEqual) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot compare")
	}
	switch operands[0].TypeName() {
	case "null":
		return types.NewBoolean(operands[1].TypeName() != "null")
	case "string":
		if operands[1].TypeName() != "string" {
			return types.NewBoolean(false)
		}
		s1 := operands[0].Value().(string)
		s2 := operands[1].Value().(string)
		return types.NewBoolean(s1 != s2)
	case "boolean":
		if operands[1].TypeName() != "boolean" {
			return types.NewBoolean(false)
		}
		b1 := operands[0].Value().(bool)
		b2 := operands[1].Value().(bool)
		return types.NewBoolean(b1 != b2)
	case "int8":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int8" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int8)
		var n2 int8
		if operands[1].TypeName() == "int8" {
			n2 = operands[1].Value().(int8)
		} else {
			n := operands[1].Value().(float64)
			n2 = int8(n)
		}
		return types.NewBoolean(n1 != n2)
	case "byte":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "byte" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(byte)
		var n2 byte
		if operands[1].TypeName() == "byte" {
			n2 = operands[1].Value().(byte)
		} else {
			n := operands[1].Value().(float64)
			n2 = byte(n)
		}
		return types.NewBoolean(n1 != n2)
	case "int16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int16)
		var n2 int16
		if operands[1].TypeName() == "int16" {
			n2 = operands[1].Value().(int16)
		} else {
			n := operands[1].Value().(float64)
			n2 = int16(n)
		}
		return types.NewBoolean(n1 != n2)

	case "int32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int32)
		var n2 int32
		if operands[1].TypeName() == "int32" {
			n2 = operands[1].Value().(int32)
		} else {
			n := operands[1].Value().(float64)
			n2 = int32(n)
		}
		return types.NewBoolean(n1 != n2)
	case "int64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int64)
		var n2 int64
		if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else {
			n := operands[1].Value().(float64)
			n2 = int64(n)
		}
		return types.NewBoolean(n1 != n2)
	case "uint16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint16)
		var n2 uint16
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint16)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint16(n)
		}
		return types.NewBoolean(n1 != n2)
	case "uint32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint32)
		var n2 uint32
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint32)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint32(n)
		}
		return types.NewBoolean(n1 != n2)
	case "uint64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint64)
		var n2 uint64
		if operands[1].TypeName() == "uint64" {
			n2 = operands[1].Value().(uint64)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint64(n)
		}
		return types.NewBoolean(n1 != n2)
	case "float32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float32)
		var n2 float32
		if operands[1].TypeName() == "float32" {
			n2 = operands[1].Value().(float32)
		} else {
			n := operands[1].Value().(float64)
			n2 = float32(n)
		}
		return types.NewBoolean(n1 != n2)
	case "float64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float64)
		n2 := operands[1].Value().(float64)
		return types.NewBoolean(n1 != n2)
	case "number":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
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
		return types.NewBoolean(n1 != n2)
	}
	return nil
}

type Less struct{}

func (l *Less) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot compare")
	}
	switch operands[0].TypeName() {
	case "null":
		panic("null cannot compared with <")
	case "string":
		panic("string cannot compared with <")
	case "boolean":
		panic("boolean cannot compared with <")
	case "int8":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int8" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int8)
		var n2 int8
		if operands[1].TypeName() == "int8" {
			n2 = operands[1].Value().(int8)
		} else {
			n := operands[1].Value().(float64)
			n2 = int8(n)
		}
		return types.NewBoolean(n1 < n2)
	case "byte":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "byte" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(byte)
		var n2 byte
		if operands[1].TypeName() == "byte" {
			n2 = operands[1].Value().(byte)
		} else {
			n := operands[1].Value().(float64)
			n2 = byte(n)
		}
		return types.NewBoolean(n1 < n2)
	case "int16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int16)
		var n2 int16
		if operands[1].TypeName() == "int16" {
			n2 = operands[1].Value().(int16)
		} else {
			n := operands[1].Value().(float64)
			n2 = int16(n)
		}
		return types.NewBoolean(n1 < n2)

	case "int32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int32)
		var n2 int32
		if operands[1].TypeName() == "int32" {
			n2 = operands[1].Value().(int32)
		} else {
			n := operands[1].Value().(float64)
			n2 = int32(n)
		}
		return types.NewBoolean(n1 < n2)
	case "int64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int64)
		var n2 int64
		if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else {
			n := operands[1].Value().(float64)
			n2 = int64(n)
		}
		return types.NewBoolean(n1 < n2)
	case "uint16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint16)
		var n2 uint16
		if operands[1].TypeName() == "uint16" {
			n2 = operands[1].Value().(uint16)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint16(n)
		}
		return types.NewBoolean(n1 < n2)
	case "uint32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint32)
		var n2 uint32
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint32)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint32(n)
		}
		return types.NewBoolean(n1 < n2)
	case "uint64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint64)
		var n2 uint64
		if operands[1].TypeName() == "uint64" {
			n2 = operands[1].Value().(uint64)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint64(n)
		}
		return types.NewBoolean(n1 < n2)
	case "float32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float32)
		var n2 float32
		if operands[1].TypeName() == "float32" {
			n2 = operands[1].Value().(float32)
		} else {
			n := operands[1].Value().(float64)
			n2 = float32(n)
		}
		return types.NewBoolean(n1 < n2)
	case "float64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float64)
		n2 := operands[1].Value().(float64)
		return types.NewBoolean(n1 < n2)
	case "number":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
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
		return types.NewBoolean(n1 < n2)
	}

	return nil
}

type More struct{}

func (m *More) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot compare")
	}
	switch operands[0].TypeName() {
	case "null":
		panic("null cannot compared with >")
	case "string":
		panic("string cannot compared with >")
	case "boolean":
		panic("boolean cannot compared with >")
	case "int8":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int8" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int8)
		var n2 int8
		if operands[1].TypeName() == "int8" {
			n2 = operands[1].Value().(int8)
		} else {
			n := operands[1].Value().(float64)
			n2 = int8(n)
		}
		return types.NewBoolean(n1 > n2)
	case "byte":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "byte" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(byte)
		var n2 byte
		if operands[1].TypeName() == "byte" {
			n2 = operands[1].Value().(byte)
		} else {
			n := operands[1].Value().(float64)
			n2 = byte(n)
		}
		return types.NewBoolean(n1 > n2)
	case "int16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int16)
		var n2 int16
		if operands[1].TypeName() == "int16" {
			n2 = operands[1].Value().(int16)
		} else {
			n := operands[1].Value().(float64)
			n2 = int16(n)
		}
		return types.NewBoolean(n1 > n2)

	case "int32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int32)
		var n2 int32
		if operands[1].TypeName() == "int32" {
			n2 = operands[1].Value().(int32)
		} else {
			n := operands[1].Value().(float64)
			n2 = int32(n)
		}
		return types.NewBoolean(n1 > n2)
	case "int64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int64)
		var n2 int64
		if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else {
			n := operands[1].Value().(float64)
			n2 = int64(n)
		}
		return types.NewBoolean(n1 > n2)
	case "uint16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint16)
		var n2 uint16
		if operands[1].TypeName() == "uint16" {
			n2 = operands[1].Value().(uint16)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint16(n)
		}
		return types.NewBoolean(n1 > n2)
	case "uint32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint32)
		var n2 uint32
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint32)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint32(n)
		}
		return types.NewBoolean(n1 > n2)
	case "uint64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint64)
		var n2 uint64
		if operands[1].TypeName() == "uint64" {
			n2 = operands[1].Value().(uint64)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint64(n)
		}
		return types.NewBoolean(n1 > n2)
	case "float32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float32)
		var n2 float32
		if operands[1].TypeName() == "float32" {
			n2 = operands[1].Value().(float32)
		} else {
			n := operands[1].Value().(float64)
			n2 = float32(n)
		}
		return types.NewBoolean(n1 > n2)
	case "float64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float64)
		n2 := operands[1].Value().(float64)
		return types.NewBoolean(n1 > n2)
	case "number":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
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
		return types.NewBoolean(n1 > n2)
	}
	return nil
}

type LessOrEqual struct{}

func (le *LessOrEqual) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot compare")
	}
	switch operands[0].TypeName() {
	case "null":
		panic("null cannot compared with <=")
	case "string":
		panic("string cannot compared with <=")
	case "boolean":
		panic("boolean cannot compared with <=")
	case "int8":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int8" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int8)
		var n2 int8
		if operands[1].TypeName() == "int8" {
			n2 = operands[1].Value().(int8)
		} else {
			n := operands[1].Value().(float64)
			n2 = int8(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "byte":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "byte" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(byte)
		var n2 byte
		if operands[1].TypeName() == "byte" {
			n2 = operands[1].Value().(byte)
		} else {
			n := operands[1].Value().(float64)
			n2 = byte(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "int16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int16)
		var n2 int16
		if operands[1].TypeName() == "int16" {
			n2 = operands[1].Value().(int16)
		} else {
			n := operands[1].Value().(float64)
			n2 = int16(n)
		}
		return types.NewBoolean(n1 <= n2)

	case "int32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int32)
		var n2 int32
		if operands[1].TypeName() == "int32" {
			n2 = operands[1].Value().(int32)
		} else {
			n := operands[1].Value().(float64)
			n2 = int32(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "int64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int64)
		var n2 int64
		if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else {
			n := operands[1].Value().(float64)
			n2 = int64(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "uint16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint16)
		var n2 uint16
		if operands[1].TypeName() == "uint" {
			n2 = operands[1].Value().(uint16)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint16(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "uint32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint32)
		var n2 uint32
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint32)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint32(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "uint64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint64)
		var n2 uint64
		if operands[1].TypeName() == "uint64" {
			n2 = operands[1].Value().(uint64)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint64(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "float32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float32)
		var n2 float32
		if operands[1].TypeName() == "float32" {
			n2 = operands[1].Value().(float32)
		} else {
			n := operands[1].Value().(float64)
			n2 = float32(n)
		}
		return types.NewBoolean(n1 <= n2)
	case "float64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float64)
		n2 := operands[1].Value().(float64)
		return types.NewBoolean(n1 <= n2)
	case "number":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
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
		return types.NewBoolean(n1 <= n2)
	}
	return nil
}

type MoreOrEqual struct{}

func (me *MoreOrEqual) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}
	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")
	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot compare")
	}
	switch operands[0].TypeName() {
	case "null":
		panic("null cannot compared with >=")
	case "string":
		panic("string cannot compared with >=")
	case "boolean":
		panic("boolean cannot compared with >=")
	case "int8":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int8" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int8)
		var n2 int8
		if operands[1].TypeName() == "int8" {
			n2 = operands[1].Value().(int8)
		} else {
			n := operands[1].Value().(float64)
			n2 = int8(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "byte":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "byte" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(byte)
		var n2 byte
		if operands[1].TypeName() == "byte" {
			n2 = operands[1].Value().(byte)
		} else {
			n := operands[1].Value().(float64)
			n2 = byte(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "int16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int16)
		var n2 int16
		if operands[1].TypeName() == "int16" {
			n2 = operands[1].Value().(int16)
		} else {
			n := operands[1].Value().(float64)
			n2 = int16(n)
		}
		return types.NewBoolean(n1 >= n2)

	case "int32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int32)
		var n2 int32
		if operands[1].TypeName() == "int32" {
			n2 = operands[1].Value().(int32)
		} else {
			n := operands[1].Value().(float64)
			n2 = int32(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "int64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "int64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(int64)
		var n2 int64
		if operands[1].TypeName() == "int64" {
			n2 = operands[1].Value().(int64)
		} else {
			n := operands[1].Value().(float64)
			n2 = int64(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "uint16":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint16" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint16)
		var n2 uint16
		if operands[1].TypeName() == "uint16" {
			n2 = operands[1].Value().(uint16)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint16(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "uint32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint32)
		var n2 uint32
		if operands[1].TypeName() == "uint32" {
			n2 = operands[1].Value().(uint32)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint32(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "uint64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "uint64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(uint64)
		var n2 uint64
		if operands[1].TypeName() == "uint64" {
			n2 = operands[1].Value().(uint64)
		} else {
			n := operands[1].Value().(float64)
			n2 = uint64(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "float32":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float32" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float32)
		var n2 float32
		if operands[1].TypeName() == "float32" {
			n2 = operands[1].Value().(float32)
		} else {
			n := operands[1].Value().(float64)
			n2 = float32(n)
		}
		return types.NewBoolean(n1 >= n2)
	case "float64":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
		if operands[1].TypeName() != "float64" && operands[1].TypeName() != "number" {
			return types.NewBoolean(false)
		}
		n1 := operands[0].Value().(float64)
		n2 := operands[1].Value().(float64)
		return types.NewBoolean(n1 >= n2)
	case "number":
		if !isNumber(operands[1]) {
			return types.NewBoolean(false)
		}
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
		return types.NewBoolean(n1 >= n2)
	}
	return nil
}
