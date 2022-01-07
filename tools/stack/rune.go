package stack

type RuneStack []rune

func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *RuneStack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	var result rune
	last := len(*s) - 1
	*s, result = (*s)[:last], (*s)[last]
	return result, true
}

func (s *RuneStack) Push(a rune) {
	*s = append(*s, a)
}

func (s *RuneStack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	last := len(*s) - 1
	return (*s)[last], true
}
