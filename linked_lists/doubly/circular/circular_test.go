package circular

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
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.expectedValues, getValues(tc.list))
			assertValues(t, tc.expectedValuesFromTail, getValuesTail(tc.list))
			if i != 0 {
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
			l.tail = n
			l.head.prev, l.tail.next = l.tail, l.head
		}
	}
	return l
}

func TestGetFilledList(t *testing.T) {
	testCases := []struct {
		name                   string
		size                   uint
		values                 []int
		expectedValues         []int
		expectedValuesFromTail []int
	}{
		{
			name:                   "Create a list whose size is equal to 1",
			size:                   1,
			values:                 []int{10},
			expectedValues:         []int{10},
			expectedValuesFromTail: []int{10},
		},
		{
			name:                   "Create a list whose size is equal to 2",
			size:                   2,
			values:                 []int{10, 20},
			expectedValues:         []int{10, 20},
			expectedValuesFromTail: []int{20, 10},
		},
		{
			name:                   "Create a list whose size is equal to 3",
			size:                   3,
			values:                 []int{10, 20, 30},
			expectedValues:         []int{10, 20, 30},
			expectedValuesFromTail: []int{30, 20, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := getFilledList(tc.values)
			assertSize(t, tc.size, l.size)
			assertValues(t, tc.expectedValues, getValues(l))
			assertValues(t, tc.expectedValuesFromTail, getValuesTail(l))
			assertPointer(t, l.head, l.tail.next)
			assertPointer(t, l.tail, l.head.prev)
		})
	}
}
