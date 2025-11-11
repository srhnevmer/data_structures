package singly

type node struct {
	next  *node
	value int
}

type list struct {
	head *node
	size int
}

func (l *list) traversal() []int {
	if l.head == nil {
		return make([]int, 0)
	}

	r := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		r = append(r, curr.value)
	}

	return r
}
