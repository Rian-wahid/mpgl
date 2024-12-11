package operator

import (
	"testing"

	"github.com/Rian-wahid/mpgl/syntax/types"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseInt8(t *testing.T) {
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	i8_1 := types.NewInt8(1)
	i8_2 := types.NewInt8(2)
	i8_3 := or.Operate(i8_1, i8_2)

	assert.Equal(t, int8(1|2), i8_3.Value().(int8))

	i8_3 = and.Operate(i8_1, i8_2)
	assert.Equal(t, int8(1&2), i8_3.Value().(int8))

	i8_3 = xor.Operate(i8_1, i8_2)
	assert.Equal(t, int8(1^2), i8_3.Value().(int8))

	i8_2 = types.NewInt8(1)
	i8_3 = sl.Operate(i8_1, i8_2)
	assert.Equal(t, int8(1<<1), i8_3.Value().(int8))
	i8_1 = types.NewInt8(2)
	i8_3 = sr.Operate(i8_1, i8_2)
	assert.Equal(t, int8(2>>1), i8_3.Value().(int8))

}

func TestBitwiseByte(t *testing.T) {
	b1 := types.NewByte(1)
	b2 := types.NewByte(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	b3 := or.Operate(b1, b2)
	assert.Equal(t, byte(1|2), b3.Value().(byte))

	b3 = and.Operate(b1, b2)
	assert.Equal(t, byte(1&2), b3.Value().(byte))

	b3 = xor.Operate(b1, b2)
	assert.Equal(t, byte(1^2), b3.Value().(byte))

	b2 = types.NewByte(1)
	b3 = sl.Operate(b1, b2)
	assert.Equal(t, byte(1<<1), b3.Value().(byte))

	b1 = types.NewByte(2)
	b3 = sr.Operate(b1, b2)
	assert.Equal(t, byte(2>>1), b3.Value().(byte))

}

func TestBitwiseInt16(t *testing.T) {
	i16_1 := types.NewInt16(1)
	i16_2 := types.NewInt16(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	i16_3 := or.Operate(i16_1, i16_2)
	assert.Equal(t, int16(1|2), i16_3.Value().(int16))

	i16_3 = and.Operate(i16_1, i16_2)
	assert.Equal(t, int16(1&2), i16_3.Value().(int16))
	i16_3 = xor.Operate(i16_1, i16_2)
	assert.Equal(t, int16(1^2), i16_3.Value().(int16))
	i16_2 = types.NewInt16(1)
	i16_3 = sl.Operate(i16_1, i16_2)
	assert.Equal(t, int16(1<<1), i16_3.Value().(int16))
	i16_1 = types.NewInt16(2)
	i16_3 = sr.Operate(i16_1, i16_2)
	assert.Equal(t, int16(2>>1), i16_3.Value().(int16))
}

func TestBitwiseInt32(t *testing.T) {
	i32_1 := types.NewInt32(1)
	i32_2 := types.NewInt32(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	i32_3 := or.Operate(i32_1, i32_2)
	assert.Equal(t, int32(1|2), i32_3.Value().(int32))

	i32_3 = and.Operate(i32_1, i32_2)
	assert.Equal(t, int32(1&2), i32_3.Value().(int32))
	i32_3 = xor.Operate(i32_1, i32_2)
	assert.Equal(t, int32(1^2), i32_3.Value().(int32))
	i32_2 = types.NewInt32(1)
	i32_3 = sl.Operate(i32_1, i32_2)
	assert.Equal(t, int32(1<<1), i32_3.Value().(int32))
	i32_1 = types.NewInt32(2)
	i32_3 = sr.Operate(i32_1, i32_2)
	assert.Equal(t, int32(2>>1), i32_3.Value().(int32))
}
