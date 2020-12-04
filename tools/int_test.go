package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAndMax(t *testing.T) {
	min, max := MinAndMax([]int{1, -100, 2333, 4545})

	assert.Equal(t, -100, min)
	assert.Equal(t, 4545, max)
}

func TestManhattanDistance(t *testing.T) {
	distance := ManhattanDistance(1, 5, 14, 8)

	// 13 down + 3 right
	assert.Equal(t, 16, distance)
}

func TestIntsProduct(t *testing.T) {
	total := IntsProduct([]int{78, 178, 75, 86, 39})

	assert.Equal(t, 3492520200, total)
}

func TestIntsSum(t *testing.T) {
	total := IntsSum([]int{78, 178, 75, 86, 39})

	assert.Equal(t, 456, total)
}

func TestIntSlicesEqual(t *testing.T) {
	assert.True(t, IntSlicesEqual([]int{1, 2, 3}, []int{1, 2, 3}))
}

func TestInRange(t *testing.T) {
	assert.True(t, InRange(10, 1, 12))
	assert.False(t, InRange(1, 10, 20))
	assert.False(t, InRange(30, 10, 20))
}

func TestOutRange(t *testing.T) {
	assert.True(t, OutRange(1, 4, 12))
	assert.True(t, OutRange(25, 10, 20))
	assert.False(t, OutRange(15, 10, 20))
}

func TestIntLength(t *testing.T) {
	assert.Equal(t, 5, IntLength(12345))
}

func TestIntIndex(t *testing.T) {
	assert.Equal(t, 3, IntIndex(123, 2))
	assert.Equal(t, 4, IntIndex(123456789, 3))
}

func TestIntToSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, IntToSlice(123))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, IntToSlice(12345))
}
