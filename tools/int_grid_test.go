package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntGrid_At(t *testing.T) {
	grid := IntGrid{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 9, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
	}

	assert.Equal(t, 9, grid.At(5, 4))
}

func TestIntGrid_Set(t *testing.T) {
	grid := IntGrid{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 9, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
	}

	grid.Set(1, 2, 9)

	assert.Equal(t, 9, grid.At(1, 2))
}

func TestIntGrid_Grow(t *testing.T) {
	grid := IntGrid{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 9, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
	}

	var original = make(IntGrid, len(grid), len(grid))
	copy(original, grid)

	newGrid := grid.GrowAll(0)

	expected := IntGrid{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 9, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, original, grid)
	assert.Equal(t, expected, newGrid)
}

func TestIntGrid_PGrow(t *testing.T) {
	grid := IntGrid{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{1, 2, 3, 4, 9, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
	}

	grid.PGrow(0)

	expected := IntGrid{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 9, 6, 7, 0},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	assert.Equal(t, expected, grid)
}

// Copies the grid once for a grow in all directions
func BenchmarkIntGrid_GrowAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := IntGrid{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 9, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		}
		grid.GrowAll(0)
	}
}

// Grows the grid itself using pointers
func BenchmarkIntGrid_PGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := IntGrid{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 9, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		}
		grid.PGrow(0)
	}
}
