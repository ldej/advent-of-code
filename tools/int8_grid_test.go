package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt8Grid_At(t *testing.T) {
	grid := Int8Grid{
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 9, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
	}

	assert.Equal(t, int8(9), grid.At(5, 4))
}

func TestInt8Grid_Set(t *testing.T) {
	grid := Int8Grid{
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 9, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
	}

	grid.Set(1, 2, 9)

	assert.Equal(t, int8(9), grid.At(1, 2))
}

func TestInt8Grid_Grow(t *testing.T) {
	grid := Int8Grid{
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 9, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
	}

	var original = make(Int8Grid, len(grid), len(grid))
	copy(original, grid)

	newGrid := grid.GrowAll(0)

	expected := Int8Grid{
		[]int8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 9, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, original, grid)
	assert.Equal(t, expected, newGrid)
}

func TestInt8Grid_PGrow(t *testing.T) {
	grid := Int8Grid{
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
		[]int8{1, 2, 3, 4, 9, 6, 7},
		[]int8{1, 2, 3, 4, 5, 6, 7},
	}

	grid.PGrow(0)

	expected := Int8Grid{
		[]int8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 9, 6, 7, 0},
		[]int8{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int8{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, expected, grid)
}

// Copies the grid once for a grow in all directions
func BenchmarkInt8Grid_GrowAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := Int8Grid{
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 9, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
		}
		grid.GrowAll(0)
	}
}

// Grows the grid itself using point8ers
func BenchmarkInt8Grid_PGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := Int8Grid{
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
			[]int8{1, 2, 3, 4, 9, 6, 7},
			[]int8{1, 2, 3, 4, 5, 6, 7},
		}
		grid.PGrow(0)
	}
}
