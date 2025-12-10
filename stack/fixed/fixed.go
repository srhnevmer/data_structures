package fixed

const max = 5

type stack struct {
	size      int
	container [max]int
}

func new() stack {
	return stack{size: -1}
}

func (s *stack) push(value int) {
	if max-1 == s.size {
		return
	}
	s.size++
	s.container[s.size] = value
}

func (s *stack) pop() {
	if s.size == -1 {
		return
	}
	s.container[s.size] = 0
	s.size--
}

func (s *stack) peek() (int, bool) {
	if s.size == -1 {
		return 0, false
	}
	return s.container[s.size], true
}

func (s *stack) isEmpty() bool {
	return s.size == -1
}

func (s *stack) isFull() bool {
	return s.size == max-1
}

func (s *stack) getSize() int {
	if s.size == -1 {
		return s.size
	}
	return s.size + 1
}
