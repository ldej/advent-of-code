package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code-2018/common"
)

func main() {

	frequenciesMap := map[int]bool{}

	frequency := 0
	for true {
		scanner := common.ReadLines("./day1/input.txt")
		for scanner.Scan() {
			a, _ := strconv.Atoi(scanner.Text())
			frequency += a
			if _, found := frequenciesMap[frequency]; found {
				fmt.Println(frequency)
				return
			} else {
				frequenciesMap[frequency] = true
			}
		}
	}
}
