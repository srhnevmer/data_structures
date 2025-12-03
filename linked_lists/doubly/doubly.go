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

		switch index {
		case l.size:
			prev.next, n.prev = n, prev
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
		l.head.next, l.head.prev = nil, nil
		l.head = nil
	case index == 0:
		target := l.head
		l.head = target.next
		target.next.prev, target.next = nil, nil
	default:
		curr := l.head
		for range index {
			curr = curr.next
		}

		switch index {
		case l.size - 1:
			curr.prev.next = nil
		default:
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
		}
		curr.prev, curr.next = nil, nil
	}
	l.size--
}

func (l *list) search(target int) (uint, bool) {
	if l.size != 0 {
		for i, curr := uint(0), l.head; curr != nil; i, curr = i+1, curr.next {
			if curr.value == target {
				return i, true
			}
		}
	}
	return 0, false
}

func (l *list) reverse() {
	for curr := l.head; curr != nil; curr = curr.prev {
		curr.prev, curr.next = curr.next, curr.prev
		if curr.prev == nil {
			l.head = curr
		}
	}
}
