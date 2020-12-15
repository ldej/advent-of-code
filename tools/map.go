package tools

import (
	"reflect"
)

func MapAll(m map[string]bool, b bool) bool {
	for _, v := range m {
		if v != b {
			return false
		}
	}
	return true
}

func MapSumValues(i interface{}) int {
	var sum int64

	iter := reflect.ValueOf(i).MapRange()
	for iter.Next() {
		sum += iter.Value().Int()
	}
	return int(sum)
}
