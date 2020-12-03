package tools

import (
	"fmt"
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

func TestIntGrid_WindowSum(t *testing.T) {
	grid := IntGrid{
		[]int{70, 0, 0, 0, 0, 0, 0},
		[]int{0, 5, 0, 0, 0, 0, 0},
		[]int{0, 5, 0, 0, 0, 1, 0},
		[]int{0, 0, 20, 200, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 190},
	}

	sum := grid.WindowSum(2, 2, 2, 2)
	assert.Equal(t, 220, sum)

	// out of bounds check
	sum = grid.WindowSum(4, 4, 3, 3)
	assert.Equal(t, 390, sum)

}

func TestIntGrid_MaxWindowSum(t *testing.T) {
	var tests = []struct {
		windowHeight int
		windowWidth  int

		x   int
		y   int
		sum int
	}{
		{1, 1, 3, 3, 200},
		{2, 2, 2, 2, 220},
		{3, 3, 1, 1, 230},
		{4, 4, 3, 3, 390},
		{5, 5, 2, 2, 412},
		{6, 6, 1, 1, 422},
		{7, 7, 0, 0, 492},
		{2, 5, 2, 1, 226},
		{6, 2, 1, 2, 221},
	}

	grid := IntGrid{
		[]int{70, 0, 0, 0, 0, 0, 0},
		[]int{0, 5, 0, 0, 0, 0, 0},
		[]int{0, 5, 0, 0, 0, 1, 0},
		[]int{0, 0, 20, 200, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0, 0, 190},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%dx%d", tt.windowHeight, tt.windowWidth), func(t *testing.T) {
			x, y, sum := grid.MaxWindowSum(tt.windowHeight, tt.windowWidth)
			if tt.x != x || tt.y != y || tt.sum != sum {
				t.Errorf("got x=%d y=%d sum=%d, want x=%d y=%d sum=%d", x, y, sum, tt.x, tt.y, tt.sum)
			}
		})
	}
}
