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