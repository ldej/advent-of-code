package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	result := run()
	fmt.Println(result)
}
func run() int {
	input := tools.ReadStringsDoubleNewlines()
	var coordinates []image.Point
	for _, line := range strings.Split(input[0], "\n") {
		parts := strings.Split(line, ",")
		coordinates = append(coordinates, image.Point{X: myints.ToInt(parts[0]), Y: myints.ToInt(parts[1])})
	}

	for i, line := range strings.Split(input[1], "\n") {
		if i > 0 {
			break
		}
		parts := strings.Split(strings.Replace(line, "fold along ", "", -1), "=")

		axis := parts[0]
		value := myints.ToInt(parts[1])

		if axis == "x" {
			for i := range coordinates {
				if coordinates[i].X > value {
					coordinates[i].X = value - (coordinates[i].X - value)
				}
			}
		} else {
			for i := range coordinates {
				if coordinates[i].Y > value {
					coordinates[i].Y = value - (coordinates[i].Y - value)
				}
			}
		}
	}

	field := map[image.Point]bool{}
	for _, coordinate := range coordinates {
		field[coordinate] = true
	}
	return len(field)
}
