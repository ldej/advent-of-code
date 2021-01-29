package tools

import (
	"fmt"
	"log"
)

func SetBit(input int, value int, position int) int {
	switch value {
	case 0:
		return input & ^(1 << position)
	case 1:
		return input | 1<<position
	default:
		log.Fatalf("can't set bit value '%d', only 0 or 1 allowed", value)
		return -1
	}
}

func ToBinary(input int) string {
	return fmt.Sprintf("%b", input)
}

func ToBinaryPadded(input int, size int) string {
	return fmt.Sprintf("%0*b", size, input)
}
