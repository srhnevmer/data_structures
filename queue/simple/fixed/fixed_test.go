package fixed

import "testing"

type expected struct {
	head, tail int
	container  [max]int
}

func TestEnqueue(t *testing.T) {
	testCases := []struct {
		name string
		queue
		value int
		expected
	}{
		{
			name:     "Enqueue a value in an empty queue",
			queue:    initQueue(),
			value:    10,
			expected: expected{0, 0, [max]int{10}},
		},
		{
			name:     "Enqueue a value in a queue of size 1",
			queue:    getFilledQueue([]int{10}),
			value:    20,
			expected: expected{0, 1, [max]int{10, 20}},
		},
		{
			name:     "Enqueue a value in a queue of size 2",
			queue:    getFilledQueue([]int{10, 20}),
			value:    30,
			expected: expected{0, 2, [max]int{10, 20, 30}},
		},
		{
			name:     "Enqueue a value in a queue of size 3",
			queue:    getFilledQueue([]int{10, 20, 30}),
			value:    40,
			expected: expected{0, 3, [max]int{10, 20, 30, 40}},
		},
		{
			name:     "Enqueue a value in a queue of size 4",
			queue:    getFilledQueue([]int{10, 20, 30, 40}),
			value:    50,
			expected: expected{0, 4, [max]int{10, 20, 30, 40, 50}},
		},
		{
			name:     "Attempt to enqueue a value in a filled queue",
			queue:    getFilledQueue([]int{10, 20, 30, 40, 50}),
			value:    60,
			expected: expected{0, 4, [max]int{10, 20, 30, 40, 50}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.queue.enqueue(tc.value)
			assertIndex(t, "head", tc.expected.head, tc.queue.head)
			assertIndex(t, "tail", tc.expected.tail, tc.queue.tail)
			assertValues(t, tc.expected.container, tc.queue.container)
		})
	}
}

func initQueue() queue {
	return new()
}

func getFilledQueue(values []int) queue {
	q := new()
	for _, v := range values {
		if q.head == -1 {
			q.head++
		}
		q.tail++
		q.container[q.tail] = v
	}
	return q
}

func assertIndex(t testing.TB, kind string, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("Expected index[%s]: %d got: %d", kind, want, got)
	}
}

func assertValues(t testing.TB, want, got [max]int) {
	t.Helper()
	if want != got {
		t.Errorf("Expected values: %v got: %v", want, got)
	}
}
