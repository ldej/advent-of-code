package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	input := tools.ReadString()
	result := run(input)
	fmt.Println(result)
}

func run(input string) int {
	binaryBuilder := strings.Builder{}
	for _, char := range input {
		binaryBuilder.WriteString(fmt.Sprintf("%04b", tools.HexToInt(string(char))))
	}
	binary := binaryBuilder.String()

	total, _ := parse(binary)
	return total
}

func parse(remaining string) (int, string) {
	var packetVersion, typeID int
	packetVersion, typeID, remaining = tools.BinaryToInt(remaining[:3]), tools.BinaryToInt(remaining[3:6]), remaining[6:]
	total := packetVersion

	switch typeID {
	case 4:
		dataBuilder := strings.Builder{}
		var notLast, value string
		for notLast != "0" {
			notLast, value, remaining = remaining[0:1], remaining[1:5], remaining[5:]
			dataBuilder.WriteString(value)
		}
		// fmt.Println(tools.BinaryToInt(dataBuilder.String()))
	default:
		var lengthTypeID string
		lengthTypeID, remaining = remaining[0:1], remaining[1:]

		var subTotal int
		switch lengthTypeID {
		case "0":
			// 15 bits
			var length int
			var subPackets string
			length, remaining = tools.BinaryToInt(remaining[:15]), remaining[15:]
			subPackets, remaining = remaining[:length], remaining[length:]
			for len(subPackets) > 0 {
				subTotal, subPackets = parse(subPackets)
				total += subTotal
			}
		case "1":
			// 11 bits
			var subPackets int
			subPackets, remaining = tools.BinaryToInt(remaining[:11]), remaining[11:]
			for i := 0; i < subPackets; i++ {
				subTotal, remaining = parse(remaining)
				total += subTotal
			}
		}
	}

	return total, remaining
}
