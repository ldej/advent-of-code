package tools

import (
	"fmt"
	"log"
	"math"
	"strconv"
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

// IntsProduct multiplies all the ints in a slice
func IntsProduct(ints []int) int {
	result := 1
	for _, i := range ints {
		result *= i
	}
	return result
}

// IntsProductVar is the variadic version of IntsSum
func IntsProductVar(ints ...int) int {
	return IntsProduct(ints)
}

// IntsSum adds up all the ints a slice
func IntsSum(ints []int) int {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
}

// IntsSumVar is the variadic version of IntsSum
func IntsSumVar(ints ...int) int {
	return IntsSum(ints)
}

// IntSlicesEqual checks if two integer slices are equal
func IntSlicesEqual(a []int, b []int) bool {
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

// IntsContain checks if an integer is part of a slice
func IntsContain(a []int, b int) bool {
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

// IntLength length of and integer
//
// example: IntLength(12345) == 5
func IntLength(value int) int {
	return len(fmt.Sprintf("%d", value))
}

// IntDigitIndex the digit on index n of an integer
//
// example: IntDigitIndex(12345, 1) == 2
func IntDigitIndex(value int, index int) int {
	return IntToSlice(value)[index]
}

// IntToSlice convert an int to its separate digits
//
// example: IntToSlice(12345) == []int{1, 2, 3, 4, 5}
func IntToSlice(value int) []int {
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

// MapInts applies a function to every integer in a slice
// a new slice is returned
func MapInts(ints []int, f func(index, value int) int) []int {
	var newInts = make([]int, len(ints), len(ints))
	for i, v := range ints {
		newInts[i] = f(i, v)
	}
	return newInts
}

// IntsIndexOf returns the index of the first element that matches the value
func IntsIndexOf(ints []int, value int) int {
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

// IntsFilter returns a slice with all ints with value n removed
func IntsFilter(ints []int, n int) []int {
	var result []int
	for _, i := range ints {
		if i != n {
			result = append(result, n)
		}
	}
	return result
}

// IntsRemoveIndex removes the item at index while preserving the original slice
func IntsRemoveIndex(ints []int, index int) []int {
	tmp := make([]int, 0)
	tmp = append(tmp, ints[:index]...)
	return append(tmp, ints[index+1:]...)
}

func IntsAppendPreserve(ints []int, item int) []int {
	tmp := make([]int, 0)
	return append(append(tmp, ints...), item)
}

func IntsPrepend(ints []int, item int) []int {
	return append([]int{item}, ints...)
}

func IntsOutOfBounds(ints []int, index int) bool {
	return index < 0 || index >= len(ints)
}
