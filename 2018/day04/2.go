package main

import (
	"github.com/ldej/advent-of-code-2018/common"
	"strconv"
	"fmt"
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
	guardSleepsMostMinute := 0
	guardSleepsMostMinuteValue := 0
	for guard, mins := range minutes {
		for minute, value := range mins {
			if value > guardSleepsMostMinuteValue {
				guardSleepsMost = guard
				guardSleepsMostMinute = minute
				guardSleepsMostMinuteValue = value
			}
		}
	}

	fmt.Println(guardSleepsMost * guardSleepsMostMinute)
}
