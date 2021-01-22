package mystrings

import (
	"sort"
	"strconv"
	"strings"
)

func Index(s string, i int) string {
	if i < 0 {
		return ""
	}
	if i > len(s) {
		return ""
	}
	return string([]rune(s)[i])
}

func CompareXOR(a string, b string, compare string) bool {
	return (a == compare || b == compare) && !(a == compare && b == compare)
}

func ToIntegers(a string) []int {
	var ints []int

	for _, char := range []rune(a) {
		i, _ := strconv.Atoi(string(char))
		ints = append(ints, i)
	}
	return ints
}

func Sort(a string) string {
	b := strings.Split(a, "")
	sort.Strings(b)
	return strings.Join(b, "")
}

func EqualSlices(a []string, b []string) bool {
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

func SliceContains(a []string, b string) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}
	return false
}

func CountLetters(a string) map[string]int {
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

// RemoveDuplicates remove duplicates while preserving order
func RemoveDuplicates(a []string) []string {
	var result []string

	for _, c := range a {
		if !SliceContains(result, c) {
			result = append(result, c)
		}
	}
	return result
}

func RemoveDuplicateLetters(a string) string {
	var result []rune
	for _, c := range a {
		if !strings.ContainsRune(string(result), c) {
			result = append(result, c)
		}
	}
	return string(result)
}

func Intersection(strs []string) []rune {
	if len(strs) == 0 {
		return []rune{}
	}
	var result []rune
	first := strs[0]

	for _, c := range first {
		if AllContain(strs, c) {
			result = append(result, c)
		}
	}
	return []rune(RemoveDuplicateLetters(string(result)))
}

func AllContain(strs []string, a rune) bool {
	for _, str := range strs {
		if !strings.ContainsRune(str, a) {
			return false
		}
	}
	return true
}

func Union(strs []string) []rune {
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

func Filter(strs []string, filter string) []string {
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

// RemoveAtIndex removes the item at index while preserving the original slice
func RemoveAtIndex(strs []string, index int) []string {
	tmp := make([]string, 0)
	tmp = append(tmp, strs[:index]...)
	return append(tmp, strs[index+1:]...)
}

func AppendPreserve(strs []string, item string) []string {
	tmp := make([]string, 0)
	return append(append(tmp, strs...), item)
}

func Prepend(strs []string, item string) []string {
	return append([]string{item}, strs...)
}

func Reverse(a string) string {
	var reversed string
	for _, v := range a {
		reversed = string(v) + reversed
	}
	return reversed
}
