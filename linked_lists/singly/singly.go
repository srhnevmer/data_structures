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
		curr := l.head
		for prevIdx := 0; prevIdx < int(idx)-1; prevIdx++ {
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

func (l *list) delete(idx uint) {
	if l.size == 0 || l.size <= idx {
		return
	}

	l.size--

	switch {
	case l.size == 1:
		l.head = nil
	case idx == 0:
		target := l.head
		l.head = target.next
		target.next = nil
	default:
		curr := l.head
		for i := uint(0); i < idx-1; i++ {
			curr = curr.next
		}

		target := curr.next
		if target == nil {
			curr.next = nil
			return
		}

		curr.next = target.next
		target.next = nil
	}
}

func (l *list) traverse() []int {
	r := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		r = append(r, curr.value)
	}
	return r
}
