package main

import (
	"container/ring"
	"fmt"
)

func main() {
	if calculate(9, 25) != 32 {
		fmt.Println("Error 9")
		return
	}
	if calculate(10, 1618) != 8317 {
		fmt.Println("Error 10")
		return
	}
	if calculate(13, 7999) != 146373 {
		fmt.Println("Error 13")
		return
	}
	if calculate(17, 1104) != 2764 {
		fmt.Println("Error 17")
		return
	}
	if calculate(21, 6111) != 54718 {
		fmt.Println("Error 21")
		return
	}
	if calculate(30, 5807) != 37305 {
		fmt.Println("Error 30")
		return
	}
	fmt.Println(calculate(418, 70769))
}

func calculate(nrOfPlayers int, lastmarble int) int {

	marbles := ring.New(1)
	marbles.Value = 0

	players := make([]int, nrOfPlayers)
	currentMarble := 1

	for i := 0; i <= lastmarble; i++ {
		idx := i % nrOfPlayers
		if currentMarble % 23 == 0 {
			players[idx] += currentMarble
			// Move back 9
			marbles = marbles.Move(-9)
			// Unlink and add to score
			players[idx] += marbles.Unlink(1).Value.(int)
			marbles = marbles.Move(2)
		} else {
			marblesNew := ring.New(1)
			marblesNew.Value = currentMarble
			marbles = marbles.Link(marblesNew)
		}

		currentMarble += 1
	}

	highscore := 0
	for idx := range players {
		if players[idx] > highscore {
			highscore = players[idx]
		}
	}

	return highscore
}
