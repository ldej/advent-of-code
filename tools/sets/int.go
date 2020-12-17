package sets

import (
	"fmt"
	"go/types"
)

type IntSet struct {
	items map[int]types.Nil
}

func NewIntSet() IntSet {
	set := new(IntSet)
	set.items = make(map[int]types.Nil)
	return *set
}

func (set *IntSet) Add(item int) {
	set.items[item] = types.Nil{}
}

func (set *IntSet) Remove(item int) {
	if _, found := set.items[item]; found {
		delete(set.items, item)
	}
}

func (set *IntSet) Clear() {
	set.items = make(map[int]types.Nil)
}

func (set *IntSet) Copy() IntSet {
	resultSet := NewIntSet()

	for item := range set.items {
		resultSet.Add(item)
	}

	return resultSet
}

func (set *IntSet) Contains(item int) bool {
	_, has := set.items[item]
	return has
}

func (set *IntSet) Len() int {
	return len(set.items)
}

func (set *IntSet) Max() int {
	max := 0

	for item := range set.items {
		if item > max {
			max = item
		}
	}

	return max
}

func (set *IntSet) Min() int {
	min := 0

	for item := range set.items {
		if item < min {
			min = item
		}
	}

	return min
}

func (set *IntSet) Intersect(setToIntersectWith IntSet) IntSet {
	resultSet := NewIntSet()

	for item := range set.items {
		if set.Contains(item) != setToIntersectWith.Contains(item) {
			resultSet.Add(item)
		}
	}

	return resultSet
}

func (set IntSet) String() string {
	str := "{ "
	for item := range set.items {
		str += fmt.Sprintf("'%d', ", item)
	}

	return str[:len(str)-2] + " }"
}
