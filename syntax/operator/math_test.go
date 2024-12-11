package operator

import (
	"testing"

	"github.com/Rian-wahid/mpgl/syntax/types"
	"github.com/stretchr/testify/assert"
)

func TestMathInt8(t *testing.T) {
	i8_1 := types.NewInt8(1)
	i8_2 := types.NewInt8(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, int8(3), add.Operate(i8_1, i8_2).Value().(int8))
	assert.Equal(t, int8(-1), sub.Operate(i8_1, i8_2).Value().(int8))
	i8_1 = types.NewInt8(8)
	assert.Equal(t, int8(16), mul.Operate(i8_1, i8_2).Value().(int8))
	assert.Equal(t, int8(4), div.Operate(i8_1, i8_2).Value().(int8))
	assert.Equal(t, int8(64), exp.Operate(i8_1, i8_2).Value().(int8))
}

func TestMathByte(t *testing.T) {
	b1 := types.NewByte(3)
	b2 := types.NewByte(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, byte(5), add.Operate(b1, b2).Value().(byte))
	assert.Equal(t, byte(1), sub.Operate(b1, b2).Value().(byte))
	b1 = types.NewByte(8)
	assert.Equal(t, byte(16), mul.Operate(b1, b2).Value().(byte))
	assert.Equal(t, byte(4), div.Operate(b1, b2).Value().(byte))
	assert.Equal(t, byte(64), exp.Operate(b1, b2).Value().(byte))
}

func TestMathInt16(t *testing.T) {
	i16_1 := types.NewInt16(1)
	i16_2 := types.NewInt16(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, int16(3), add.Operate(i16_1, i16_2).Value().(int16))
	assert.Equal(t, int16(-1), sub.Operate(i16_1, i16_2).Value().(int16))
	i16_1 = types.NewInt16(8)
	assert.Equal(t, int16(16), mul.Operate(i16_1, i16_2).Value().(int16))
	assert.Equal(t, int16(4), div.Operate(i16_1, i16_2).Value().(int16))
	assert.Equal(t, int16(64), exp.Operate(i16_1, i16_2).Value().(int16))
}

func TestMathInt32(t *testing.T) {
	a := types.NewInt32(1)
	b := types.NewInt32(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, int32(3), add.Operate(a, b).Value().(int32))
	assert.Equal(t, int32(-1), sub.Operate(a, b).Value().(int32))
	a = types.NewInt32(8)
	assert.Equal(t, int32(16), mul.Operate(a, b).Value().(int32))
	assert.Equal(t, int32(4), div.Operate(a, b).Value().(int32))
	assert.Equal(t, int32(64), exp.Operate(a, b).Value().(int32))
}

func TestMathInt64(t *testing.T) {
	a := types.NewInt64(1)
	b := types.NewInt64(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, int64(3), add.Operate(a, b).Value().(int64))
	assert.Equal(t, int64(-1), sub.Operate(a, b).Value().(int64))
	a = types.NewInt64(8)
	assert.Equal(t, int64(16), mul.Operate(a, b).Value().(int64))
	assert.Equal(t, int64(4), div.Operate(a, b).Value().(int64))
	assert.Equal(t, int64(64), exp.Operate(a, b).Value().(int64))
}

func TestMathUint16(t *testing.T) {
	a := types.NewUint16(3)
	b := types.NewUint16(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, uint16(5), add.Operate(a, b).Value().(uint16))
	assert.Equal(t, uint16(1), sub.Operate(a, b).Value().(uint16))
	a = types.NewUint16(8)
	assert.Equal(t, uint16(16), mul.Operate(a, b).Value().(uint16))
	assert.Equal(t, uint16(4), div.Operate(a, b).Value().(uint16))
	assert.Equal(t, uint16(64), exp.Operate(a, b).Value().(uint16))
}

func TestMathUint32(t *testing.T) {
	a := types.NewUint32(3)
	b := types.NewUint32(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, uint32(5), add.Operate(a, b).Value().(uint32))
	assert.Equal(t, uint32(1), sub.Operate(a, b).Value().(uint32))
	a = types.NewUint32(8)
	assert.Equal(t, uint32(16), mul.Operate(a, b).Value().(uint32))
	assert.Equal(t, uint32(4), div.Operate(a, b).Value().(uint32))
	assert.Equal(t, uint32(64), exp.Operate(a, b).Value().(uint32))
}

func TestMathUint64(t *testing.T) {
	a := types.NewUint64(3)
	b := types.NewUint64(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, uint64(5), add.Operate(a, b).Value().(uint64))
	assert.Equal(t, uint64(1), sub.Operate(a, b).Value().(uint64))
	a = types.NewUint64(8)
	assert.Equal(t, uint64(16), mul.Operate(a, b).Value().(uint64))
	assert.Equal(t, uint64(4), div.Operate(a, b).Value().(uint64))
	assert.Equal(t, uint64(64), exp.Operate(a, b).Value().(uint64))
}

func TestMathFloat32(t *testing.T) {
	a := types.NewFloat32(1)
	b := types.NewFloat32(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, float32(3), add.Operate(a, b).Value().(float32))
	assert.Equal(t, float32(-1), sub.Operate(a, b).Value().(float32))
	a = types.NewFloat32(8)
	assert.Equal(t, float32(16), mul.Operate(a, b).Value().(float32))
	assert.Equal(t, float32(4), div.Operate(a, b).Value().(float32))
	assert.Equal(t, float32(64), exp.Operate(a, b).Value().(float32))
}

func TestMathFloat64(t *testing.T) {
	a := types.NewFloat64(1)
	b := types.NewFloat64(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, float64(3), add.Operate(a, b).Value().(float64))
	assert.Equal(t, float64(-1), sub.Operate(a, b).Value().(float64))
	a = types.NewFloat64(8)
	assert.Equal(t, float64(16), mul.Operate(a, b).Value().(float64))
	assert.Equal(t, float64(4), div.Operate(a, b).Value().(float64))
	assert.Equal(t, float64(64), exp.Operate(a, b).Value().(float64))
}

func TestMathNumber(t *testing.T) {
	a := types.NewNumber(1)
	b := types.NewNumber(2)
	add := &Add{}
	sub := &Sub{}
	mul := &Mul{}
	div := &Div{}
	exp := &Exp{}
	assert.Equal(t, float64(3), add.Operate(a, b).Value().(float64))
	assert.Equal(t, float64(-1), sub.Operate(a, b).Value().(float64))
	a = types.NewNumber(8)
	assert.Equal(t, float64(16), mul.Operate(a, b).Value().(float64))
	assert.Equal(t, float64(4), div.Operate(a, b).Value().(float64))
	assert.Equal(t, float64(64), exp.Operate(a, b).Value().(float64))
}
