package operator

import (
	"testing"

	"github.com/Rian-wahid/mpgl/syntax/types"
	"github.com/stretchr/testify/assert"
)

func TestBooleanOperator(t *testing.T) {
	not := &Not{}
	and := &And{}
	or := &Or{}
	assert.Equal(t, true, not.Operate(types.NewBoolean(false)).Value().(bool))
	assert.Equal(t, false, not.Operate(types.NewBoolean(true)).Value().(bool))
	assert.Equal(t, false, and.Operate(types.NewBoolean(false), types.NewBoolean(false)).Value().(bool))
	assert.Equal(t, false, and.Operate(types.NewBoolean(false), types.NewBoolean(true)).Value().(bool))
	assert.Equal(t, false, and.Operate(types.NewBoolean(true), types.NewBoolean(false)).Value().(bool))
	assert.Equal(t, true, and.Operate(types.NewBoolean(true), types.NewBoolean(true)).Value().(bool))
	assert.Equal(t, false, or.Operate(types.NewBoolean(false), types.NewBoolean(false)).Value().(bool))
	assert.Equal(t, true, or.Operate(types.NewBoolean(true), types.NewBoolean(false)).Value().(bool))
	assert.Equal(t, true, or.Operate(types.NewBoolean(false), types.NewBoolean(true)).Value().(bool))
	assert.Equal(t, true, or.Operate(types.NewBoolean(true), types.NewBoolean(true)).Value().(bool))
}
