package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/2018/common"
)

type statement struct {
	Input        []int
	OpNumber     int
	OpParameters []int
	Output       []int
	Index        int
}

var opcodeMap = map[string]func([]int, int, int, int) []int{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"aqrr": eqrr,
}

var opcodes = []func([]int, int, int, int) []int{
	addr,
	addi,
	mulr,
	muli,
	banr,
	bani,
	borr,
	bori,
	setr,
	seti,
	gtir,
	gtri,
	gtrr,
	eqir,
	eqri,
	eqrr,
}

func addr(registers []int, a, b, c int) []int {
	valueA := registers[a]
	valueB := registers[b]
	registers[c] = valueA + valueB
	return registers
}

func addi(registers []int, a, valueB, c int) []int {
	valueA := registers[a]
	registers[c] = valueA + valueB
	return registers
}

func mulr(registers []int, a, b, c int) []int {
	valueA := registers[a]
	valueB := registers[b]
	registers[c] = valueA * valueB
	return registers
}

func muli(registers []int, a, valueB, c int) []int {
	valueA := registers[a]
	registers[c] = valueA * valueB
	return registers
}

func banr(registers []int, a, b, c int) []int {
	valueA := registers[a]
	valueB := registers[b]
	registers[c] = valueA & valueB
	return registers
}

func bani(registers []int, a, valueB, c int) []int {
	valueA := registers[a]
	registers[c] = valueA & valueB
	return registers
}

func borr(registers []int, a, b, c int) []int {
	valueA := registers[a]
	valueB := registers[b]
	registers[c] = valueA | valueB
	return registers
}

func bori(registers []int, a, valueB, c int) []int {
	valueA := registers[a]
	registers[c] = valueA | valueB
	return registers
}

func setr(registers []int, a, _, c int) []int {
	registers[c] = registers[a]
	return registers
}

func seti(registers []int, valueA, _, c int) []int {
	registers[c] = valueA
	return registers
}

func gtir(registers []int, valueA, b, c int) []int {
	if valueA > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func gtri(registers []int, a, valueB, c int) []int {
	if registers[a] > valueB {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func gtrr(registers []int, a, b, c int) []int {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func eqir(registers []int, valueA, b, c int) []int {
	if valueA == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func eqri(registers []int, a, valueB, c int) []int {
	if registers[a] == valueB {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func eqrr(registers []int, a, b, c int) []int {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

func toInt(inputs []string) []int {
	var ints []int
	for _, input := range inputs {
		i, _ := strconv.Atoi(input)
		ints = append(ints, i)
	}
	return ints
}

func main() {
	results := common.ReadAllLines("./day16/input.txt", `(?P<data>.*)`)

	var statements []statement
	for i := 0; i < len(results)-3; i += 4 {
		input := results[i]["data"]
		input = strings.Replace(input, "Before: [", "", 1)
		input = strings.Replace(input, "]", "", 1)

		operation := results[i+1]["data"]
		vals := toInt(strings.Split(operation, " "))
		opNumber, params := vals[0], vals[1:]

		output := results[i+2]["data"]
		output = strings.Replace(output, "After:  [", "", 1)
		output = strings.Replace(output, "]", "", 1)

		statements = append(statements, statement{
			Input:        toInt(strings.Split(input, ", ")),
			OpNumber:     opNumber,
			OpParameters: params,
			Output:       toInt(strings.Split(output, ", ")),
			Index:        i,
		})
	}

	do(statements)
	//fmt.Println(result)
}

func isEqual(a []int, b []int) bool {
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func do(statements []statement) map[int][]string {
	opNumberMap := map[int][]string{}
	ambiguous := 0
	for _, statement := range statements {
		ops := []int{}
		for _, opFunc := range opcodes {
			registers := make([]int, 4)
			copy(registers, statement.Input)

			registers = opFunc(registers, statement.OpParameters[0], statement.OpParameters[1], statement.OpParameters[2])

			if isEqual(registers, statement.Output) {
				ops = append(ops, statement.OpNumber)
			}
		}
		if len(ops) >= 3 {
			ambiguous += 1
		}
		//opNumberMap[statement.OpNumber] = filterDuplicates(append(opNumberMap[statement.OpNumber], ops...))
	}
	fmt.Println(ambiguous)
	for opNumber, names := range opNumberMap {
		opNumberMap[opNumber] = filterDuplicates(names)
	}
	return opNumberMap
}

func filterDuplicates(a []string) []string {
	b := map[string]bool{}
	for _, c := range a {
		b[c] = true
	}
	var names []string
	for d := range b {
		names = append(names, d)
	}
	return names
}
