package operator

import (
	"fmt"

	"github.com/Rian-wahid/mpgl/syntax/types"
)

type Concat struct{}

func (c *Concat) Operate(operands ...types.Type) types.Type {
	if operands == nil {
		panic("operands must not nil")
	}

	if len(operands) < 2 || len(operands) > 2 {
		panic("need 2 operands")

	}
	if !operands[0].IsComparable() || !operands[1].IsComparable() {
		panic("cannot concat non-comparable type")
	}
	if operands[0].TypeName() != "string" && operands[1].TypeName() != "string" {
		panic("1 operand must be a string")
	}

	var f1 string
	switch operands[0].TypeName() {
	case "string":
		f1 = "%s"
		break
	case "null":
		f1 = "%s"
		break
	case "boolean":
		f1 = "%t"
		break
	case "int8":
		f1 = "%d"
		break
	case "byte":
		f1 = "%d"
		break
	case "int16":
		f1 = "%d"
		break
	case "int32":
		f1 = "%d"
		break
	case "int64":
		f1 = "%d"
		break
	case "uint16":
		f1 = "%d"
		break
	case "uint32":
		f1 = "%d"
		break
	case "uint64":
		f1 = "%d"
		break
	case "float32":
		f1 = "%f"
		break
	case "float64":
		f1 = "%f"
		break
	case "number":
		f1 = "%f"
		break
	}

	switch operands[1].TypeName() {
	case "string":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%s"), val1, val2)
		return types.NewString(str)
	case "null":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%s"), val1, val2)
		return types.NewString(str)
	case "boolean":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%t"), val1, val2)
		return types.NewString(str)
	case "int8":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "byte":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "int16":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "int32":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "int64":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "uint16":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "uint32":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "uint64":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "float32":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%f"), val1, val2)
		return types.NewString(str)
	case "float64":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%f"), val1, val2)
		return types.NewString(str)
	case "number":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%f"), val1, val2)
		return types.NewString(str)
	}

	return nil
}
