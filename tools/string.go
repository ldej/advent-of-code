package tools

import (
	"strconv"
)

func StringIndex(s string, i int) string {
	if i < 0 {
		return ""
	}
	if i > len(s) {
		return ""
	}
	return string([]rune(s)[i])
}

func StringCompareXOR(a string, b string, compare string) bool {
	return (a == compare || b == compare) && !(a == compare && b == compare)
}

func StringToIntegers(a string) []int {
	var ints []int

	for _, char := range []rune(a) {
		i, _ := strconv.Atoi(string(char))
		ints = append(ints, i)
	}
	return ints
}