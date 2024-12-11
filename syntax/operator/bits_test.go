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

func TestBitwiseInt64(t *testing.T) {
	i64_1 := types.NewInt64(1)
	i64_2 := types.NewInt64(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	i64_3 := or.Operate(i64_1, i64_2)
	assert.Equal(t, int64(1|2), i64_3.Value().(int64))

	i64_3 = and.Operate(i64_1, i64_2)
	assert.Equal(t, int64(1&2), i64_3.Value().(int64))
	i64_3 = xor.Operate(i64_1, i64_2)
	assert.Equal(t, int64(1^2), i64_3.Value().(int64))
	i64_2 = types.NewInt64(1)
	i64_3 = sl.Operate(i64_1, i64_2)
	assert.Equal(t, int64(1<<1), i64_3.Value().(int64))
	i64_1 = types.NewInt64(2)
	i64_3 = sr.Operate(i64_1, i64_2)
	assert.Equal(t, int64(2>>1), i64_3.Value().(int64))
}

func TestBitwiseUint16(t *testing.T) {
	u16_1 := types.NewUint16(1)
	u16_2 := types.NewUint16(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	assert.Equal(t, uint16(1|2), or.Operate(u16_1, u16_2).Value().(uint16))
	assert.Equal(t, uint16(1&2), and.Operate(u16_1, u16_2).Value().(uint16))
	assert.Equal(t, uint16(1^2), xor.Operate(u16_1, u16_2).Value().(uint16))
	u16_2 = types.NewUint16(1)
	assert.Equal(t, uint16(1<<1), sl.Operate(u16_1, u16_2).Value().(uint16))
	u16_1 = types.NewUint16(2)
	assert.Equal(t, uint16(2>>1), sr.Operate(u16_1, u16_2).Value().(uint16))
}

func TestBitwiseUint32(t *testing.T) {
	u32_1 := types.NewUint32(1)
	u32_2 := types.NewUint32(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	assert.Equal(t, uint32(1|2), or.Operate(u32_1, u32_2).Value().(uint32))
	assert.Equal(t, uint32(1&2), and.Operate(u32_1, u32_2).Value().(uint32))
	assert.Equal(t, uint32(1^2), xor.Operate(u32_1, u32_2).Value().(uint32))
	u32_2 = types.NewUint32(1)
	assert.Equal(t, uint32(1<<1), sl.Operate(u32_1, u32_2).Value().(uint32))
	u32_1 = types.NewUint32(2)
	assert.Equal(t, uint32(2>>1), sr.Operate(u32_1, u32_2).Value().(uint32))
}

func TestBitwiseUint64(t *testing.T) {
	u64_1 := types.NewUint64(1)
	u64_2 := types.NewUint64(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	assert.Equal(t, uint64(1|2), or.Operate(u64_1, u64_2).Value().(uint64))
	assert.Equal(t, uint64(1&2), and.Operate(u64_1, u64_2).Value().(uint64))
	assert.Equal(t, uint64(1^2), xor.Operate(u64_1, u64_2).Value().(uint64))
	u64_2 = types.NewUint64(1)
	assert.Equal(t, uint64(1<<1), sl.Operate(u64_1, u64_2).Value().(uint64))
	u64_1 = types.NewUint64(2)
	assert.Equal(t, uint64(2>>1), sr.Operate(u64_1, u64_2).Value().(uint64))
}

func TestBitwiseFloat32(t *testing.T) {
	f32_1 := types.NewFloat32(1)
	f32_2 := types.NewFloat32(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	assert.Equal(t, float32(1|2), or.Operate(f32_1, f32_2).Value().(float32))
	assert.Equal(t, float32(1&2), and.Operate(f32_1, f32_2).Value().(float32))
	assert.Equal(t, float32(1^2), xor.Operate(f32_1, f32_2).Value().(float32))
	f32_2 = types.NewFloat32(1)
	assert.Equal(t, float32(1<<1), sl.Operate(f32_1, f32_2).Value().(float32))
	f32_1 = types.NewFloat32(2)
	assert.Equal(t, float32(2>>1), sr.Operate(f32_1, f32_2).Value().(float32))
}

func TestBitwiseFloat64(t *testing.T) {
	f64_1 := types.NewFloat64(1)
	f64_2 := types.NewFloat64(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	assert.Equal(t, float64(1|2), or.Operate(f64_1, f64_2).Value().(float64))
	assert.Equal(t, float64(1&2), and.Operate(f64_1, f64_2).Value().(float64))
	assert.Equal(t, float64(1^2), xor.Operate(f64_1, f64_2).Value().(float64))
	f64_2 = types.NewFloat64(1)
	assert.Equal(t, float64(1<<1), sl.Operate(f64_1, f64_2).Value().(float64))
	f64_1 = types.NewFloat64(2)
	assert.Equal(t, float64(2>>1), sr.Operate(f64_1, f64_2).Value().(float64))
}

func TestBitwiseNumber(t *testing.T) {
	num_1 := types.NewNumber(1)
	num_2 := types.NewNumber(2)
	or := &BitOr{}
	and := &BitAnd{}
	xor := &BitXor{}
	sl := &BitLeftShift{}
	sr := &BitRightShift{}
	assert.Equal(t, float64(1|2), or.Operate(num_1, num_2).Value().(float64))
	assert.Equal(t, float64(1&2), and.Operate(num_1, num_2).Value().(float64))
	assert.Equal(t, float64(1^2), xor.Operate(num_1, num_2).Value().(float64))
	num_2 = types.NewNumber(1)
	assert.Equal(t, float64(1<<1), sl.Operate(num_1, num_2).Value().(float64))
	num_1 = types.NewNumber(2)
	assert.Equal(t, float64(2>>1), sr.Operate(num_1, num_2).Value().(float64))
}
