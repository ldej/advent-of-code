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
