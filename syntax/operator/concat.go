package operator

import (
	"fmt"

	"github.com/Rian-wahid/mpgl/syntax/types"
)

type Concat struct{}

func (c *Concat) Operate(operands ...types.Type) types.Type {

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
	case "null":
		f1 = "%s"
		break
	case "boolean":
		f1 = "%t"
		break
	case "int8":
	case "byte":
	case "int16":
	case "int":
	case "int32":
	case "int64":
	case "uint":
	case "uint16":
	case "uint32":
	case "uint64":
		f1 = "%d"
		break
	case "float32":
	case "float64":
	case "number":
		f1 = "%f"
		break
	}

	switch operands[1].TypeName() {
	case "string":
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
	case "byte":
	case "int16":
	case "int":
	case "int32":
	case "int64":
	case "uint":
	case "uint16":
	case "uint32":
	case "uint64":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%d"), val1, val2)
		return types.NewString(str)
	case "float32":
	case "float64":
	case "number":
		val1 := operands[0].Value()
		val2 := operands[1].Value()
		str := fmt.Sprintf(fmt.Sprintf("%s%s", f1, "%f"), val1, val2)
		return types.NewString(str)
	}

	return nil
}
