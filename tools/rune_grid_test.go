package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneGrid_At(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#.#G#"),
		[]rune("#..G#E#"),
		[]rune("#...X.#"),
		[]rune("#######"),
	}

	assert.Equal(t, 'X', grid.At(5, 4))
}

func TestRuneGrid_Set(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#.#G#"),
		[]rune("#..G#E#"),
		[]rune("#...X.#"),
		[]rune("#######"),
	}

	grid.Set(1, 2, 'X')

	assert.Equal(t, 'X', grid.At(1, 2))
}

func TestRuneGrid_Grow(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#.#G#"),
		[]rune("#..G#E#"),
		[]rune("#...X.#"),
		[]rune("#######"),
	}
	var original = make(RuneGrid, len(grid), len(grid))
	copy(original, grid)

	newGrid := grid.GrowAll('*')

	expected := RuneGrid{
		[]rune("*********"),
		[]rune("*#######*"),
		[]rune("*#.G...#*"),
		[]rune("*#...EG#*"),
		[]rune("*#.#.#G#*"),
		[]rune("*#..G#E#*"),
		[]rune("*#...X.#*"),
		[]rune("*#######*"),
		[]rune("*********"),
	}

	assert.Equal(t, original, grid)
	assert.Equal(t, expected, newGrid)
}

func TestRuneGrid_PGrow(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#.#G#"),
		[]rune("#..G#E#"),
		[]rune("#...X.#"),
		[]rune("#######"),
	}

	grid.PGrow('*')

	expected := RuneGrid{
		[]rune("*********"),
		[]rune("*#######*"),
		[]rune("*#.G...#*"),
		[]rune("*#...EG#*"),
		[]rune("*#.#.#G#*"),
		[]rune("*#..G#E#*"),
		[]rune("*#...X.#*"),
		[]rune("*#######*"),
		[]rune("*********"),
	}

	assert.Equal(t, expected, grid)
}

// Copies the grid once for a grow in all directions
func BenchmarkRuneGrid_GrowAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := RuneGrid{
			[]rune("#######"),
			[]rune("#.G...#"),
			[]rune("#...EG#"),
			[]rune("#.#.#G#"),
			[]rune("#..G#E#"),
			[]rune("#...X.#"),
			[]rune("#######"),
		}

		grid.GrowAll('*')
	}
}

// Grows the grid itself using pointers
func BenchmarkRuneGrid_PGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid := RuneGrid{
			[]rune("#######"),
			[]rune("#.G...#"),
			[]rune("#...EG#"),
			[]rune("#.#.#G#"),
			[]rune("#..G#E#"),
			[]rune("#...X.#"),
			[]rune("#######"),
		}

		grid.PGrow('*')
	}
}

func TestRuneGrid_Count(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#.#G#"),
		[]rune("#..G#E#"),
		[]rune("#...X.#"),
		[]rune("#######"),
	}

	assert.Equal(t, 27, grid.Count('#'))
}
