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

func TestRegexNamedGroupsRepeat(t *testing.T) {
	input := "dim white bags contain 5 shiny indigo bags, 4 posh tan bags, 3 faded blue bags."
	regex := `(?P<count>\d+) (?P<color>.*?) bag`
	result := RegexNamedGroupsRepeat(input, regex)

	expected := []map[string]string{
		{"count": "5", "color": "shiny indigo"},
		{"count": "4", "color": "posh tan"},
		{"count": "3", "color": "faded blue"},
	}
	assert.Equal(t, expected, result)
}

func TestRegexNamedGroups(t *testing.T) {
	input := "#1 @ 342,645: 25x20"
	regex := `^#(?P<id>\d+) @ (?P<left>\d+),(?P<top>\d+): (?P<width>\d+)x(?P<height>\d+)$`
	result := RegexNamedGroups(input, regex)

	expected := map[string]string{
		"id":     "1",
		"left":   "342",
		"top":    "645",
		"width":  "25",
		"height": "20",
	}
	assert.Equal(t, expected, result)
}
