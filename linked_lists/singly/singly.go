package singly

type node struct {
	next  *node
	value int
}

type list struct {
	head *node
	size uint
}

func (l *list) insert(idx uint, val int) {
	if l.size < idx {
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
		var prev *node
		curr := l.head
		for range idx {
			prev = curr
			curr = curr.next
		}
		n.next = curr
		prev.next = n
	}
}

func (l *list) delete(idx uint) {
	if l.size == 0 || l.size <= idx {
		return
	}

	l.size--
	switch {
	case l.size == 1:
		l.head = nil
	case idx == 0:
		t := l.head
		l.head = t.next
		t.next = nil
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
}

func (l *list) traverse() []int {
	r := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		r = append(r, curr.value)
	}
	return r
}
