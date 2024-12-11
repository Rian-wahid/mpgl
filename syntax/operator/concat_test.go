package operator

import (
	"fmt"
	"testing"

	"github.com/Rian-wahid/mpgl/syntax/types"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	str := types.NewString("this is ")
	concat := &Concat{}
	assert.Equal(t, "this is null", concat.Operate(str, types.NewNull()).Value().(string))
	assert.Equal(t, "nullthis is ", concat.Operate(types.NewNull(), str).Value().(string))
	assert.Equal(t, "this is true", concat.Operate(str, types.NewBoolean(true)).Value().(string))
	assert.Equal(t, "falsethis is ", concat.Operate(types.NewBoolean(false), str).Value().(string))
	assert.Equal(t, "this is 1", concat.Operate(str, types.NewInt8(1)).Value().(string))
	assert.Equal(t, "1this is ", concat.Operate(types.NewInt8(1), str).Value().(string))
	assert.Equal(t, "this is 2", concat.Operate(str, types.NewByte(2)).Value().(string))
	assert.Equal(t, "2this is ", concat.Operate(types.NewByte(2), str).Value().(string))
	assert.Equal(t, "this is 3", concat.Operate(str, types.NewInt16(3)).Value().(string))
	assert.Equal(t, "3this is ", concat.Operate(types.NewInt16(3), str).Value().(string))
	assert.Equal(t, "this is 4", concat.Operate(str, types.NewInt32(4)).Value().(string))
	assert.Equal(t, "4this is ", concat.Operate(types.NewInt32(4), str).Value().(string))
	assert.Equal(t, "this is 5", concat.Operate(str, types.NewInt64(5)).Value().(string))
	assert.Equal(t, "5this is ", concat.Operate(types.NewInt64(5), str).Value().(string))
	assert.Equal(t, "this is 6", concat.Operate(str, types.NewUint16(6)).Value().(string))
	assert.Equal(t, "6this is ", concat.Operate(types.NewUint16(6), str).Value().(string))
	assert.Equal(t, "this is 7", concat.Operate(str, types.NewUint32(7)).Value().(string))
	assert.Equal(t, "7this is ", concat.Operate(types.NewUint32(7), str).Value().(string))
	assert.Equal(t, "this is 8", concat.Operate(str, types.NewUint64(8)).Value().(string))
	assert.Equal(t, "8this is ", concat.Operate(types.NewUint64(8), str).Value().(string))
	assert.Equal(t, fmt.Sprintf("%s%f", str.Value().(string), float32(9.5)), concat.Operate(str, types.NewFloat32(9.5)).Value().(string))
	assert.Equal(t, fmt.Sprintf("%f%s", float32(9.5), str.Value().(string)), concat.Operate(types.NewFloat32(9.5), str).Value().(string))
	assert.Equal(t, fmt.Sprintf("%s%f", str.Value().(string), float64(10.1)), concat.Operate(str, types.NewFloat64(10.1)).Value().(string))
	assert.Equal(t, fmt.Sprintf("%f%s", float64(10.1), str.Value().(string)), concat.Operate(types.NewFloat64(10.1), str).Value().(string))
	assert.Equal(t, fmt.Sprintf("%s%f", str.Value().(string), float64(11.2)), concat.Operate(str, types.NewNumber(11.2)).Value().(string))
	assert.Equal(t, fmt.Sprintf("%f%s", float64(11.2), str.Value().(string)), concat.Operate(types.NewNumber(11.2), str).Value().(string))
}
