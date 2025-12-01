package singly

import (
	"slices"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		list
		size           uint
		index          uint
		value          int
		expectedValues []int
	}{
		{
			name:           "Attempt to insert into an empty list with an invalid index",
			list:           list{},
			size:           0,
			index:          1,
			value:          50,
			expectedValues: []int{},
		},
		{
			name:           "Insert into an empty list with a valid index",
			list:           list{},
			size:           1,
			index:          0,
			value:          50,
			expectedValues: []int{50},
		},
		{
			name:           "Attempt to insert into a filled list with an invalid index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           3,
			index:          4,
			value:          50,
			expectedValues: []int{10, 20, 30},
		},
		{
			name:           "Insert into a filled list at the zero index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           4,
			index:          0,
			value:          50,
			expectedValues: []int{50, 10, 20, 30},
		},
		{
			name:           "Insert into a filled list at the first index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           4,
			index:          1,
			value:          50,
			expectedValues: []int{10, 50, 20, 30},
		},
		{
			name:           "Insert into a filled list at the second index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           4,
			index:          2,
			value:          50,
			expectedValues: []int{10, 20, 50, 30},
		},
		{
			name:           "Insert into a filled list at the end of the list",
			list:           getFilledList([]int{10, 20, 30}),
			size:           4,
			index:          3,
			value:          50,
			expectedValues: []int{10, 20, 30, 50},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.expectedValues, getValues(tc.list))
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name string
		list
		size           uint
		index          uint
		expectedValues []int
	}{
		{
			name:           "Attempt to delete a value from an empty list",
			list:           list{},
			size:           0,
			index:          0,
			expectedValues: []int{},
		},
		{
			name:           "Delete a value from a list of size 1 with a valid index",
			list:           getFilledList([]int{10}),
			size:           0,
			index:          0,
			expectedValues: []int{},
		},
		{
			name:           "Attempt to delete a value from a list of size 1 with an invalid index",
			list:           getFilledList([]int{10}),
			size:           1,
			index:          1,
			expectedValues: []int{10},
		},
		{
			name:           "Attempt to delete a value from a filled list with an invalid index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           3,
			index:          3,
			expectedValues: []int{10, 20, 30},
		},
		{
			name:           "Delete a value from a filled list at the zero index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           2,
			index:          0,
			expectedValues: []int{20, 30},
		},
		{
			name:           "Delete a value from a filled list at the first index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           2,
			index:          1,
			expectedValues: []int{10, 30},
		},
		{
			name:           "Delete a value from a filled list at the second index",
			list:           getFilledList([]int{10, 20, 30}),
			size:           2,
			index:          2,
			expectedValues: []int{10, 20},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.delete(tc.index)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.expectedValues, getValues(tc.list))
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
