package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/myints"
)

func main() {
	example1 := run(10, "./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	result := run(2000000)
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(lineNumber int, file ...string) int {
	input := tools.ReadStrings(file...)

	sensors := map[tools.Point]int{}
	beacons := map[tools.Point]bool{}

	var min, max *tools.Point

	for _, line := range input {
		numbers := tools.FindInts(line)

		sensor := tools.Point{X: numbers[0], Y: numbers[1]}
		beacon := tools.Point{X: numbers[2], Y: numbers[3]}

		distance := sensor.ManhattanDistance(beacon)

		if min == nil || max == nil {
			min = &tools.Point{Y: sensor.Y - distance, X: sensor.X - distance}
			max = &tools.Point{Y: sensor.Y + distance, X: sensor.X + distance}
		} else {
			min.X = myints.Min(min.X, sensor.X-distance)
			min.Y = myints.Min(min.Y, sensor.Y-distance)
			max.X = myints.Max(max.X, sensor.X+distance)
			max.Y = myints.Max(max.Y, sensor.Y+distance)
		}

		sensors[sensor] = distance
		beacons[beacon] = true
	}

	var cannotContain int
	for x := min.X; x <= max.X; x++ {
		target := tools.Point{X: x, Y: lineNumber}
		_, sensorFound := sensors[target]
		_, beaconFound := beacons[target]
		if sensorFound || beaconFound {
			continue
		}
		for sensor, distance := range sensors {
			if sensor.ManhattanDistance(target) <= distance {
				cannotContain++
				break
			}
		}
	}

	return cannotContain
}
