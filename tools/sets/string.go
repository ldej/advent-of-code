package sets

import (
	"fmt"
	"go/types"
)

type StringSet struct {
	items map[string]types.Nil
}

func NewStringSet() StringSet {
	set := new(StringSet)
	set.items = make(map[string]types.Nil)
	return *set
}

func (set *StringSet) Add(item string) {
	if item != "" {
		set.items[item] = types.Nil{}
	}
}

func (set *StringSet) Remove(item string) {
	if _, found := set.items[item]; found {
		delete(set.items, item)
	}
}

func (set *StringSet) Clear() {
	set.items = make(map[string]types.Nil)
}

func (set *StringSet) Copy() StringSet {
	resultSet := NewStringSet()

	for item := range set.items {
		resultSet.Add(item)
	}

	return resultSet
}

func (set *StringSet) Contains(item string) bool {
	_, has := set.items[item]
	return has
}

func (set *StringSet) Len() int {
	return len(set.items)
}

func (set *StringSet) Max() string {
	max := ""

	for item := range set.items {
		if item > max {
			max = item
		}
	}

	return max
}

func (set *StringSet) Min() string {
	min := ""

	for item := range set.items {
		if item < min {
			min = item
		}
	}

	return min
}

func (set *StringSet) Intersect(setToIntersectWith StringSet) StringSet {
	resultSet := NewStringSet()

	for item := range set.items {
		if set.Contains(item) != setToIntersectWith.Contains(item) {
			resultSet.Add(item)
		}
	}

	return resultSet
}

func (set StringSet) String() string {
	str := "{ "
	for item := range set.items {
		str += fmt.Sprintf("'%s', ", item)
	}

	return str[:len(str)-2] + " }"
}