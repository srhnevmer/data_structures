package dynamic

type node struct {
	next  *node
	value int
}

type stack struct {
	size uint
	head *node
}

func (s *stack) push(value int) {
	switch s.size {
	case 0:
		s.head = &node{nil, value}
	default:
		s.head = &node{s.head, value}
	}
	s.size++
}

func (s *stack) pop() {
	if s.size == 0 {
		return
	}
	target := s.head
	s.head = target.next
	target.next = nil
	s.size--
}

func (s *stack) peek() (int, bool) {
	switch s.size {
	case 0:
		return 0, false
	default:
		return s.head.value, true
	}
}
