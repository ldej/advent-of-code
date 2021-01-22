package runegrid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadRuneGrid(t *testing.T) {
	grid := Read("../testdata/rune_grid.txt")

	expected := RuneGrid{
		[]rune("#######"),
		[]rune("#.G...#"),
		[]rune("#...EG#"),
		[]rune("#.#.#G#"),
		[]rune("#..G#E#"),
		[]rune("#.....#"),
		[]rune("#######"),
	}

	assert.Equal(t, expected, grid)
}
