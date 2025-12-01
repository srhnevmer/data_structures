package circular

import (
	"fmt"
	"slices"
	"testing"
)

type expected struct {
	size   uint
	values []int
}

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		list
		index uint
		value int
		expected
	}{
		{
			name:     "Attempt to insert a value into an empty list with an invalid index",
			list:     list{},
			index:    1,
			value:    10,
			expected: expected{0, []int{}},
		},
		{
			name:     "Insert a value into an empty list at the 0 index",
			list:     list{},
			index:    0,
			value:    10,
			expected: expected{1, []int{10}},
		},
		{
			name:     "Attempt to insert a value into a list of size 1 with an invalid index",
			list:     getFilledList([]int{10}),
			index:    2,
			value:    20,
			expected: expected{1, []int{10}},
		},
		{
			name:     "Insert a value into a list of size 1 at the 0 index",
			list:     getFilledList([]int{10}),
			index:    0,
			value:    20,
			expected: expected{2, []int{20, 10}},
		},
		{
			name:     "Insert a value into a list of size 1 at the 1 index",
			list:     getFilledList([]int{10}),
			index:    1,
			value:    20,
			expected: expected{2, []int{10, 20}},
		},
		{
			name:     "Attempt to insert a value into a list of size 2 with an invalid index",
			list:     getFilledList([]int{10, 20}),
			index:    3,
			value:    30,
			expected: expected{2, []int{10, 20}},
		},
		{
			name:     "Insert a value into a list of size 2 at the 0 index",
			list:     getFilledList([]int{10, 20}),
			index:    0,
			value:    30,
			expected: expected{3, []int{30, 10, 20}},
		},
		{
			name:     "Insert a value into a list of size 2 at the 1 index",
			list:     getFilledList([]int{10, 20}),
			index:    1,
			value:    30,
			expected: expected{3, []int{10, 30, 20}},
		},
		{
			name:     "Insert a value into a list of size 2 at the 2 index",
			list:     getFilledList([]int{10, 20}),
			index:    2,
			value:    30,
			expected: expected{3, []int{10, 20, 30}},
		},
		{
			name:     "Attempt to insert a value into a list of size 3 with an invalid index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    4,
			value:    40,
			expected: expected{3, []int{10, 20, 30}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 0 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    0,
			value:    40,
			expected: expected{4, []int{40, 10, 20, 30}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 1 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    1,
			value:    40,
			expected: expected{4, []int{10, 40, 20, 30}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 2 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    2,
			value:    40,
			expected: expected{4, []int{10, 20, 40, 30}},
		},
		{
			name:     "Insert a value into a list of size 3 at the 3 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    3,
			value:    40,
			expected: expected{4, []int{10, 20, 30, 40}},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.expected.size, tc.list.size)
			assertValues(t, tc.expected.values, getValues(tc.list))
			if i != 0 && i != 2 && i != 5 && i != 9 {
				assertPointer(t, tc.list.head, tc.list.tail.next)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name string
		list
		index uint
		expected
	}{
		{
			name:     "Delete from an empty list with an invalid index",
			list:     list{},
			index:    1,
			expected: expected{0, []int{}},
		},
		{
			name:     "Delete from an empty list at the 0 index",
			list:     list{},
			index:    0,
			expected: expected{0, []int{}},
		},
		{
			name:     "Delete from a list of size 1 with an invalid index",
			list:     getFilledList([]int{10}),
			index:    1,
			expected: expected{1, []int{10}},
		},
		{
			name:     "Delete from a list of size 1 at the 0 index",
			list:     getFilledList([]int{10}),
			index:    0,
			expected: expected{0, []int{}},
		},
		{
			name:     "Delete from a list of size 2 with an invalid index",
			list:     getFilledList([]int{10, 20}),
			index:    2,
			expected: expected{2, []int{10, 20}},
		},
		{
			name:     "Delete from a list of size 2 at the 0 index",
			list:     getFilledList([]int{10, 20}),
			index:    0,
			expected: expected{1, []int{20}},
		},
		{
			name:     "Delete from a list of size 2 at the 1 index",
			list:     getFilledList([]int{10, 20}),
			index:    1,
			expected: expected{1, []int{10}},
		},
		{
			name:     "Delete from a list of size 3 with an invalid index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    3,
			expected: expected{3, []int{10, 20, 30}},
		},
		{
			name:     "Delete from a list of size 3 at the 0 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    0,
			expected: expected{2, []int{20, 30}},
		},
		{
			name:     "Delete from a list of size 3 at the 1 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    1,
			expected: expected{2, []int{10, 30}},
		},
		{
			name:     "Delete from a list of size 3 at the 2 index",
			list:     getFilledList([]int{10, 20, 30}),
			index:    2,
			expected: expected{2, []int{10, 20}},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.delete(tc.index)
			assertSize(t, tc.expected.size, tc.list.size)
			assertValues(t, tc.expected.values, getValues(tc.list))
			if i != 0 && i != 1 && i != 3 {
				assertPointer(t, tc.list.head, tc.list.tail.next)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type expected struct {
		index  uint
		result bool
	}
	testCases := []struct {
		name string
		list
		target int
		expected
	}{
		{
			name:     "Search in an empty list but a value doesn't exist",
			list:     list{},
			target:   10,
			expected: expected{0, false},
		},
		{
			name:     "Search in a list of size 1 but a value doesn't exist",
			list:     getFilledList([]int{10}),
			target:   20,
			expected: expected{0, false},
		},
		{
			name:     "Search in a list of size 1 at the 0 index",
			list:     getFilledList([]int{10}),
			target:   10,
			expected: expected{0, true},
		},
		{
			name:     "Search in a list of size 2 but a value doesn't exist",
			list:     getFilledList([]int{10, 20}),
			target:   30,
			expected: expected{0, false},
		},
		{
			name:     "Search in a list of size 2 at the 0 index",
			list:     getFilledList([]int{10, 20}),
			target:   10,
			expected: expected{0, true},
		},
		{
			name:     "Search in a list of size 2 at the 1 index",
			list:     getFilledList([]int{10, 20}),
			target:   20,
			expected: expected{1, true},
		},
		{
			name:     "Search in a list of size 3 but a value doesn't exist",
			list:     getFilledList([]int{10, 20, 30}),
			target:   40,
			expected: expected{0, false},
		},
		{
			name:     "Search in a list of size 3 at the 0 index",
			list:     getFilledList([]int{10, 20, 30}),
			target:   10,
			expected: expected{0, true},
		},
		{
			name:     "Search in a list of size 3 at the 1 index",
			list:     getFilledList([]int{10, 20, 30}),
			target:   20,
			expected: expected{1, true},
		},
		{
			name:     "Search in a list of size 3 at the 2 index",
			list:     getFilledList([]int{10, 20, 30}),
			target:   30,
			expected: expected{2, true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			idx, ok := tc.list.search(tc.target)
			if want, got := tc.expected.index, idx; want != got {
				t.Fatalf("Expected index: %d got: %d", want, got)
			}
			if want, got := tc.expected.result, ok; want != got {
				t.Errorf("Expected result: %t got: %t", want, got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		name string
		list
		expectedValues []int
	}{
		{
			name:           "Reverse a list of size 1",
			list:           getFilledList([]int{10}),
			expectedValues: []int{10},
		},
		{
			name:           "Reverse a list of size 2",
			list:           getFilledList([]int{10, 20}),
			expectedValues: []int{20, 10},
		},
		{
			name:           "Reverse a list of size 3",
			list:           getFilledList([]int{10, 20, 30}),
			expectedValues: []int{30, 20, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.reverse()
			assertValues(t, tc.expectedValues, getValues(tc.list))
			assertPointer(t, tc.list.head, tc.list.tail.next)
		})
	}
}

func getFilledList(values []int) list {
	l := list{size: uint(len(values))}
	n := &node{nil, values[0]}
	values = values[1:]
	n.next = n
	l.head, l.tail = n, n

	for i := 0; i < len(values); i++ {
		n := &node{l.head, values[i]}
		l.tail.next, l.tail = n, n
	}

	return l
}

func TestGetFilledList(t *testing.T) {
	testCases := []struct {
		size           uint
		values         []int
		expectedValues []int
	}{
		{size: 1, values: []int{10}, expectedValues: []int{10}},
		{size: 2, values: []int{10, 20}, expectedValues: []int{10, 20}},
		{size: 3, values: []int{10, 20, 30}, expectedValues: []int{10, 20, 30}},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case number: %d", i+1), func(t *testing.T) {
			l := getFilledList(tc.values)
			assertSize(t, tc.size, l.size)
			assertValues(t, tc.expectedValues, getValues(l))
			assertPointer(t, l.head, l.tail.next)
		})
	}
}

func getValues(l list) []int {
	values := make([]int, 0, l.size)
	for i, curr := uint(0), l.head; i < l.size; i, curr = i+1, curr.next {
		values = append(values, curr.value)
	}
	return values
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
		t.Fatalf("Expected values: %v got: %v", want, got)
	}
}

func assertPointer(t testing.TB, want, got *node) {
	t.Helper()
	if want != got {
		t.Errorf("Expected pointer: [ptr: %p value: %[1]v] got: [ptr: %p value: %[2]v]", want, got)
	}
}
