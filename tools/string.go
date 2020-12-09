package tools

import (
	"sort"
	"strconv"
	"strings"
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

func StringSort(a string) string {
	b := strings.Split(a, "")
	sort.Strings(b)
	return strings.Join(b, "")
}

func StringSlicesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func StringSliceContains(a []string, b string) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}
	return false
}

func StringCountLetters(a string) map[string]int {
	counts := map[string]int{}
	for i := 0; i < len(a); i++ {
		letter := string(a[i])
		if _, found := counts[letter]; found {
			counts[letter] += 1
		} else {
			counts[letter] = 1
		}
	}
	return counts
}

// StringsRemoveDuplicates remove duplicates while preserving order
func StringsRemoveDuplicates(a []string) []string {
	var result []string

	for _, c := range a {
		if !StringSliceContains(result, c) {
			result = append(result, c)
		}
	}
	return result
}

func StringRemoveDuplicates(a string) string {
	var result []rune
	for _, c := range a {
		if !strings.ContainsRune(string(result), c) {
			result = append(result, c)
		}
	}
	return string(result)
}

func StringsIntersection(strs []string) []rune {
	if len(strs) == 0 {
		return []rune{}
	}
	var result []rune
	first := strs[0]

	for _, c := range first {
		if AllStringsContain(strs, c) {
			result = append(result, c)
		}
	}
	return []rune(StringRemoveDuplicates(string(result)))
}

func AllStringsContain(strs []string, a rune) bool {
	for _, str := range strs {
		if !strings.ContainsRune(str, a) {
			return false
		}
	}
	return true
}

func StringsUnion(strs []string) []rune {
	var result []rune
	mapped := map[rune]bool{}

	for _, s := range strs {
		for _, c := range s {
			if _, found := mapped[c]; !found {
				result = append(result, c)
				mapped[c] = true
			}
		}
	}
	return result
}

func StringsFilter(strs []string, filter string) []string {
	var result []string
	for _, s := range strs {
		if s != filter {
			result = append(result, s)
		}
	}
	return result
}

// MapStrings applies a function to every string in a slice
// a new slice is returned
func MapStrings(strs []string, f func(index int, value string) string) []string {
	var result = make([]string, len(strs), len(strs))
	for i, v := range strs {
		result[i] = f(i, v)
	}
	return result
}
