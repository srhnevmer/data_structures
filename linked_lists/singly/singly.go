package singly

type node struct {
	next  *node
	value int
}

type list struct {
	head *node
	size int
}

func (l *list) add(val int) {
	l.size++
	n := &node{nil, val}

	switch l.head {
	case nil:
		l.head = n
	default:
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
}
