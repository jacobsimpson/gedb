package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	assert := assert.New(t)
	v := NewVarInt(0)
	assert.Equal(9, len(v.value))
	assert.Equal(int64(0), v.ToInt64())
}

func TestOne(t *testing.T) {
	assert := assert.New(t)
	v := NewVarInt(1)
	assert.Equal(9, len(v.value))
	assert.Equal(int64(1), v.ToInt64())
}

func TestNegativeOne(t *testing.T) {
	assert := assert.New(t)
	v := NewVarInt(-1)
	assert.Equal(9, len(v.value))
	//assert.Equal(int64(-1), v.ToInt64())
}

func TestOneTwentyEight(t *testing.T) {
	assert := assert.New(t)
	v := NewVarInt(127)
	assert.Equal(9, len(v.value))
	//assert.Equal(int64(127), v.ToInt64())
}

func TestOneTwentyNine(t *testing.T) {
	assert := assert.New(t)
	v := NewVarInt(128)
	assert.Equal(9, len(v.value))
	//assert.Equal(int64(128), v.ToInt64())
}
