package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	example1 := run(0, 20, "./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run(0, 4000000)
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(minLimit int, maxLimit int, file ...string) int {
	input := tools.ReadStrings(file...)

	sensors := map[tools.Point]int{}

	for _, line := range input {
		numbers := tools.FindInts(line)

		sensor := tools.Point{X: numbers[0], Y: numbers[1]}
		beacon := tools.Point{X: numbers[2], Y: numbers[3]}

		sensors[sensor] = sensor.ManhattanDistance(beacon)
	}

	for y := minLimit; y <= maxLimit; y++ {
		for x := minLimit; x <= maxLimit; x++ {
			target := tools.Point{X: x, Y: y}

			var inRange bool
			for sensor, strength := range sensors {
				distanceToTarget := sensor.ManhattanDistance(target)
				if distanceToTarget <= strength {
					inRange = true
					x += strength - distanceToTarget
					break
				}
			}
			if !inRange {
				return (target.X * 4000000) + target.Y
			}
		}
	}
	return -1
}
