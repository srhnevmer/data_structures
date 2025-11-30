package circular

import (
	"slices"
	"testing"
)

type expected struct {
	size       uint
	values     []int
	valuesTail []int
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		list
		index uint
		value int
		expected
	}{
		// 0
		{
			name:     "Attempt to insert a value into an empty list with an invalid index",
			list:     list{},
			index:    1,
			value:    50,
			expected: expected{0, []int{}, []int{}},
		},
		{
			name:     "Insert a value into an empty list at the 0 index",
			list:     list{},
			index:    0,
			value:    50,
			expected: expected{1, []int{50}, []int{50}},
		},
		// 1
		{
			name:     "Attempt to insert a value into a list of size 1 with an invalid index",
			list:     getFilledList([]int{10}),
			index:    2,
			value:    50,
			expected: expected{1, []int{10}, []int{10}},
		},
		{
			name:     "Insert a value into a list of size 1 at the 0 index",
			list:     getFilledList([]int{10}),
			index:    0,
			value:    50,
			expected: expected{2, []int{50, 10}, []int{10, 50}},
		},
		{
			name:     "Insert a value into a list of size 1 at the 1 index",
			list:     getFilledList([]int{10}),
			index:    1,
			value:    50,
			expected: expected{2, []int{10, 50}, []int{50, 10}},
		},
		// 2
		{
			name:     "Attempt to insert a value into a list of size 2 with an invalid index",
			list:     getFilledList([]int{10, 20}),
			index:    3,
			value:    50,
			expected: expected{2, []int{10, 20}, []int{20, 10}},
		},
		{
			name:     "Insert a value into a list of size 2 at the 0 index",
			list:     getFilledList([]int{10, 20}),
			index:    0,
			value:    50,
			expected: expected{3, []int{50, 10, 20}, []int{20, 10, 50}},
		},
		{
			name:     "Insert a value into a list of size 2 at the 1 index",
			list:     getFilledList([]int{10, 20}),
			index:    1,
			value:    50,
			expected: expected{3, []int{10, 50, 20}, []int{20, 50, 10}},
		},
		{
			name:     "Insert a value into a list of size 2 at the 2 index",
			list:     getFilledList([]int{10, 20}),
			index:    2,
			value:    50,
			expected: expected{3, []int{10, 20, 50}, []int{50, 20, 10}},
		},
		// 3
		{
			name:     "Attempt to insert a value into a list of size 3 with an invalid index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    4,
			value:    50,
			expected: expected{3, []int{10, 20, 30}, []int{30, 20, 10}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 0 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    0,
			value:    50,
			expected: expected{4, []int{50, 10, 20, 30}, []int{30, 20, 10, 50}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 1 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    1,
			value:    50,
			expected: expected{4, []int{10, 50, 20, 30}, []int{30, 20, 50, 10}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 2 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    2,
			value:    50,
			expected: expected{4, []int{10, 20, 50, 30}, []int{30, 50, 20, 10}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 3 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    3,
			value:    50,
			expected: expected{4, []int{10, 20, 30, 50}, []int{50, 30, 20, 10}},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.expected.size, tc.list.size)
			assertValues(t, tc.expected.values, getValues(tc.list))
			assertValues(t, tc.expected.valuesTail, getValuesTail(tc.list))
			if i != 0 && i != 2 && i != 5 && i != 9 {
				assertPointer(t, tc.list.head, tc.list.tail.next)
				assertPointer(t, tc.list.tail, tc.list.head.prev)
			}
		})
	}
}

func assertSize(t testing.TB, want, got uint) {
	t.Helper()
	if want != got {
		t.Fatalf("Expected size: %d got: %d", want, got)
	}
}

func assertValues(t testing.TB, want, got []int) {
	t.Helper()
	if slices.Compare(want, got) != 0 {
		t.Fatalf("Expected values: %d got: %d", want, got)
	}
}

func assertPointer(t testing.TB, want, got *node) {
	t.Helper()
	if want != got {
		t.Fatalf("Expected pointer: [ptr: %p val: %v] got: [ptr: %p val: %v]", want, *want, got, *got)
	}
}

func getValues(l list) []int {
	v := make([]int, 0, l.size)
	for i, curr := uint(0), l.head; i < l.size; i, curr = i+1, curr.next {
		v = append(v, curr.value)
	}
	return v
}

func getValuesTail(l list) []int {
	v := make([]int, 0, l.size)
	for i, curr := uint(0), l.tail; i < l.size; i, curr = i+1, curr.prev {
		v = append(v, curr.value)
	}
	return v
}

func getFilledList(values []int) list {
	l := list{size: uint(len(values))}
	for i := range values {
		n := &node{nil, nil, values[i]}
		switch l.head {
		case nil:
			n.next, n.prev = n, n
			l.head, l.tail = n, n
		default:
			l.tail.next, n.prev = n, l.tail
			l.head.prev, n.next = n, l.head
			l.tail = n
		}
	}
	return l
}

func TestGetFilledList(t *testing.T) {
	testCases := []struct {
		name   string
		values []int
		expected
	}{
		{
			name:     "Create a list whose size is equal to 1",
			values:   []int{10},
			expected: expected{1, []int{10}, []int{10}},
		},
		{
			name:     "Create a list whose size is equal to 2",
			values:   []int{10, 20},
			expected: expected{2, []int{10, 20}, []int{20, 10}},
		},
		{
			name:     "Create a list whose size is equal to 3",
			values:   []int{10, 20, 30},
			expected: expected{3, []int{10, 20, 30}, []int{30, 20, 10}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := getFilledList(tc.values)
			assertSize(t, tc.expected.size, l.size)
			assertValues(t, tc.expected.values, getValues(l))
			assertValues(t, tc.expected.valuesTail, getValuesTail(l))
			assertPointer(t, l.head, l.tail.next)
			assertPointer(t, l.tail, l.head.prev)
		})
	}
}
