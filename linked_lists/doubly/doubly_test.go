package doubly

import (
	"slices"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		list
		size                   uint
		index                  uint
		value                  int
		expectedValues         []int
		expectedValuesFromTail []int
	}{
		{
			name:                   "Attempt to insert into an empty list with an invalid index",
			list:                   list{},
			size:                   0,
			index:                  1,
			value:                  50,
			expectedValues:         []int{},
			expectedValuesFromTail: []int{},
		},
		{
			name:                   "Insert into an empty list with a valid index",
			list:                   list{},
			size:                   1,
			index:                  0,
			value:                  50,
			expectedValues:         []int{50},
			expectedValuesFromTail: []int{50},
		},
		{
			name:                   "Attempt to insert into a filled list with an invalid index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   3,
			index:                  4,
			value:                  50,
			expectedValues:         []int{10, 20, 30},
			expectedValuesFromTail: []int{30, 20, 10},
		},
		{
			name:                   "Insert into a filled list at the zero index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   4,
			index:                  0,
			value:                  50,
			expectedValues:         []int{50, 10, 20, 30},
			expectedValuesFromTail: []int{30, 20, 10, 50},
		},
		{
			name:                   "Insert into a filled list at the first index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   4,
			index:                  1,
			value:                  50,
			expectedValues:         []int{10, 50, 20, 30},
			expectedValuesFromTail: []int{30, 20, 50, 10},
		},
		{
			name:                   "Insert into a filled list at the second index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   4,
			index:                  2,
			value:                  50,
			expectedValues:         []int{10, 20, 50, 30},
			expectedValuesFromTail: []int{30, 50, 20, 10},
		},
		{
			name:                   "Insert into a filled list at the end of the list",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   4,
			index:                  3,
			value:                  50,
			expectedValues:         []int{10, 20, 30, 50},
			expectedValuesFromTail: []int{50, 30, 20, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.expectedValues, getValues(tc.list))
			assertValues(t, tc.expectedValuesFromTail, getValuesFromTail(tc.list))
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name string
		list
		size                   uint
		index                  uint
		expectedValues         []int
		expectedValuesFromTail []int
	}{
		{
			name:                   "Attempt to delete a value from an empty list",
			list:                   list{},
			size:                   0,
			index:                  0,
			expectedValues:         []int{},
			expectedValuesFromTail: []int{},
		},
		{
			name:                   "Delete a value from a list of size 1 with a valid index",
			list:                   getFilledList([]int{10}),
			size:                   0,
			index:                  0,
			expectedValues:         []int{},
			expectedValuesFromTail: []int{},
		},
		{
			name:                   "Attempt to delete a value from a list of size 1 with an invalid index",
			list:                   getFilledList([]int{10}),
			size:                   1,
			index:                  1,
			expectedValues:         []int{10},
			expectedValuesFromTail: []int{10},
		},
		{
			name:                   "Attempt to delete a value from a filled list with an invalid index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   3,
			index:                  3,
			expectedValues:         []int{10, 20, 30},
			expectedValuesFromTail: []int{30, 20, 10},
		},
		{
			name:                   "Delete a value from a filled list at the zero index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   2,
			index:                  0,
			expectedValues:         []int{20, 30},
			expectedValuesFromTail: []int{30, 20},
		},
		{
			name:                   "Delete a value from a filled list at the first index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   2,
			index:                  1,
			expectedValues:         []int{10, 30},
			expectedValuesFromTail: []int{30, 10},
		},
		{
			name:                   "Delete a value from a filled list at the second index",
			list:                   getFilledList([]int{10, 20, 30}),
			size:                   2,
			index:                  2,
			expectedValues:         []int{10, 20},
			expectedValuesFromTail: []int{20, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.delete(tc.index)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.expectedValues, getValues(tc.list))
			assertValues(t, tc.expectedValuesFromTail, getValuesFromTail(tc.list))
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
		t.Fatalf("Expected values: %v got: %v", want, got)
	}
}

func getValues(l list) []int {
	v := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		v = append(v, curr.value)
	}
	return v
}

func getValuesFromTail(l list) []int {
	if l.head == nil {
		return []int{}
	}
	curr := l.head
	v := make([]int, 0, l.size)
	for curr.next != nil {
		curr = curr.next
	}
	for curr != nil {
		v = append(v, curr.value)
		curr = curr.prev
	}
	return v
}

func getFilledList(values []int) list {
	var prev *node
	l := list{size: uint(len(values))}
	for i := range len(values) {
		n := &node{nil, nil, values[i]}
		switch l.head {
		case nil:
			l.head = n
		default:
			prev.next, n.prev = n, prev
		}
		prev = n
	}
	return l
}
