package myrunes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt(t *testing.T) {
	assert.Equal(t, 1, ToInt('a'))
	assert.Equal(t, 26, ToInt('z'))
	assert.Equal(t, 27, ToInt('A'))
	assert.Equal(t, 52, ToInt('Z'))
}
