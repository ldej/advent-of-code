package main

import (
	"fmt"
	"sort"

	"github.com/ldej/advent-of-code/2018/common"
	"strings"
	"time"
)

type point struct {
	X int
	Y int
	Steps int
}

func (p point) NextTo(a point) bool {
	if p.X == a.X {
		if p.Y == a.Y-1 || p.Y == a.Y+1 {
			return true
		}
	} else if p.Y == a.Y {
		if p.X == a.X-1 || p.X == a.X+1 {
			return true
		}
	}
	return false
}

func (p point) ComesBefore(a point) bool {
	if p.Y == a.Y {
		return p.X < a.X
	}
	return p.Y < a.Y
}

type character struct {
	ID        int
	Type      byte
	Point     point
	Attack    int
	HitPoints int
}

type characters []character

func main() {
	results := common.ReadAllLines("./day15/input.txt", `(?P<data>.*)`)

	grid := []string{}
	for _, result := range results {
		if result["data"] != "" {
			grid = append(grid, result["data"])
		}
	}
	var c characters
	c, grid = findCharacters(grid)

	printGrid(grid, c)

	rounds := 0
	somebodyWon := false
	finishedRound := false
	for !somebodyWon {
		killed := false
		for idx := 0; idx < len(c); idx++ {
			char := c[idx]
			char, c = move(char, c, grid)
			c, killed = attack(char, c)
			printGrid(grid, c)
			time.Sleep(20 * time.Millisecond)

			if killed {
				idx -= 1
				killed = false
			}
			finishedRound = idx == len(c) - 1
			if somebodyWins(c) {
				somebodyWon = true
				break
			}
		}

		if finishedRound {
			c = sortCharacters(c)
			rounds += 1
			printGrid(grid, c)
			fmt.Println("Rounds: ", rounds)
		}
	}

	hp := calculateHP(c)
	fmt.Println(rounds, "*", hp, "=", rounds * hp)
	return
}

func calculateHP(c characters) int {
	total := 0
	for _, char := range c {
		total += char.HitPoints
	}
	return total
}

func attack(char character, c characters) (characters, bool) {
	enemies := findEnemies(char.Type, c)
	if enemies, foundAny := nextTo(char, enemies); foundAny {
		lowestHP := enemies[0]
		for _, enemy := range enemies {
			if enemy.HitPoints < lowestHP.HitPoints {
				lowestHP = enemy
			}
		}
		lowestHP.HitPoints -= char.Attack
		if lowestHP.HitPoints <= 0 {
			return removeCharacter(c, lowestHP), lowestHP.Point.ComesBefore(char.Point)
		}
		return replaceCharacter(c, lowestHP), false
	}
	return c, false
}

func move(char character, c characters, grid []string) (character, characters) {
	enemies := findEnemies(char.Type, c)
	if _, foundAny := nextTo(char, enemies); foundAny {
		return char, c
	}
	targets := findTargets(enemies, grid)
	//reachable, _ := findReachable(char.Point, c, targets, grid)
	p, found := findNearest(char.Point, targets, c ,grid)
	if !found {
		return char, c
	}
	return doImprovingMove(char, p, c, grid)
}

func doImprovingMove(char character, destination point, c characters, grid []string) (character, characters) {
	available := findAvailable(getAdjecent(char.Point), c, grid)
	for idx := range available {
		if available[idx].X == destination.X && available[idx].Y == destination.Y {
			char.Point = destination
			return char, replaceCharacter(c, char)
		}
	}
	sorted := sortedPoints(available)
	if nearest, found := findNearest(destination, sorted, c, grid); found {
		char.Point = nearest
		return char, replaceCharacter(c, char)
	}
	return char, c
}

func removeCharacter(c characters, char character) characters {
	for idx := range c {
		if c[idx].ID == char.ID {
			return append(c[:idx], c[idx+1:]...)
		}
	}
	return c
}

func replaceCharacter(c characters, char character) characters {
	for idx := range c {
		if c[idx].ID == char.ID {
			c[idx] = char
			return c
		}
	}
	return c
}

func findNearest(p point, points []point, c characters, grid []string) (point, bool) {
	reachable, distanceMap := findReachable(p, c, sortedPoints(points), grid)

	//printDistanceMap(distanceMap, grid, c)

	var nearest point
	minDistance := 100
	reachable = sortedPoints(reachable)
	for _, r := range reachable {
		if distanceMap[r.Y][r.X] < minDistance {
			nearest = r
			minDistance = distanceMap[r.Y][r.X]
		}
	}
	return nearest, len(reachable) > 0
}

func sortedPoints(points []point) []point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].Y == points[j].Y {
			return points[i].X < points[j].X
		}
		return points[i].Y < points[j].Y
	})
	return points
}

