package main

import (
		"fmt"
)

func main() {
	after := 846601
	recipes :=  []int{3, 7}
	elf1Idx := 0
	elf2Idx := 1

	for len(recipes) < after*2 + 2 {

		sum := recipes[elf1Idx] + recipes[elf2Idx]
		digits := splitInt(sum)
		recipes = append(recipes, digits...)

		elf1Idx = (elf1Idx + recipes[elf1Idx] + 1) % len(recipes)
		elf2Idx = (elf2Idx + recipes[elf2Idx] + 1) % len(recipes)

		for i := 0; i < len(recipes); i++ {
			if i == elf1Idx {
				fmt.Printf("(%d)", recipes[i])
			} else if i == elf2Idx {
				fmt.Printf("[%d]", recipes[i])
			} else {
				fmt.Printf(" %d ", recipes[i])
			}
		}
		fmt.Println()
	}

	nrs := recipes[after:after+10]
	for _, nr := range nrs {
		fmt.Printf("%d", nr)
	}
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
