package fixed

const max = 5

type queue struct {
	head, tail int
	container  [max]int
}

func new() queue {
	return queue{head: -1, tail: -1}
}

func (q *queue) enqueue(value int) {
	if q.tail == max-1 {
		return
	}

	if q.head == -1 {
		q.head++
	}
	q.tail++
	q.container[q.tail] = value
}

func (q *queue) dequeue() {
	switch q.head {
	case -1:
		return
	case q.tail:
		q.head, q.tail = -1, -1
	default:
		q.head++
	}
}

func (q *queue) peek() int {
	switch q.head {
	case -1:
		return 0
	default:
		return q.container[q.head]
	}
}

func (q *queue) isEmpty() bool {
	return q.head == -1
}

func (q *queue) isFull() bool {
	return q.tail == max-1
}
