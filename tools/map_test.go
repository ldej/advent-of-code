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
