package dynamic

import (
	"slices"
	"testing"
)

type expected struct {
	headValue int
	tailValue int
	values    []int
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
			queue:    queue{},
			value:    10,
			expected: expected{10, 10, []int{10}},
		},
		{
			name:     "Enqueue a value in a queue of size 1",
			queue:    getFilledQueue([]int{10}),
			value:    20,
			expected: expected{10, 20, []int{10, 20}},
		},
		{
			name:     "Enqueue a value in a queue of size 2",
			queue:    getFilledQueue([]int{10, 20}),
			value:    30,
			expected: expected{10, 30, []int{10, 20, 30}},
		},
		{
			name:     "Enqueue a value in a queue of size 3",
			queue:    getFilledQueue([]int{10, 20, 30}),
			value:    40,
			expected: expected{10, 40, []int{10, 20, 30, 40}},
		},
		{
			name:     "Enqueue a value in a queue of size 4",
			queue:    getFilledQueue([]int{10, 20, 30, 40}),
			value:    50,
			expected: expected{10, 50, []int{10, 20, 30, 40, 50}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.queue.enqueue(tc.value)
			assertValue(t, "Head", tc.expected.headValue, tc.queue.head.value)
			assertValue(t, "Tail", tc.expected.tailValue, tc.queue.tail.value)
			assertValues(t, tc.expected.values, getValues(tc.queue))
		})
	}
}

func getFilledQueue(values []int) queue {
	q := queue{}
	for i, v := range values {
		n := &node{nil, v}
		switch i {
		case 0:
			q.head, q.tail = n, n
		default:
			q.tail.next = n
			q.tail = n
		}
	}
	return q
}

func assertValue(t testing.TB, kind string, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("[%s] Expected value: %d got: %d", kind, want, got)
	}
}

func assertValues(t testing.TB, want, got []int) {
	t.Helper()
	if slices.Compare(want, got) != 0 {
		t.Errorf("Expected values: %v got: %v", want, got)
	}
}

func getValues(q queue) []int {
	max := 5
	v := make([]int, 0, max)
	for curr := q.head; curr != nil; curr = curr.next {
		v = append(v, curr.value)
	}
	return v
}
