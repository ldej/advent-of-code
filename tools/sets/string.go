package sets

import (
	"fmt"
	"go/types"
)

type StringSet struct {
	items map[string]types.Nil
}

func NewStringSet(items ...string) *StringSet {
	set := new(StringSet)
	set.items = make(map[string]types.Nil)
	for _, item := range items {
		set.Add(item)
	}
	return set
}

func (set *StringSet) Add(items ...string) *StringSet {
	for _, item := range items {
		if item != "" {
			set.items[item] = types.Nil{}
		}
	}
	return set
}

func (set *StringSet) Remove(item string) *StringSet {
	if _, found := set.items[item]; found {
		delete(set.items, item)
	}
	return set
}

func (set *StringSet) Clear() {
	set.items = make(map[string]types.Nil)
}

func (set *StringSet) Copy() *StringSet {
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

func (set *StringSet) Intersect(setToIntersectWith *StringSet) *StringSet {
	resultSet := NewStringSet()

	for item := range set.items {
		if setToIntersectWith.Contains(item) {
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

func (set StringSet) Items() []string {
	var strings []string
	for key := range set.items {
		strings = append(strings, key)
	}
	return strings
}
