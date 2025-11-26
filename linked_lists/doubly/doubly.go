package doubly

type node struct {
	prev, next *node
	value      int
}

type list struct {
	head *node
	size uint
}

func (l *list) insert(index uint, value int) {
	if l.size < index {
		return
	}

	n := &node{nil, nil, value}
	switch {
	case l.head == nil:
		l.head = n
	case index == 0:
		l.head.prev, n.next = n, l.head
		l.head = n
	default:
		var prev *node
		curr := l.head
		for range index {
			prev, curr = curr, curr.next
		}

		if index == l.size {
			prev.next, n.prev = n, prev
		} else {
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
	case index == 0 && l.size == 1:
		l.head = nil
	case index == 0:
		target := l.head
		target.next.prev = nil
		l.head = target.next
		target = nil
	default:
		curr := l.head
		for range index {
			curr = curr.next
		}

		if index == l.size-1 {
			curr.prev.next = nil
		} else {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
		}
		curr.prev, curr.next = nil, nil
	}
	l.size--
}
