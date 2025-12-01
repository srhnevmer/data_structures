package singly

import (
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
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.expected.size, tc.list.size)
			assertValues(t, tc.expected.values, getValues(tc.list))
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
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.delete(tc.index)
			assertSize(t, tc.expected.size, tc.list.size)
			assertValues(t, tc.expected.values, getValues(tc.list))
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name string
		list
		target         int
		expectedIndex  uint
		expectedResult bool
	}{
		{
			name:           "Attempt to find a value in an empty list",
			list:           list{},
			target:         50,
			expectedIndex:  0,
			expectedResult: false,
		},
		{
			name:           "Attempt to find a value in a filled list but a value doesn't exist",
			list:           list{},
			target:         50,
			expectedIndex:  0,
			expectedResult: false,
		},
		{
			name:           "Find the value that equals 10",
			list:           getFilledList([]int{10, 20, 30}),
			target:         10,
			expectedIndex:  0,
			expectedResult: true,
		},
		{
			name:           "Find the value that equals 20",
			list:           getFilledList([]int{10, 20, 30}),
			target:         20,
			expectedIndex:  1,
			expectedResult: true,
		},
		{
			name:           "Find the value that equals 30",
			list:           getFilledList([]int{10, 20, 30}),
			target:         30,
			expectedIndex:  2,
			expectedResult: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			idx, ok := tc.list.search(tc.target)
			if want, got := tc.expectedIndex, idx; want != got {
				t.Fatalf("Expected index: %d got: %d", want, got)
			}
			if want, got := tc.expectedResult, ok; want != got {
				t.Errorf("Expected result: %t, got: %t", want, got)
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
			name:           "Attempt to reverse an empty list",
			list:           list{},
			expectedValues: []int{},
		},
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
		t.Errorf("Expected values: %v, got: %v", want, got)
	}
}

func getValues(l list) []int {
	v := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		v = append(v, curr.value)
	}
	return v
}

func getFilledList(values []int) list {
	var prev *node
	l := list{size: uint(len(values))}
	for i := range values {
		n := &node{nil, values[i]}
		switch i {
		case 0:
			l.head = n
		default:
			prev.next = n
		}
		prev = n
	}
	return l
}
