package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadBytes(location string) []byte {
	bytes, err := ioutil.ReadFile(location)
	if err != nil {
		fmt.Print(err)
	}
	return bytes
}

func ReadLines(location string) *bufio.Scanner {
	file, err := os.Open(location)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}
