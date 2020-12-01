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

	for noun := 0; noun < 100; noun += 1 {
		for verb := 0; verb < 100; verb += 1 {
			valuesCopy := make([]int, len(values))
			copy(valuesCopy, values)

			if run(valuesCopy, noun, verb) == 19690720 {
				fmt.Println(noun, verb)
			}
		}
	}
}

func run(values []int, noun int, verb int) int {
	values[1] = noun
	values[2] = verb

	for i := 0; ; i += 4 {
		instruction, address1, address2, addressDestination := values[i], values[i+1], values[i+2], values[i+3]
		switch instruction {
		case 1:
			values[addressDestination] = values[address1] + values[address2]
		case 2:
			values[addressDestination] = values[address1] * values[address2]
		case 99:
			return values[0]
		default:
			log.Fatal(instruction)
		}
	}
}
