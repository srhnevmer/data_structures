package circular

type node struct {
	next  *node
	value int
}

type list struct {
	head, tail *node
	size       uint
}

func (l *list) insert(idx uint, val int) {
	if l.size < idx {
		return
	}

	l.size++
	n := &node{nil, val}
	switch {
	case l.head == nil:
		n.next = n
		l.head, l.tail = n, n
	case idx == 0:
		n.next = l.head
		l.head = n
		l.tail.next = l.head
	case idx == l.size-1:
		n.next = l.head
		l.tail.next = n
		l.tail = n
	default:
		var prev *node
		curr := l.head
		for range idx {
			prev = curr
			curr = curr.next
		}
		n.next, prev.next = prev.next, n
	}
}

func (l *list) delete(idx uint) {
	if l.size <= idx || l.head == nil {
		return
	}

	switch {
	case idx == 0 && l.size == 1:
		l.head, l.tail = nil, nil
	case idx == 0:
		target := l.head
		l.head = target.next
		l.tail.next = l.head
		target.next = nil
	case idx == l.size-1:
		var prev *node
		curr := l.head
		for range idx {
			prev = curr
			curr = curr.next
		}
		prev.next = l.head
		l.tail = prev
	default:
		var prev *node
		curr := l.head
		for range idx {
			prev = curr
			curr = curr.next
		}
		prev.next = curr.next
		curr.next = nil
	}
	l.size--
}
