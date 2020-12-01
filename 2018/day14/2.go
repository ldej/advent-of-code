package main

import (
	"fmt"

	"github.com/ldej/advent-of-code-2018/common"
)

func main() {
	after := 846601
	recipes := []int{3, 7}
	elf1Idx := 0
	elf2Idx := 1

	afterDigits := splitInt(after)

	found := false

	for !found {

		sum := recipes[elf1Idx] + recipes[elf2Idx]
		digits := splitInt(sum)

		for _, digit := range digits {
			recipes = append(recipes, digit)

			lastN := recipes[common.Max(len(recipes)-len(afterDigits), 0):]
			if areEqual(lastN, afterDigits) {
				fmt.Println(len(recipes[:len(recipes)-len(afterDigits)]))
				found = true
				break
			}
		}
		if found {
			break
		}

		elf1Idx += recipes[elf1Idx] + 1
		for elf1Idx >= len(recipes) {
			elf1Idx -= len(recipes)
		}
		elf2Idx += recipes[elf2Idx] + 1
		for elf2Idx >= len(recipes) {
			elf2Idx -= len(recipes)
		}
	}

}

func areEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func splitInt(a int) []int {
	if a == 0 {
		return []int{0}
	}
	if a < 10 {
		return []int{a}
	}
	list := []int{}
	for a > 0 {

		list = append([]int{a % 10}, list...)
		a = a / 10
	}
	return list
}
