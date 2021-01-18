package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSum(t *testing.T) {
	assert.Equal(t, 3, MapSumValues(map[int]int{
		1: 1,
		2: 2,
	}))
	assert.Equal(t, 3, MapSumValues(map[string]int{
		"1": 1,
		"2": 2,
	}))
}

func TestMapContainsKeys(t *testing.T) {
	stringMap := map[string]int{
		"1": 1,
		"2": 10,
		"3": 3,
	}
	stringKeys := []string{"1", "2"}

	assert.True(t, MapHasKeys(stringMap, stringKeys))

	intMap := map[int]int{
		1: 1,
		2: 10,
		3: 3,
	}
	intKeys := []int{1, 2}

	assert.True(t, MapHasKeys(intMap, intKeys))
}
