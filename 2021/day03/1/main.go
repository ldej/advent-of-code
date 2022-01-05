package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadStrings()
	length := len(input[0])

	gamma := make([]rune, length)
	epsilon := make([]rune, length)

	for i := 0; i < length; i++ {
		zeroes, ones := ZeroesOnes(input, i)

		if zeroes > ones {
			gamma[i] = '0'
			epsilon[i] = '1'
		} else {
			gamma[i] = '1'
			epsilon[i] = '0'
		}
	}

	g := tools.BinaryToInt(string(gamma))
	e := tools.BinaryToInt(string(epsilon))

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
