package tools

import (
	"math"
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

func IntsProduct(ints []int) int {
	result := 1
	for _, i := range ints {
		result *= i
	}
	return result
}

func IntsSum(ints []int) int {
	result := 0
	for _, i := range ints {
		result += i
	}
	return result
}

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

func IntSliceContains(a []int, b int) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}
	return false
}
