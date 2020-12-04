package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHexColor(t *testing.T) {
	assert.True(t, IsHexColor("#123abc"))
	assert.True(t, IsHexColor("#AA99FF"))
	assert.False(t, IsHexColor("ddfgjfdg"))
}
