package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntegers(t *testing.T) {
	ints := StringToIntegers("123456789")

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, expected, ints)
}

func TestStringSort(t *testing.T) {
	sorted := StringSort("LaurencedeJong")

	assert.Equal(t, "JLacdeeegnnoru", sorted)
}

func TestStringSlicesEqual(t *testing.T) {
	assert.True(t, StringSlicesEqual([]string{"Laurence", "de", "Jong"}, []string{"Laurence", "de", "Jong"}))
}

func TestStringSlicesNotEqual(t *testing.T) {
	assert.False(t, StringSlicesEqual([]string{"Laurence", "de", "dJong"}, []string{"Laurence", "de", "Jong"}))
}

func TestStringCountLetters(t *testing.T) {
	result := StringCountLetters("aabbbcccdddd")

	expected := map[string]int{
		"a": 2,
		"b": 3,
		"c": 3,
		"d": 4,
	}

	assert.Equal(t, expected, result)
}

func TestStringRemoveDuplicates(t *testing.T) {
	result := StringsRemoveDuplicates([]string{"Laurence", "de", "de", "Jong"})

	expected := []string{"Laurence", "de", "Jong"}

	assert.Equal(t, expected, result)
}

func TestStringsIntersection(t *testing.T) {
	result := StringsIntersection([]string{"ffekaasfsef", "friaagrjijqw", "aygbffdaywgbd", "adhvthwfd"})

	assert.Equal(t, []rune{'f', 'a'}, result)
}

func TestStringsUnion(t *testing.T) {
	result := StringsUnion([]string{"ffekaasfsef", "friaagrjijqw", "aygbffdaywgbd", "adhvthwfd"})

	assert.Equal(t, []rune("abdefghijkstywv"), result)
}
