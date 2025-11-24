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
			prev, curr = curr, curr.next
		}
		n.next, prev.next = curr, n
	}
	l.size++
}

func (l *list) delete(idx uint) {
	if l.size == 0 || l.size <= idx {
		return
	}

	switch {
	case idx == 0 && l.size == 1:
		l.head = nil
	case idx == 0:
		target := l.head
		l.head = target.next
		target = nil
	default:
		var prev *node
		curr := l.head
		for range idx {
			prev, curr = curr, curr.next
		}
		prev.next = curr.next
		curr = nil
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
	var prev *node
	curr := l.head
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}
