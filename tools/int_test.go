package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAndMax(t *testing.T) {
	min, max := MinAndMax([]int{1, -100, 2333, 4545})

	assert.Equal(t, -100, min)
	assert.Equal(t, 4545, max)
}
