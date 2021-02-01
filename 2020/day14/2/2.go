package main

import (
	"fmt"
	"github.com/ldej/advent-of-code/tools/myints"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day14/example2.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day14/input.txt")
	fmt.Println("Result:", result)
}

func run(input string) int {
	lines := tools.ReadStrings(input)

	var memory = make(map[int]int)
	var mask string

	for _, line := range lines {
		instruction := strings.Split(line, " = ")
		if instruction[0] == "mask" {
			mask = instruction[1]
		} else {
			memoryAddress, value := tools.FindInt(instruction[0]), myints.ToInt(instruction[1])

			memoryMask := createMemoryMask(memoryAddress, mask)

			addresses := getAddresses(memoryAddress, memoryMask)

			for _, addr := range addresses {
				memory[addr] = value
			}
		}
	}

	return tools.MapSumValues(memory)
}

func getAddresses(memoryAddress int, memoryMask string) []int {
	var addresses = []int{memoryAddress}

	for i, bit := range memoryMask {
		pos := len(memoryMask) - 1 - i
		switch bit {
		case '1':
			for j, addr := range addresses {
				addresses[j] = tools.SetBit(addr, 1, pos)
			}
		case 'X':
			var updated []int
			for _, addr := range addresses {
				updated = append(updated, tools.SetBit(addr, 0, pos))
				updated = append(updated, tools.SetBit(addr, 1, pos))
			}
			addresses = updated
		}
	}
	return addresses
}

func createMemoryMask(memoryAddress int, mask string) string {
	var memoryMask = make([]rune, len(mask))
	ints := tools.ToBinaryPadded(memoryAddress, len(mask))

	for i, bit := range mask {
		switch bit {
		case '0':
			memoryMask[i] = rune(ints[i])
		default:
			memoryMask[i] = bit
		}
	}
	return string(memoryMask)
}
