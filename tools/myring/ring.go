package myring

import (
	"container/ring"
	"fmt"
	"strings"
)

func NewIntRing(ints []int) *ring.Ring {
	r := ring.New(len(ints))
	for _, i := range ints {
		r.Value = i
		r = r.Next()
	}
	return r
}

func String(r *ring.Ring) string {
	var s strings.Builder
	for i := 0; i < r.Len(); i++ {
		s.WriteString(fmt.Sprintf("%-v ", r.Value))
		r = r.Next()
	}
	return s.String()
}

func Contains(r *ring.Ring, a int) bool {
	for i := 0; i < r.Len(); i++ {
		if r.Value.(int) == a {
			return true
		}
		r = r.Next()
	}
	return false
}
