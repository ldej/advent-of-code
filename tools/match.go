package tools

import (
	"regexp"
	"strings"
)

const (
	ReHexColor = `^#[a-f0-9]{6}$`
)

var (
	CompiledHexColorRegex = regexp.MustCompile(ReHexColor)
)

func IsHexColor(a string) bool {
	return CompiledHexColorRegex.MatchString(strings.ToLower(a))
}
