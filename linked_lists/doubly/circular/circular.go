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

func (l *list) delete(index uint) {
	if l.size <= index {
		return
	}

	switch {
	case l.size == 1:
		l.head.prev, l.head.next = nil, nil
		l.head, l.tail = nil, nil
	case index == 0:
		target := l.head
		l.head = target.next
		l.head.prev, l.tail.next = l.tail, l.head
		target.prev, target.next = nil, nil
	default:
		curr := l.head
		for range index {
			curr = curr.next
		}

		switch index {
		case l.size - 1:
			curr.prev.next, l.head.prev = l.head, curr.prev
			l.tail = curr.prev
		default:
			curr.prev.next, curr.next.prev = curr.next, curr.prev
		}
		curr.prev, curr.next = nil, nil
	}
	l.size--
}
