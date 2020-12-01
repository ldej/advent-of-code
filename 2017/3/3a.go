package main

import "fmt"

func main() {
	input := 361527

	width := 1
	steps := 0
	minLimit := 1

	for width*width <= input {
		minLimit = width * width
		width += 2
		steps += 1
	}

	maxLimit := width * width
	diff := maxLimit - minLimit

	extra := (input - minLimit) % (diff / 8)

	totalSteps := steps + extra

	fmt.Println(totalSteps)
}
