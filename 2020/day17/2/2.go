package main

import (
	"fmt"

	"github.com/ldej/advent-of-code/tools/runegrid"
)

func main() {
	fmt.Println("Part 2")

	result := run("./2020/day17/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day17/input.txt")
	fmt.Println("Result:", result)
}

type Cube struct {
	X int
	Y int
	Z int
	W int
}

func run(input string) int {
	grid := runegrid.Read(input)

	cubes := map[Cube]bool{}

	for cell := range grid.Cells() {
		cube := Cube{X: cell.X, Y: cell.Y, Z: 0, W: 0}
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

	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := Bounds(cubes)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				for w := minW - 1; w <= maxW+1; w++ {
					currentCube := Cube{X: x, Y: y, Z: z, W: w}
					isActive := cubes[currentCube]

					var activeNeighbours int
					for _, i := range directions {
						for _, j := range directions {
							for _, k := range directions {
								for _, l := range directions {
									if i == 0 && j == 0 && k == 0 && l == 0 {
										continue
									}
									neighbour := Cube{X: x + i, Y: y + j, Z: z + k, W: w + l}
									if isNeighbourActive := cubes[neighbour]; isNeighbourActive {
										activeNeighbours++
									}
								}
							}
						}
					}
					if (isActive && (activeNeighbours == 2 || activeNeighbours == 3)) || (!isActive && activeNeighbours == 3) {
						newCubes[currentCube] = true
					}
				}
			}
		}
	}
	return newCubes
}

func Bounds(cubes map[Cube]bool) (int, int, int, int, int, int, int, int) {
	f := AnyCube(cubes)
	var minX, maxX, minY, maxY, minZ, maxZ, minW, maxW = f.X, f.X, f.Y, f.Y, f.Z, f.Z, f.W, f.W

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
		if cube.W < minW {
			minW = cube.W
		} else if cube.W > maxW {
			maxW = cube.W
		}
	}
	return minX, maxX, minY, maxY, minZ, maxZ, minW, maxW
}

func AnyCube(cubes map[Cube]bool) Cube {
	for cube := range cubes {
		return cube
	}
	return Cube{}
}
