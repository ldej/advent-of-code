package main

import (
	"fmt"
	"sort"

	"github.com/ldej/advent-of-code/2018/common"
)

type direction byte

const (
	Left     direction = 0
	Straight direction = 1
	Right    direction = 2
)

var NextCrossingMove = map[direction]direction{
	Left:     Straight,
	Straight: Right,
	Right:    Left,
}

type cars []car

type car struct {
	X            int
	Y            int
	Orientation  byte
	NextCrossing direction
	ID           int
}

func main() {
	results := common.ReadAllLines("./day13/input.txt", `(?P<data>.*)`)

	grid := []string{}
	for _, result := range results {
		if result["data"] != "" {
			grid = append(grid, result["data"])
		}
	}
	var c cars
	c, grid = findCars(grid)

	for {
		//printCarsOnGrid(c, grid)
		c = moveCars(c, grid)
		if len(c) == 1 {
			fmt.Printf("Last car at %d,%d", c[0].X, c[0].Y)
			return
		}
	}
}

func printCarsOnGrid(c cars, grid []string) {
	fmt.Printf("\033[0;0H")
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if ca, _, found := getCarOnPosition(c, i, j); found {
				fmt.Print(string(ca.Orientation))
			} else {
				fmt.Print(string(grid[j][i]))
			}
		}
		fmt.Println()
	}
	fmt.Printf("\nCars: %2d\n", len(c))
}

func moveCars(c cars, grid []string) cars {
	for idx := 0; idx < len(c); idx++ {
		car := c[idx]
		switch car.Orientation {
		case '>':
			if _, index, found := getCarOnPosition(c, car.X+1, car.Y); found {
				fmt.Printf("Removing car %d\n", c[idx].ID)
				fmt.Printf("Removing car %d\n", c[index].ID)
				if idx > index {
					c = append(c[:idx], c[idx+1:]...)
					c = append(c[:index], c[index+1:]...)
					idx -= 2
				} else {
					c = append(c[:index], c[index+1:]...)
					c = append(c[:idx], c[idx+1:]...)
					idx -= 1
				}
			} else {
				switch grid[car.Y][car.X+1] {
				case '/':
					c[idx].Orientation = '^'
				case '\\':
					c[idx].Orientation = 'v'
				case '-':
				case '+':
					switch c[idx].NextCrossing {
					case Left:
						c[idx].Orientation = '^'
					case Straight:
						c[idx].Orientation = '>'
					case Right:
						c[idx].Orientation = 'v'
					}
					c[idx].NextCrossing = NextCrossingMove[c[idx].NextCrossing]
				default:
					fmt.Println("Help 1", grid[car.Y][car.X+1])
				}
				c[idx].X += 1
			}

		case '<':
			if _, index, found := getCarOnPosition(c, car.X-1, car.Y); found {
				fmt.Printf("Removing car %d\n", c[idx].ID)
				fmt.Printf("Removing car %d\n", c[index].ID)
				if idx > index {
					c = append(c[:idx], c[idx+1:]...)
					c = append(c[:index], c[index+1:]...)
					idx -= 2
				} else {
					c = append(c[:index], c[index+1:]...)
					c = append(c[:idx], c[idx+1:]...)
					idx -= 1
				}
			} else {
				switch grid[car.Y][car.X-1] {
				case '/':
					c[idx].Orientation = 'v'
				case '\\':
					c[idx].Orientation = '^'
				case '-':
				case '+':
					switch c[idx].NextCrossing {
					case Left:
						c[idx].Orientation = 'v'
					case Straight:
						c[idx].Orientation = '<'
					case Right:
						c[idx].Orientation = '^'
					}
					c[idx].NextCrossing = NextCrossingMove[c[idx].NextCrossing]
				default:
					fmt.Println("Help 2", grid[car.Y][car.X-1])
				}
				c[idx].X -= 1
			}
		case '^':
			if _, index, found := getCarOnPosition(c, car.X, car.Y-1); found {
				fmt.Printf("Removing car %d\n", c[idx].ID)
				fmt.Printf("Removing car %d\n", c[index].ID)
				if idx > index {
					c = append(c[:idx], c[idx+1:]...)
					c = append(c[:index], c[index+1:]...)
					idx -= 2
				} else {
					c = append(c[:index], c[index+1:]...)
					c = append(c[:idx], c[idx+1:]...)
					idx -= 1
				}
			} else {
				switch grid[car.Y-1][car.X] {
				case '/':
					c[idx].Orientation = '>'
				case '\\':
					c[idx].Orientation = '<'
				case '|':
				case '+':
					switch c[idx].NextCrossing {
					case Left:
						c[idx].Orientation = '<'
					case Straight:
						c[idx].Orientation = '^'
					case Right:
						c[idx].Orientation = '>'
					}
					c[idx].NextCrossing = NextCrossingMove[c[idx].NextCrossing]
				default:
					fmt.Println("Help 3", grid[car.Y-1][car.X])
				}
				c[idx].Y -= 1
			}

		case 'v':
			if _, index, found := getCarOnPosition(c, car.X, car.Y+1); found {
				fmt.Printf("Removing car %d\n", c[idx].ID)
				fmt.Printf("Removing car %d\n", c[index].ID)
				if idx > index {
					c = append(c[:idx], c[idx+1:]...)
					c = append(c[:index], c[index+1:]...)
					idx -= 2
				} else {
					c = append(c[:index], c[index+1:]...)
					c = append(c[:idx], c[idx+1:]...)
					idx -= 1
				}
			} else {
				switch grid[car.Y+1][car.X] {
				case '/':
					c[idx].Orientation = '<'
				case '\\':
					c[idx].Orientation = '>'
				case '|':
				case '+':
					switch c[idx].NextCrossing {
					case Left:
						c[idx].Orientation = '>'
					case Straight:
						c[idx].Orientation = 'v'
					case Right:
						c[idx].Orientation = '<'
					}
					c[idx].NextCrossing = NextCrossingMove[c[idx].NextCrossing]
				default:
					fmt.Println("Help 4", grid[car.Y+1][car.X])
				}
				c[idx].Y += 1
			}

		default:
			fmt.Println("Help 5", car.Orientation)
		}
	}
	c = sortCars(c)
	return c
}

func getCarOnPosition(c cars, x int, y int) (car, int, bool) {
	for idx, ca := range c {
		if ca.X == x && ca.Y == y {
			return ca, idx, true
		}
	}
	return car{}, 0, false
}

func findCars(grid []string) (cars, []string) {
	c := cars{}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == '>' || grid[j][i] == '<' || grid[j][i] == '^' || grid[j][i] == 'v' {
				c = append(c, car{
					X:            i,
					Y:            j,
					Orientation:  grid[j][i],
					NextCrossing: Left,
				})
				grid[j] = replaceCar(grid[j], i, grid[j][i])
			}
		}
	}
	c = sortCars(c)
	for idx := range c {
		c[idx].ID = idx + 1
	}
	return c, grid
}

func sortCars(c cars) cars {
	sort.Slice(c, func(i, j int) bool {
		if c[i].Y == c[j].Y {
			return c[i].X < c[j].X
		}
		return c[i].Y < c[j].Y
	})
	return c
}

func replaceCar(row string, index int, c byte) string {
	out := []rune(row)
	switch c {
	case '>', '<':
		out[index] = '-'
	case '^', 'v':
		out[index] = '|'
	default:
		panic(fmt.Sprintf("not a car: %c\n", c))
	}
	return string(out)
}
