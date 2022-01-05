package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
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

	_, result := parse(binary)
	return result
}

func parse(remaining string) (string, int) {
	var packetVersion, typeID int
	packetVersion, typeID, remaining = tools.BinaryToInt(remaining[:3]), tools.BinaryToInt(remaining[3:6]), remaining[6:]
	if packetVersion < 0 {
		// don't need packet version
	}

	if typeID == 4 {
		dataBuilder := strings.Builder{}
		var notLast, valuePart string
		for notLast != "0" {
			notLast, valuePart, remaining = remaining[0:1], remaining[1:5], remaining[5:]
			dataBuilder.WriteString(valuePart)
		}
		return remaining, tools.BinaryToInt(dataBuilder.String())
	}

	var lengthTypeID string
	lengthTypeID, remaining = remaining[0:1], remaining[1:]

	var values []int

	switch lengthTypeID {
	case "0":
		// 15 bits
		var length, value int
		var subPackets string
		length, remaining = tools.BinaryToInt(remaining[:15]), remaining[15:]
		subPackets, remaining = remaining[:length], remaining[length:]
		for len(subPackets) > 0 {
			subPackets, value = parse(subPackets)
			values = append(values, value)
		}
	case "1":
		// 11 bits
		var subPackets, value int
		subPackets, remaining = tools.BinaryToInt(remaining[:11]), remaining[11:]
		for i := 0; i < subPackets; i++ {
			remaining, value = parse(remaining)
			values = append(values, value)
		}
	}

	var result int
	switch typeID {
	case 0:
		result = myints.Sum(values)
	case 1:
		result = myints.Product(values)
	case 2:
		result = myints.Min(values...)
	case 3:
		result = myints.Max(values...)
	case 5:
		if values[0] > values[1] {
			result = 1
		}
	case 6:
		if values[0] < values[1] {
			result = 1
		}
	case 7:
		if values[0] == values[1] {
			result = 1
		}
	}

	return remaining, result
}
