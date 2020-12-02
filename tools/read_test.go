package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInts(t *testing.T) {
	ints := ReadInts("./testdata/ints.txt")

	expected := []int{1757, 1890, -1750, 1440}

	assert.Equal(t, expected, ints)
}

func TestReadIntSlice(t *testing.T) {
	ints := ReadIntSlice("./testdata/int_slice.txt")

	expected := []int{10, -4, 8, 12}

	assert.Equal(t, expected, ints)
}

func TestReadIntSlices(t *testing.T) {
	ints := ReadIntSlices("./testdata/int_slices.txt")

	expected := [][]int{
		{5806, 6444, 1281, 38},
		{74, 127, -226, 84, 174},
		{1332, 52, -54, 655, 56, 170},
		{5653, 236, 1944, 3807},
	}

	assert.Equal(t, expected, ints)
}

func TestReadIntCsv(t *testing.T) {
	ints := ReadIntCsv("./testdata/int_csv.txt")

	expected := [][]int{
		{108, 350},
		{113, 132},
		{115, -53},
		{117, 175},
	}
	assert.Equal(t, expected, ints)
}

func TestReadStrings(t *testing.T) {
	strings := ReadStrings("./testdata/strings.txt")

	expected := []string{
		"rmyxgdlihczskunpfwbgqoeybv",
		"rmyxgdlksczskunpfwbjqkeatv",
		"rmybgdxibczskunpfwbjqoeatv",
		"rmyxgdlirczskuopfwbjqzeatv",
	}

	assert.Equal(t, expected, strings)
}

func TestReadStringSlices(t *testing.T) {
	strings := ReadStringSlices("./testdata/string_slices.txt")

	expected := [][]string{
		{"pphsv", "ojtou", "brvhsj", "cer"},
		{"vxjnf", "fzqitnj", "uyfck"},
		{"caibh", "nfuk", "kfnu"},
		{"qiho", "qif", "eupwww", "avyglnj"},
	}

	assert.Equal(t, expected, strings)
}

func TestReadRuneGrid(t *testing.T) {
	grid := ReadRuneGrid("./testdata/rune_grid.txt")

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