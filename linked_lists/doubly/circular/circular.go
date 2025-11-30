package circular

type node struct {
	prev, next *node
	value      int
}

type list struct {
	head, tail *node
	size       uint
}

func (l *list) insert(index uint, value int) {
	if l.size < index {
		return
	}

	n := &node{nil, nil, value}
	switch {
	case l.size == 0:
		n.prev, n.next = n, n
		l.head, l.tail = n, n
	case index == 0:
		n.prev, l.tail.next = l.tail, n
		n.next, l.head.prev = l.head, n
		l.head = n
	default:
		var prev *node
		curr := l.head
		for range index {
			prev, curr = curr, curr.next
		}

		switch index {
		case l.size:
			prev.next, n.prev = n, prev
			l.head.prev, n.next = n, l.head
			l.tail = n
		default:
			prev.next, n.prev = n, prev
			curr.prev, n.next = n, curr
		}
	}
	l.size++
}
