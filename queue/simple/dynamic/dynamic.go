package dynamic

type node struct {
	next  *node
	value int
}

type queue struct {
	head, tail *node
}

func (q *queue) enqueue(value int) {
	n := &node{nil, value}
	switch q.head {
	case nil:
		q.head, q.tail = n, n
	default:
		q.tail.next = n
		q.tail = n
	}
}
