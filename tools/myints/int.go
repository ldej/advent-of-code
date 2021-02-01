package myints

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func Max(ints ...int) int {
	return MaxSlice(ints)
}

func MaxSlice(ints []int) int {
	max := ints[0]
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func MaxSliceIndex(ints []int) int {
	max := ints[0]
	index := 0
	for i, value := range ints {
		if value > max {
			max = value
			index = i
		}
	}
	return index
}

func Min(ints ...int) int {
	return MinList(ints)
}

func MinList(ints []int) int {
	min := ints[0]
	for _, i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}

func MinListIndex(ints []int) int {
	min := ints[0]
	index := 0
	for i, value := range ints {
		if value < min {
			min = value
			index = i
		}
	}
	return index
}

func MinAndMax(ints []int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32

	for _, i := range ints {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	x := 0
	if x1 > x2 {
		x = x1 - x2
	} else {
		x = x2 - x1
	}
	y := 0
	if y1 > y2 {
		y = y1 - y2
	} else {
		y = y2 - y1
	}
	return x + y
}

// Product multiplies all the ints in a slice
func Product(ints []int) int {
	result := 1
	for _, i := range ints {
		result *= i
	}
	return result
}

// ProductVar is the variadic version of Sum
func ProductVar(ints ...int) int {
	return Product(ints)
}

// Sum adds up all the ints a slice
func Sum(ints []int) int {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
}

// IntsSum is the variadic version of Sum
func IntsSum(ints ...int) int {
	return Sum(ints)
}

// SlicesEqual checks if two integer slices are equal
func SlicesEqual(a []int, b []int) bool {
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

// SliceContains checks if an integer is part of a slice
func SliceContains(a []int, b int) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}
	return false
}

// ToInt converts a string to int and will log.Fatalf when it fails
func ToInt(a string) int {
	v, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("tools.ToInt: not an int: %q", a)
	}
	return v
}

func ToInts(strs []string) []int {
	var values []int
	for _, i := range strs {
		values = append(values, ToInt(i))
	}
	return values
}

func ToIntOr(a string, or int) int {
	v, err := strconv.Atoi(a)
	if err != nil {
		return or
	}
	return v
}

// InRange from min to max including min and max
func InRange(value int, min int, max int) bool {
	if min >= max {
		log.Fatalf("tools.InRange: min is bigger or equal to max: min=%d max=%d", min, max)
	}
	if value < min || value > max {
		return false
	}
	return true
}

// OutRange lower than min or higher than max
func OutRange(value int, min int, max int) bool {
	if min >= max {
		log.Fatalf("tools.OutRange: min is bigger or equal to max: min=%d max=%d", min, max)
	}
	if value >= min && value <= max {
		return false
	}
	return true
}

// Length length of and integer
//
// example: Length(12345) == 5
func Length(value int) int {
	return len(fmt.Sprintf("%d", value))
}

// Index the digit on index n of an integer
//
// example: Index(12345, 1) == 2
func Index(value int, index int) int {
	return ToDigits(value)[index]
}

// ToDigits convert an int to its separate digits
//
// example: ToDigits(12345) == []int{1, 2, 3, 4, 5}
func ToDigits(value int) []int {
	str := strconv.Itoa(value)
	var ints []int
	for _, i := range str {
		a, _ := strconv.Atoi(string(i))
		ints = append(ints, a)
	}
	return ints
}

func GreatestCommonDivisor(a, b int) int {
	// Euclidean
	if a < b {
		a, b = b, a
	}
	quotient := a / b
	remainder := a - (quotient * b)
	if remainder == 0 {
		return b
	}

	return GreatestCommonDivisor(b, remainder)
}

func GreatestCommonDivisorSlice(numbers []int) int {
	var gcd = numbers[0]
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		gcd = GreatestCommonDivisor(gcd, number)
	}
	return gcd
}

// Map applies a function to every integer in a slice
// a new slice is returned
func Map(ints []int, f func(index, value int) int) []int {
	var newInts = make([]int, len(ints), len(ints))
	for i, v := range ints {
		newInts[i] = f(i, v)
	}
	return newInts
}

// IndexOf returns the index of the first element that matches the value
func IndexOf(ints []int, value int) int {
	for index, v := range ints {
		if v == value {
			return index
		}
	}
	return -1
}

// IntsNonN returns the first integer that is not n, otherwise -1
func IntsNonN(ints []int, n int) int {
	for _, v := range ints {
		if v != n {
			return v
		}
	}
	return -1
}

// Filter returns a slice with all ints with value n removed
func Filter(ints []int, n int) []int {
	var result []int
	for _, i := range ints {
		if i != n {
			result = append(result, i)
		}
	}
	return result
}

// RemoveIndex removes the item at index while preserving the original slice
func RemoveIndex(ints []int, index int) []int {
	tmp := make([]int, 0)
	tmp = append(tmp, ints[:index]...)
	return append(tmp, ints[index+1:]...)
}

func AppendPreserve(ints []int, item int) []int {
	tmp := make([]int, 0)
	return append(append(tmp, ints...), item)
}

func Prepend(ints []int, item int) []int {
	return append([]int{item}, ints...)
}

func ParseCsv(a string) [][]int {
	var result [][]int

	lines := strings.Split(a, "\n")
	for _, line := range lines {
		var ints []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			ints = append(ints, ToInt(part))
		}
		result = append(result, ints)
	}
	return result
}

func ToCsv(ints []int) string {
	var s []string
	for _, i := range ints {
		s = append(s, strconv.Itoa(i))
	}
	return strings.Join(s, ", ")
}

func RemoveDuplicates(ints []int) []int {
	var result []int

	for _, i := range ints {
		if !SliceContains(result, i) {
			result = append(result, i)
		}
	}
	return result
}

func IsOdd(i int) bool {
	return i&1 == 1
}

func IsEven(i int) bool {
	return i&1 == 0
}
