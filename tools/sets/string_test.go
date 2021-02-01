package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSet_Intersect(t *testing.T) {
	a := NewStringSet().Add("one", "two", "three")
	b := NewStringSet().Add("three", "four", "five")
	expected := NewStringSet().Add("three")

	assert.Equal(t, expected, a.Intersect(b))
	assert.Equal(t, expected, b.Intersect(a))
}
