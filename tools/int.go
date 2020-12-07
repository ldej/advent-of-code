package tools

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func MaxList(ints []int) int {
	max := ints[0]
	for _, i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func MaxListIndex(ints []int) int {
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

func MaxVar(ints ...int) int {
	return MaxList(ints)
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
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

func MinVar(ints ...int) int {
	return MinList(ints)
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

// IntsSum adds up all the ints a slice
func IntsSum(ints []int) int {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
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

// IntSliceContains checks if an integer is part of a slice
func IntSliceContains(a []int, b int) bool {
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

// IntIndex the digit on index n of an integer
//
// example: IntIndex(12345, 1) == 2
func IntIndex(value int, index int) int {
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
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GreatestCommonDivisorSlice(numbers []int) int {
	var gcd = numbers[0]
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		gcd = GreatestCommonDivisor(gcd, number)
	}
	return gcd
}
