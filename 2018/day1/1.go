package main

import (
	"fmt"
	"strconv"

	"github.com/ldej/advent-of-code-2018/common"
)

func main() {
	scanner := common.ReadLines("./day1/input.txt")

	frequency := 0
	for scanner.Scan() {
		a, _ := strconv.Atoi(scanner.Text())
		frequency += a
	}
	fmt.Println(frequency)
}
