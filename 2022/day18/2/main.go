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

func Neighbors(cube tools.Point3D) []tools.Point3D {
	return []tools.Point3D{
		{X: cube.X + 1, Y: cube.Y, Z: cube.Z},
		{X: cube.X - 1, Y: cube.Y, Z: cube.Z},
		{X: cube.X, Y: cube.Y + 1, Z: cube.Z},
		{X: cube.X, Y: cube.Y - 1, Z: cube.Z},
		{X: cube.X, Y: cube.Y, Z: cube.Z + 1},
		{X: cube.X, Y: cube.Y, Z: cube.Z - 1},
	}
}

func run(file ...string) int {
	input := tools.ReadIntCsv(file...)
	cubes := map[tools.Point3D]bool{}

	max := tools.Point3D{}
	for _, line := range input {
		cube := tools.Point3D{
			X: line[0],
			Y: line[1],
			Z: line[2],
		}
		cubes[cube] = true
		if cube.X > max.X {
			max.X = cube.X
		}
		if cube.Y > max.Y {
			max.Y = cube.Y
		}
		if cube.Z > max.Z {
			max.Z = cube.Z
		}
	}
	// Make one bigger for flooding algorithm
	max.X += 1
	max.Y += 1
	max.Z += 1

	queue := map[tools.Point3D]bool{tools.Point3D{X: 0, Y: 0, Z: 0}: true}
	visited := make(map[tools.Point3D]bool)

	for len(queue) > 0 {
		var next tools.Point3D
		for next = range queue {
			break
		}
		delete(queue, next)

		for _, n := range Neighbors(next) {
			outOfBounds := n.X < 0 || n.Y < 0 || n.Z < 0 || n.X > max.X || n.Y > max.Y || n.Z > max.Z
			if cubes[n] || visited[n] || outOfBounds {
				continue
			}
			queue[n] = true
		}
		visited[next] = true
	}

	// if it is not visited from the outside, and it's not a cube, then it's trapped
	trapped := map[tools.Point3D]bool{}
	for x := 0; x <= max.X; x++ {
		for y := 0; y <= max.Y; y++ {
			for z := 0; z <= max.Z; z++ {
				c := tools.Point3D{X: x, Y: y, Z: z}
				if visited[c] || cubes[c] {
					continue
				}
				trapped[c] = true
			}
		}
	}

	var totalSurface int
	for cube := range cubes {
		surface := 6
		for _, n := range Neighbors(cube) {
			if cubes[n] || trapped[n] {
				surface -= 1
			}
		}
		totalSurface += surface
	}

	return totalSurface
}
