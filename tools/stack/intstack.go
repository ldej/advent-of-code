package stack

type IntStack []int

func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	var result int
	last := len(*s) - 1
	*s, result = (*s)[:last], (*s)[last]
	return result, true
}

func (s *IntStack) Push(a int) {
	*s = append(*s, a)
}
