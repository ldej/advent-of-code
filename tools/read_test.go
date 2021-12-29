package tools

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func PrepareFile(name string) func() {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("../input.txt", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return func() {
		err = os.Remove("../input.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestReadInts(t *testing.T) {
	clean := PrepareFile("./testdata/ints.txt")
	defer clean()

	ints := ReadInts()

	expected := []int{1757, 1890, -1750, 1440}

	assert.Equal(t, expected, ints)
}

func TestReadIntSlice(t *testing.T) {
	clean := PrepareFile("./testdata/int_slice.txt")
	defer clean()

	ints := ReadIntSlice()

	expected := []int{10, -4, 8, 12}

	assert.Equal(t, expected, ints)
}

func TestReadIntSlices(t *testing.T) {
	clean := PrepareFile("./testdata/int_slices.txt")
	defer clean()

	ints := ReadIntSlices()

	expected := [][]int{
		{5806, 6444, 1281, 38},
		{74, 127, -226, 84, 174},
		{1332, 52, -54, 655, 56, 170},
		{5653, 236, 1944, 3807},
	}

	assert.Equal(t, expected, ints)
}

func TestReadIntCsv(t *testing.T) {
	clean := PrepareFile("./testdata/int_csv.txt")
	defer clean()

	ints := ReadIntCsv()

	expected := [][]int{
		{108, 350},
		{113, 132},
		{115, -53},
		{117, 175},
	}
	assert.Equal(t, expected, ints)
}

func TestReadIntCsvOneLine(t *testing.T) {
	clean := PrepareFile("./testdata/int_csv2.txt")
	defer clean()

	ints := ReadIntCsvOneLine()

	expected := []int{1, 0, 0, 3, 1, 1, -2, 3, 1, 3, 4, 3}
	assert.Equal(t, expected, ints)
}

func TestReadStrings(t *testing.T) {
	clean := PrepareFile("./testdata/strings.txt")
	defer clean()

	strings := ReadStrings()

	expected := []string{
		"rmyxgdlihczskunpfwbgqoeybv",
		"rmyxgdlksczskunpfwbjqkeatv",
		"rmybgdxibczskunpfwbjqoeatv",
		"rmyxgdlirczskuopfwbjqzeatv",
	}

	assert.Equal(t, expected, strings)
}

func TestReadStringSlices(t *testing.T) {
	clean := PrepareFile("./testdata/string_slices.txt")
	defer clean()

	strings := ReadStringSlices()

	expected := [][]string{
		{"pphsv", "ojtou", "brvhsj", "cer"},
		{"vxjnf", "fzqitnj", "uyfck"},
		{"caibh", "nfuk", "kfnu"},
		{"qiho", "qif", "eupwww", "avyglnj"},
	}

	assert.Equal(t, expected, strings)
}

func TestReadString(t *testing.T) {
	clean := PrepareFile("./testdata/string.txt")
	defer clean()

	str := ReadString()

	assert.Equal(t, "wWuUJjXxqQrqQmKBzZbZzLlkWNqQntxXBZzM", str)
}
