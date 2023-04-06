package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	example1 := run("./example1.txt")
	fmt.Printf("\nExample 1:\n%v\n", example1)

	example2 := run("./example2.txt")
	fmt.Printf("\nExample 2:\n%v\n", example2)

	result := run()
	fmt.Printf("\nFinal:\n%v\n", result)
}

func run(file ...string) int {
	input := tools.ReadIntCsv(file...)
	cubes := map[tools.Point3D]bool{}

	for _, line := range input {
		cubes[tools.Point3D{
			X: line[0],
			Y: line[1],
			Z: line[2],
		}] = true
	}

	var totalSurface int
	for cube := range cubes {
		surface := 6
		if cubes[tools.Point3D{X: cube.X + 1, Y: cube.Y, Z: cube.Z}] {
			surface -= 1
		}
		if cubes[tools.Point3D{X: cube.X - 1, Y: cube.Y, Z: cube.Z}] {
			surface -= 1
		}
		if cubes[tools.Point3D{X: cube.X, Y: cube.Y + 1, Z: cube.Z}] {
			surface -= 1
		}
		if cubes[tools.Point3D{X: cube.X, Y: cube.Y - 1, Z: cube.Z}] {
			surface -= 1
		}
		if cubes[tools.Point3D{X: cube.X, Y: cube.Y, Z: cube.Z + 1}] {
			surface -= 1
		}
		if cubes[tools.Point3D{X: cube.X, Y: cube.Y, Z: cube.Z - 1}] {
			surface -= 1
		}
		totalSurface += surface
	}
	return totalSurface
}
