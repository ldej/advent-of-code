package tools

import (
	"github.com/ldej/advent-of-code/tools/myints"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationsInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var result [][]int
	for res := range CombinationsInt(input, 2) {
		result = append(result, res)
	}

	expected := [][]int{
		{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8}, {1, 9}, {1, 10},
		{2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8}, {2, 9}, {2, 10},
		{3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9}, {3, 10},
		{4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9}, {4, 10},
		{5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10},
		{6, 7}, {6, 8}, {6, 9}, {6, 10},
		{7, 8}, {7, 9}, {7, 10},
		{8, 9}, {8, 10},
		{9, 10},
	}
	assert.Equal(t, expected, result)
}

func TestGenMapInts(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	var result []int
	for result = range GenMapInts(input, func(i, v int) int { return v * v }) {
		if myints.IndexOf(result, 9) >= 0 {
			break
		}
	}

	assert.Equal(t, []int{1, 4, 9, 4, 5}, result)
}
