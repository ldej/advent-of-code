package stack

type StringStack []string

func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	var result string
	last := len(*s) - 1
	*s, result = (*s)[:last], (*s)[last]
	return result, true
}

func (s *StringStack) Push(a string) {
	*s = append(*s, a)
}

func (s *StringStack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	last := len(*s) - 1
	return (*s)[last], true
}
