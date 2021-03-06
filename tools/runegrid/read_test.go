package runegrid

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
