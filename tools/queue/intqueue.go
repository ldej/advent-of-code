package queue

type IntQueue []int

func NewIntQueue(ints ...int) IntQueue {
	return ints
}

func (s *IntQueue) Copy() IntQueue {
	var newQueue = make(IntQueue, len(*s))
	copy(newQueue, *s)
	return newQueue
}

func (s *IntQueue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *IntQueue) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	var result int
	result, *s = (*s)[0], (*s)[1:]
	return result, true
}

func (s *IntQueue) Push(a ...int) {
	*s = append(*s, a...)
}
