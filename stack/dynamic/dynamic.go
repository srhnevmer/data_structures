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
