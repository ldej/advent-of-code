package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day17/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day17/input.txt")
	fmt.Println("Result:", result)
}

type Cube struct {
	X int
	Y int
	Z int
}

func run(input string) int {
	grid := tools.ReadRuneGrid(input)

	cubes := map[Cube]bool{}

	for cell := range grid.Cells() {
		cube := Cube{X: cell.X, Y: cell.Y, Z: 0}
		cubes[cube] = cell.Value == '#'
	}

	for i := 0; i < 6; i++ {
		cubes = iterate(cubes)
	}

	var activeCount = 0
	for _, value := range cubes {
		if value {
			activeCount += 1
		}
	}

	return activeCount
}

func iterate(cubes map[Cube]bool) map[Cube]bool {
	var newCubes = make(map[Cube]bool)
	directions := []int{-1, 0, 1}

	minX, maxX, minY, maxY, minZ, maxZ := Bounds(cubes)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				currentCube := Cube{X: x, Y: y, Z: z}
				isActive := cubes[currentCube]

				var activeCount int
				for _, i := range directions {
					for _, j := range directions {
						for _, k := range directions {
							if i == 0 && j == 0 && k == 0 {
								continue
							}
							neighbour := Cube{X: x + i, Y: y + j, Z: z + k}
							if isNeighbourActive := cubes[neighbour]; isNeighbourActive {
								activeCount++
							}
						}
					}
				}
				if (isActive && (activeCount == 2 || activeCount == 3)) || (!isActive && activeCount == 3) {
					newCubes[currentCube] = true
				}
			}
		}
	}
	return newCubes
}

func Bounds(cubes map[Cube]bool) (int, int, int, int, int, int) {
	f := AnyCube(cubes)
	var minX, maxX = f.X, f.X
	var minY, maxY = f.Y, f.Y
	var minZ, maxZ = f.Z, f.Z

	for cube := range cubes {
		if cube.X < minX {
			minX = cube.X
		} else if cube.X > maxX {
			maxX = cube.X
		}
		if cube.Y < minY {
			minY = cube.Y
		} else if cube.Y > maxY {
			maxY = cube.Y
		}
		if cube.Z < minZ {
			minZ = cube.Z
		} else if cube.Z > maxZ {
			maxZ = cube.Z
		}
	}
	return minX, maxX, minY, maxY, minZ, maxZ
}

func AnyCube(cubes map[Cube]bool) Cube {
	for cube := range cubes {
		return cube
	}
	return Cube{}
}
