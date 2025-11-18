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
			values: []int{10, 20, 30, 40},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			if want, got := tc.size, tc.list.size; want != got {
				t.Fatalf("Expected size: %d got: %d", want, got)
			}

			if want, got := tc.values, getValuesFromList(tc.list); slices.Compare(want, got) != 0 {
				t.Fatalf("Expected values: %v got: %v", want, got)
			}

			if i != 0 {
				if want, got := tc.list.head, tc.list.tail.next; want != got {
					t.Errorf("Expected result: [ptr: %p value: %[1]v] got: [ptr: %p value: %[2]v]", want, got)
				}
			}
		})
	}
}

func getFilledList() list {
	l := list{}
	for i := range values {
		l.insert(uint(i), values[i])
	}
	return l
}

func getValuesFromList(l list) []int {
	values := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		values = append(values, curr.value)
	}
	return values
}
