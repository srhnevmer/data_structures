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

	n := &node{nil, val}
	switch {
	case l.head == nil:
		n.next = n
		l.head, l.tail = n, n
	case idx == 0:
		n.next = l.head
		l.tail.next, l.head = n, n
	case idx == l.size:
		n.next = l.head
		l.tail.next, l.tail = n, n
	default:
		var prev *node
		curr := l.head
		for range idx {
			prev = curr
			curr = curr.next
		}
		n.next, prev.next = prev.next, n
	}
	l.size++
}

func (l *list) delete(idx uint) {
	if l.head == nil || l.size <= idx {
		return
	}

	switch {
	case idx == 0 && l.size == 1:
		l.head, l.tail = nil, nil
	case idx == 0:
		target := l.head
		l.head = target.next
		l.tail.next = l.head
		target = nil
	default:
		var prev *node
		curr := l.head
		for range idx {
			prev = curr
			curr = curr.next
		}

		if idx == l.size-1 {
			l.tail = prev
			prev.next = l.head
			curr = nil
		} else {
			prev.next = curr.next
			curr = nil
		}
	}
	l.size--
}

func (l *list) search(target int) (uint, bool) {
	if l.size != 0 {
		for i, curr := uint(0), l.head; i < l.size; i, curr = i+1, curr.next {
			if curr.value == target {
				return i, true
			}
		}
	}
	return 0, false
}

func (l *list) reverse() {
	var prev *node
	curr := l.head
	for i := uint(0); i < l.size; i++ {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head.next = l.tail
	l.head, l.tail = l.tail, l.head
}
