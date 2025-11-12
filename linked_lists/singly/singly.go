package singly

type node struct {
	next  *node
	value int
}

type list struct {
	head *node
	size int
}

func (l *list) insert(idx, val int) {
	if idx < 0 || l.size < idx {
		return
	}

	l.size++
	n := &node{nil, val}

	switch {
	case l.head == nil:
		l.head = n
	case idx == 0:
		n.next = l.head
		l.head = n
	default:
		curr := l.head
		for prevIdx := 0; prevIdx < idx-1; prevIdx++ {
			curr = curr.next
		}

		if curr.next == nil {
			curr.next = n
			return
		}

		n.next = curr.next
		curr.next = n
	}

}
