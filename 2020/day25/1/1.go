package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Part 1")

	result := run(5764801, 17807724)
	fmt.Println("Example:", result)

	result = run(1327981, 2822615)
	fmt.Println("Result:", result)
}

func run(card int, door int) int {
	var cardLoopSize = getLoopSize(card, 7)
	var doorLoopSize = getLoopSize(door, 7)

	var cardEncryptionKey = transformSubject(card, doorLoopSize)
	var doorEncryptionKey = transformSubject(door, cardLoopSize)

	if cardEncryptionKey != doorEncryptionKey {
		log.Fatal("Keys should be equal")
	}
	return cardEncryptionKey
}

func getLoopSize(key int, subject int) int {
	var loopSize = 1
	var temp = 1
	for {
		temp = (temp * subject) % 20201227
		if temp == key {
			return loopSize
		}
		loopSize++
	}
}

func transformSubject(key int, loopSize int) int {
	var temp = 1
	for i := 0; i < loopSize; i++ {
		temp = (temp * key) % 20201227
	}
	return temp
}
