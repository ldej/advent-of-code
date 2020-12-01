package main

import (
	"fmt"
	"github.com/ldej/advent-of-code-2018/common"
	"strconv"
	"strings"
)

func main() {
	results := common.ReadAllLines("./day4/input.txt", `^\[(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2}) (?P<hour>\d{2}):(?P<minutes>\d{2})\] (?P<log>.*)`)

	minutes := map[int]map[int]int{}

	activeGuard := 0
	fallsAsleep := 0
	for _, resultMap := range results {
		if resultMap["log"] == "falls asleep" {
			fallsAsleep, _ = strconv.Atoi(resultMap["minutes"])
		} else if resultMap["log"] == "wakes up" {
			wakesUp, _ := strconv.Atoi(resultMap["minutes"])
			for i := fallsAsleep; i < wakesUp; i++ {
				if min, ok := minutes[activeGuard]; ok {
					if _, minOk := min[i]; minOk {
						minutes[activeGuard][i] += 1
					} else {
						minutes[activeGuard][i] = 1
					}
				} else {
					minutes[activeGuard] = map[int]int{i: 1}
				}
			}
		} else {
			// Guard starts
			res := strings.Split(resultMap["log"], " ")
			if len(res) > 1 {
				activeGuard, _ = strconv.Atoi(res[1][1:])
			}
		}
	}

	guardSleepsMost := 0
	currentMost := 0
	for guard, mins := range minutes {
		currentTotal := 0
		for _, min := range mins {
			currentTotal += min
		}
		if currentTotal > currentMost {
			guardSleepsMost = guard
			currentMost = currentTotal
		}
	}

	mostMinute := 0
	mostMinuteValue := 0
	for key, minute := range minutes[guardSleepsMost] {
		if minute > mostMinuteValue {
			mostMinute = key
			mostMinuteValue = minute
		}
	}

	fmt.Println(guardSleepsMost * mostMinute)
}
