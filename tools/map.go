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

func MapHasKeys(aMap interface{}, keySlice interface{}) bool {
	mapKeys := reflect.ValueOf(aMap).MapKeys()
	keys := reflect.ValueOf(keySlice)

	for i := 0; i < keys.Len(); i++ {
		var found bool
		for _, mapKey := range mapKeys {
			if mapKey.Interface() == keys.Index(i).Interface() {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func MapIntValues(i interface{}) []int {
	var ints []int

	iter := reflect.ValueOf(i).MapRange()
	for iter.Next() {
		ints = append(ints, int(iter.Value().Int()))
	}
	return ints
}
