package main

import (
	"strconv"
	"strings"

	"github.com/ldej/advent-of-code/2018/common"
	"fmt"
)

type operation struct {
	OpNumber     int
	OpParameters []int
}

type statement struct {
	Input        []int
	Operation operation
	Output       []int
	Index        int
}

var opcodeMap = map[string]func([]int, int, int, int) []int {
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

		oper := results[i+1]["data"]
		vals := toInt(strings.Split(oper, " "))
		opNumber, params := vals[0], vals[1:]

		output := results[i+2]["data"]
		output = strings.Replace(output, "After:  [", "", 1)
		output = strings.Replace(output, "]", "", 1)

		statements = append(statements, statement{
			Input:        toInt(strings.Split(input, ", ")),
			Operation: operation{
				OpNumber:     opNumber,
				OpParameters: params,
			},
			Output:       toInt(strings.Split(output, ", ")),
			Index:        i,
		})
	}

	var program []operation
	results = common.ReadAllLines("./day16/input2.txt", `(?P<data>.*)`)
	for i := 0; i < len(results); i += 1 {
		oper := results[i]["data"]
		if oper == "" {
			continue
		}
		vals := toInt(strings.Split(oper, " "))
		program = append(program, operation{
			OpNumber: vals[0],
			OpParameters: vals[1:],
		})
	}

	res := do(statements, program)
	fmt.Println(res)
}

func isEqual(a []int, b []int) bool {
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func do(statements []statement, program []operation) int {
	opNumberMap := map[int][]string{}

	ambiguous := 0

	for _, statement := range statements {
		ops := []string{}
		for opName, opFunc := range opcodeMap {
			registers := make([]int, 4)
			copy(registers, statement.Input)

			registers = opFunc(registers, statement.Operation.OpParameters[0], statement.Operation.OpParameters[1], statement.Operation.OpParameters[2])

			if isEqual(registers, statement.Output) {
				ops = append(ops, opName)
			}
		}

		if len(ops) >= 3 {
			ambiguous += 1
		}
		opNumberMap[statement.Operation.OpNumber] = filterDuplicates(append(opNumberMap[statement.Operation.OpNumber], ops...))
	}

	for opNumber, names := range opNumberMap {
		opNumberMap[opNumber] = filterDuplicates(names)
	}

	finalOpNumberMap := findOpNumberMap(opNumberMap, map[int]string{})

	registers := execute(finalOpNumberMap, program)

	return registers[0]
}

func execute(onm map[int]string, program []operation) []int {
	registers := []int{0, 0, 0, 0}

	for _, op := range program {
		opFunc := opcodeMap[onm[op.OpNumber]]
		registers = opFunc(registers, op.OpParameters[0], op.OpParameters[1], op.OpParameters[2])
	}
	return registers
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

func findOpNumberMap(onm map[int][]string, fonm map[int]string) map[int]string{
	allEmpty := true
	for number, names := range onm {
		if len(names) > 0 {
			allEmpty = false
		}

		if len(names) == 1 {
			fonm[number] = names[0]
			for number, n := range onm {
				onm[number] = filterString(n, names[0])
			}
		}
	}
	if allEmpty {
		return fonm
	}
	return findOpNumberMap(onm, fonm)
}

func filterString(a []string, b string) []string {
	var result []string
	for _, c := range a {
		if c != b {
			result = append(result, c)
		}
	}
	return result
}
