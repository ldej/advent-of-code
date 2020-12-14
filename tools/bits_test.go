package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBit(t *testing.T) {
	assert.Equal(t, 1, SetBit(0, 1, 0))
	assert.Equal(t, 2, SetBit(0, 1, 1))
	assert.Equal(t, 3, SetBit(2, 1, 0))
	assert.Equal(t, 4, SetBit(0, 1, 2))

	assert.Equal(t, 0, SetBit(1, 0, 0))
	assert.Equal(t, 2, SetBit(3, 0, 0))
}

func TestToBinary(t *testing.T) {
	assert.Equal(t, "1", ToBinary(1))
	assert.Equal(t, "10", ToBinary(2))
	assert.Equal(t, "11", ToBinary(3))
	assert.Equal(t, "100", ToBinary(4))
}

func TestToBinaryPadded(t *testing.T) {
	assert.Equal(t, "0001", ToBinaryPadded(1, 4))
}
