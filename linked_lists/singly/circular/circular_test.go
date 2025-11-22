package circular

import (
	"slices"
	"testing"
)

var (
	val    = 50
	values = []int{10, 20, 30}
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		list
		size   uint
		index  uint
		value  int
		values []int
	}{
		{
			name:   "Insert into an empty list with an invalid index",
			list:   list{},
			size:   0,
			index:  1,
			value:  val,
			values: []int{},
		},
		{
			name:   "Insert into an empty list with a valid index",
			list:   list{},
			size:   1,
			index:  0,
			value:  val,
			values: []int{val},
		},
		{
			name:   "Insert into the filled list with an invalid index",
			list:   getFilledList(),
			size:   3,
			index:  4,
			value:  val,
			values: values,
		},
		{
			name:   "Insert into the filled list at the zero index",
			list:   getFilledList(),
			size:   4,
			index:  0,
			value:  val,
			values: []int{val, 10, 20, 30},
		},
		{
			name:   "Insert into the filled list at the first index",
			list:   getFilledList(),
			size:   4,
			index:  1,
			value:  val,
			values: []int{10, val, 20, 30},
		},
		{
			name:   "Insert into the filled list at the second index",
			list:   getFilledList(),
			size:   4,
			index:  2,
			value:  val,
			values: []int{10, 20, val, 30},
		},
		{
			name:   "Insert into the filled list at the end of the list",
			list:   getFilledList(),
			size:   4,
			index:  3,
			value:  val,
			values: []int{10, 20, 30, val},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.values, getValuesFromList(tc.list))
			if i != 0 {
				assertPointer(t, tc.list.head, tc.list.tail.next)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testCase := []struct {
		name string
		list
		size   uint
		index  uint
		values []int
	}{
		{
			name:   "Delete a value from an empty list with an invalid index",
			list:   list{},
			size:   0,
			index:  1,
			values: []int{},
		},
		{
			name:   "Delete a value from an empty list with a valid index",
			list:   list{},
			size:   0,
			index:  1,
			values: []int{},
		},
		{
			name:   "Delete a value from a filled list with an invalid index",
			list:   getFilledList(),
			size:   3,
			index:  3,
			values: []int{10, 20, 30},
		},
		{
			name:   "Delete a value from a filled list at the zero index",
			list:   getFilledList(),
			size:   2,
			index:  0,
			values: []int{20, 30},
		},
		{
			name:   "Delete a value from a filled list at the first index",
			list:   getFilledList(),
			size:   2,
			index:  1,
			values: []int{10, 30},
		},
		{
			name:   "Delete a value from a filled list at the second index",
			list:   getFilledList(),
			size:   2,
			index:  2,
			values: []int{10, 20},
		},
	}
	for i, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.delete(tc.index)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.values, getValuesFromList(tc.list))
			if i > 1 {
				assertPointer(t, tc.list.head, tc.list.tail.next)
			}
		})
	}
}

func getFilledList() list {
	l := list{size: 3}
	n0 := &node{nil, values[0]}
	n1 := &node{nil, values[1]}
	n2 := &node{nil, values[2]}
	l.head, l.tail = n0, n2
	n0.next, n1.next, n2.next = n1, n2, n0
	return l
}

func getValuesFromList(l list) []int {
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
		t.Errorf("Expected result: [ptr: %p value: %[1]v] got: [ptr: %p value: %[2]v]", want, got)
	}
}
