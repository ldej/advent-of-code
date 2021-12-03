package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/mystrings"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int64 {
	input := tools.ReadStrings("./2021/day03/input.txt")
	length := len(input[0])

	oxygen := input
	co2 := input

	for i := 0; i < length; i++ {
		zeroes, ones := ZeroesOnes(oxygen, i)

		oxygen = mystrings.FilterFunc(oxygen, func(line string) bool {
			return (zeroes > ones && line[i] == '0') ||
				(ones > zeroes && line[i] == '1') ||
				(ones == zeroes && line[i] == '1')
		})

		if len(oxygen) == 1 {
			break
		}
	}

	for i := 0; i < length; i++ {
		zeroes, ones := ZeroesOnes(co2, i)

		co2 = mystrings.FilterFunc(co2, func(line string) bool {
			return (zeroes < ones && line[i] == '0') ||
				(ones < zeroes && line[i] == '1') ||
				(ones == zeroes && line[i] == '0')
		})

		if len(co2) == 1 {
			break
		}
	}

	g := tools.BinaryToInt(oxygen[0])
	e := tools.BinaryToInt(co2[0])

	return g * e
}

func ZeroesOnes(input []string, position int) (int, int) {
	zeroes := 0
	ones := 0
	for _, line := range input {
		if line[position] == '0' {
			zeroes++
		} else {
			ones++
		}
	}
	return zeroes, ones
}
