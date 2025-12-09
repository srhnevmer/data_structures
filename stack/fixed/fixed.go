package fixed

const max = 5

type stack struct {
	size      int
	container [max]int
}

func (s *stack) peek() (int, bool) {
	if s.size == -1 {
		return 0, false
	}
	return s.container[s.size], true
}
