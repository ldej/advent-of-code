package runegrid

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

func TestRuneGrid_Transpose(t *testing.T) {
	grid := RuneGrid{
		[]rune("1..2"),
		[]rune("===="),
		[]rune("||||"),
		[]rune("----"),
		[]rune("####"),
		[]rune("3**4"),
	}

	expected := RuneGrid{
		[]rune("1=|-#3"),
		[]rune(".=|-#*"),
		[]rune(".=|-#*"),
		[]rune("2=|-#4"),
	}
	assert.Equal(t, expected, grid.Transpose())
}

func TestRuneGrid_Rotate(t *testing.T) {
	grid := RuneGrid{
		[]rune("1..2"),
		[]rune("===="),
		[]rune("||||"),
		[]rune("----"),
		[]rune("####"),
		[]rune("3**4"),
	}

	rotated := grid.Rotate(90)
	expected := RuneGrid{
		[]rune("3#-|=1"),
		[]rune("*#-|=."),
		[]rune("*#-|=."),
		[]rune("4#-|=2"),
	}
	assert.Equal(t, expected, rotated)

	expected = RuneGrid{
		[]rune("4**3"),
		[]rune("####"),
		[]rune("----"),
		[]rune("||||"),
		[]rune("===="),
		[]rune("2..1"),
	}
	rotated = rotated.Rotate(90)
	assert.Equal(t, expected, rotated)

	expected = RuneGrid{
		[]rune("3#-|=1"),
		[]rune("*#-|=."),
		[]rune("*#-|=."),
		[]rune("4#-|=2"),
	}
	rotated = rotated.Rotate(-90)
	assert.Equal(t, expected, rotated)

	rotated = grid.Rotate(180)
	expected = RuneGrid{
		[]rune("4**3"),
		[]rune("####"),
		[]rune("----"),
		[]rune("||||"),
		[]rune("===="),
		[]rune("2..1"),
	}
	assert.Equal(t, expected, rotated)

	rotated = grid.Rotate(-180)
	expected = RuneGrid{
		[]rune("4**3"),
		[]rune("####"),
		[]rune("----"),
		[]rune("||||"),
		[]rune("===="),
		[]rune("2..1"),
	}
	assert.Equal(t, expected, rotated)

	rotated = grid.Rotate(-270)
	expected = RuneGrid{
		[]rune("3#-|=1"),
		[]rune("*#-|=."),
		[]rune("*#-|=."),
		[]rune("4#-|=2"),
	}
	assert.Equal(t, expected, rotated)

	rotated = grid.Rotate(270)
	expected = RuneGrid{
		[]rune("2=|-#4"),
		[]rune(".=|-#*"),
		[]rune(".=|-#*"),
		[]rune("1=|-#3"),
	}
	assert.Equal(t, expected, rotated)
}

func TestRuneGrid_FlipVertical(t *testing.T) {
	grid := RuneGrid{
		[]rune("1..2"),
		[]rune("===="),
		[]rune("||||"),
		[]rune("----"),
		[]rune("####"),
		[]rune("3**4"),
	}

	expected := RuneGrid{
		[]rune("3**4"),
		[]rune("####"),
		[]rune("----"),
		[]rune("||||"),
		[]rune("===="),
		[]rune("1..2"),
	}

	assert.Equal(t, expected, grid.FlipVertical())
}

func TestRuneGrid_FlipHorizontal(t *testing.T) {
	grid := RuneGrid{
		[]rune("2=|-#4"),
		[]rune(".=|-#*"),
		[]rune(".=|-#*"),
		[]rune("1=|-#3"),
	}

	expected := RuneGrid{
		[]rune("4#-|=2"),
		[]rune("*#-|=."),
		[]rune("*#-|=."),
		[]rune("3#-|=1"),
	}

	assert.Equal(t, expected, grid.FlipHorizontal())
}

func TestRuneGrid_Window_3x3_1(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#X#G#"),
		[]rune("#..G#E#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
	expected := RuneWindow{
		Grid: RuneGrid{
			[]rune("..E"),
			[]rune("#X#"),
			[]rune(".G#"),
		},
		X:       3,
		Y:       3,
		CenterX: 1,
		CenterY: 1,
	}
	assert.Equal(t, expected, grid.Window(3, 3, 3, 3))
}

func TestRuneGrid_Window_3x3_2(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#X#G#"),
		[]rune("#..G#E#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
	expected := RuneWindow{
		Grid: RuneGrid{
			[]rune("##"),
			[]rune("#."),
		},
		X:       0,
		Y:       0,
		CenterX: 0,
		CenterY: 0,
	}
	assert.Equal(t, expected, grid.Window(3, 3, 0, 0))
}

func TestRuneGrid_Window_5x5_1(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#X#G#"),
		[]rune("#..G#E#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
	expected := RuneWindow{
		Grid: RuneGrid{
			[]rune(".G..."),
			[]rune("...EG"),
			[]rune(".#X#G"),
			[]rune("..G#E"),
			[]rune("....."),
		},
		X:       3,
		Y:       3,
		CenterX: 2,
		CenterY: 2,
	}
	assert.Equal(t, expected, grid.Window(5, 5, 3, 3))
}

func TestRuneGrid_Window_5x5_2(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#X#G#"),
		[]rune("#..G#E#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
	expected := RuneWindow{
		Grid: RuneGrid{
			[]rune("####"),
			[]rune("#.G."),
			[]rune("#..."),
			[]rune("#.#X"),
		},
		X:       1,
		Y:       1,
		CenterX: 1,
		CenterY: 1,
	}
	assert.Equal(t, expected, grid.Window(5, 5, 1, 1))
}

func TestRuneGrid_Window_5x5_3(t *testing.T) {
	grid := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#X#G#"),
		[]rune("#..G#E#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}
	expected := RuneWindow{
		Grid: RuneGrid{
			[]rune("X#G#"),
			[]rune("G#E#"),
			[]rune("...#"),
			[]rune("####"),
		},
		X:       5,
		Y:       5,
		CenterX: 2,
		CenterY: 2,
	}
	assert.Equal(t, expected, grid.Window(5, 5, 5, 5))
}