func findReachable(p point, c characters, points []point, grid []string) ([]point, [][]int) {
	var mapped [][]int
	for jdx := 0; jdx < len(grid); jdx++ {
		mapped = append(mapped, []int{})
		for idx := 0; idx < len(grid[jdx]); idx++ {
			mapped[jdx] = append(mapped[jdx], -1)
		}
	}

	var reachable []point
	toMap := findAvailable(getAdjecent(p), c, grid)
	mapped[p.Y][p.X] = 0
	for idx, a := range toMap {
		toMap[idx].Steps = 1
		mapped[a.Y][a.X] = toMap[idx].Steps
	}

	for len(toMap) > 0 {
		adjecent := findAvailable(getAdjecent(toMap[0]), c, grid)
		for _, adj := range adjecent {
			if d := mapped[adj.Y][adj.X]; d < 0 {
				adj.Steps = toMap[0].Steps + 1
				toMap = append(toMap, adj)
				mapped[adj.Y][adj.X] = adj.Steps
			}
		}
		for idx := 0; idx < len(points); idx++ {
			if points[idx].X == toMap[0].X && points[idx].Y == toMap[0].Y {
				reachable = append(reachable, points[idx])
				points = append(points[:idx], points[idx+1:]...)
				idx -= 1
			}
		}
		toMap = toMap[1:]
	}
	return reachable, mapped
}

func isAvailable(grid []string, c characters, x, y int) bool {
	if grid[y][x] != '.' {
		return false
	}
	for _, char := range c {
		if char.Point.X == x && char.Point.Y == y {
			return false
		}
	}
	return true
}

func findAvailable(points []point, c characters, grid []string) []point {
	var available []point
	for _, p := range points {
		if isAvailable(grid, c, p.X, p.Y) {
			available = append(available, p)
		}
	}
	return available
}

func getAdjecent(p point) []point {
	return []point{
		{
			X: p.X,
			Y: p.Y - 1,
		},
		{
			X: p.X,
			Y: p.Y + 1,
		},
		{
			X: p.X - 1,
			Y: p.Y,
		},
		{
			X: p.X + 1,
			Y: p.Y,
		},
	}
}

func findTargets(c characters, grid []string) []point {
	var targets []point
	for _, char := range c {
		if grid[char.Point.Y-1][char.Point.X] == '.' {
			targets = append(targets, point{
				X: char.Point.X,
				Y: char.Point.Y - 1,
			})
		}
		if grid[char.Point.Y][char.Point.X-1] == '.' {
			targets = append(targets, point{
				X: char.Point.X - 1,
				Y: char.Point.Y,
			})
		}
		if grid[char.Point.Y][char.Point.X+1] == '.' {
			targets = append(targets, point{
				X: char.Point.X + 1,
				Y: char.Point.Y,
			})
		}
		if grid[char.Point.Y+1][char.Point.X] == '.' {
			targets = append(targets, point{
				X: char.Point.X,
				Y: char.Point.Y + 1,
			})
		}
	}
	return targets
}

func findEnemies(typ byte, c characters) characters {
	var enemies characters
	for _, char := range c {
		if char.Type != typ {
			enemies = append(enemies, char)
		}
	}
	return enemies
}

func nextTo(me character, c characters) (characters, bool) {
	var n characters
	for _, char := range c {
		if char.Point.NextTo(me.Point) {
			n = append(n, char)
		}
	}
	return n, len(n) > 0
}

func findCharacters(grid []string) (characters, []string) {
	c := characters{}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == 'E' || grid[j][i] == 'G' {
				c = append(c, character{
					ID:   len(c) + 1,
					Type: grid[j][i],
					Point: point{
						X: i,
						Y: j,
					},
					Attack:    3,
					HitPoints: 200,
				})
				grid[j] = replaceLetter(grid[j], i)
			}
		}
	}
	c = sortCharacters(c)
	return c, grid
}

func replaceLetter(grid string, index int) string {
	out := []rune(grid)
	out[index] = '.'
	return string(out)
}

func sortCharacters(c characters) characters {
	sort.Slice(c, func(i, j int) bool {
		if c[i].Point.Y == c[j].Point.Y {
			return c[i].Point.X < c[j].Point.X
		}
		return c[i].Point.Y < c[j].Point.Y
	})
	return c
}

func somebodyWins(c characters) bool {
	first := c[0]
	for _, char := range c {
		if char.Type != first.Type {
			return false
		}
	}
	return true
}

func printGrid(grid []string, c characters) {
	fmt.Printf("\033[0;0H")
	for i := 0; i < len(grid)+1; i++ {
		fmt.Println("                                                                                                ")
	}
	fmt.Printf("\033[0;0H")
	for j := 0; j < len(grid); j++ {
		var scores characters
		for i := 0; i < len(grid[j]); i++ {
			if ca, found := getCharacterOnPosition(c, i, j); found {
				fmt.Print(string(ca.Type))
				scores = append(scores, ca)
			} else {
				fmt.Print(string(grid[j][i]))
			}
		}
		var s []string
		for _, a := range scores {
			s = append(s, fmt.Sprintf("%s(%d)", string(a.Type), a.HitPoints))
		}
		fmt.Printf(" %s\n", strings.Join(s, ", "))
	}
}

func printDistanceMap(distanceMap [][]int, grid []string, c characters) {
	fmt.Printf("\033[0;0H")
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if ca, found := getCharacterOnPosition(c, i, j); found {
				fmt.Print(string(ca.Type))
			} else {
				if grid[j][i] == '#' {
					fmt.Print(string(grid[j][i]))
				} else {
					if distanceMap[j][i] >=0 && distanceMap[j][i] < 10 {
						fmt.Print(distanceMap[j][i])
					} else {
						fmt.Print(string(grid[j][i]))
					}
				}
			}
		}
		fmt.Print("\n")
	}
}

func getCharacterOnPosition(c characters, i, j int) (character, bool) {
	for _, char := range c {
		if char.Point.X == i && char.Point.Y == j {
			return char, true
		}
	}
	return character{}, false
}
