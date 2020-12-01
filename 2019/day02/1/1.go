package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code-2019/common"
)

func main() {
	scanner := common.ReadLines("./day02/input.txt")

	line := ""
	for scanner.Scan() {
		line = scanner.Text()
	}
	string_values := strings.Split(line, ",")

	values := []int{}
	for _, a := range string_values {
		val, _ := strconv.Atoi(a)
		values = append(values, val)
	}

	values[1] = 12
	values[2] = 2

loop:
	for i := 0; ; i += 4 {

		instruction, address1, address2, addressDestination := values[i], values[i+1], values[i+2], values[i+3]
		switch instruction {
		case 1:
			values[addressDestination] = values[address1] + values[address2]
		case 2:
			values[addressDestination] = values[address1] * values[address2]
		case 99:
			break loop
		default:
			log.Fatal(instruction)
		}
	}

	fmt.Println(values[0])
}
