package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code/2017/common"
)

func main() {
	scanner := common.ReadLines("./5/input1.txt")

	var numbers []int
	for scanner.Scan() {

		if number, err := strconv.Atoi(scanner.Text()); err == nil {
			numbers = append(numbers, number)
		}
	}

	location := 0
	length := len(numbers)
	step := 0

	for location >= 0 && location < length {
		step++
		value := numbers[location]
		numbers[location]++
		location += value
	}
	fmt.Println(step)
}
