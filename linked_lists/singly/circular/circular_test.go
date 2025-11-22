package circular

import (
	"fmt"
	"slices"
	"testing"
)

var (
	val    = 50
	values = []int{10, 20, 30}
)

// func TestInsert(t *testing.T) {
// 	testCases := []struct {
// 		name string
// 		list
// 		size   uint
// 		index  uint
// 		value  int
// 		values []int
// 	}{
// 		{
// 			name:   "Insert into an empty list with an invalid index",
// 			list:   list{},
// 			size:   0,
// 			index:  1,
// 			value:  val,
// 			values: []int{},
// 		},
// 		{
// 			name:   "Insert into an empty list with a valid index",
// 			list:   list{},
// 			size:   1,
// 			index:  0,
// 			value:  val,
// 			values: []int{val},
// 		},
// 		{
// 			name:   "Insert into the filled list with an invalid index",
// 			list:   getFilledList(values),
// 			size:   3,
// 			index:  4,
// 			value:  val,
// 			values: values,
// 		},
// 		{
// 			name:   "Insert into the filled list at the zero index",
// 			list:   getFilledList(values),
// 			size:   4,
// 			index:  0,
// 			value:  val,
// 			values: []int{val, 10, 20, 30},
// 		},
// 		{
// 			name:   "Insert into the filled list at the first index",
// 			list:   getFilledList(values),
// 			size:   4,
// 			index:  1,
// 			value:  val,
// 			values: []int{10, val, 20, 30},
// 		},
// 		{
// 			name:   "Insert into the filled list at the second index",
// 			list:   getFilledList(values),
// 			size:   4,
// 			index:  2,
// 			value:  val,
// 			values: []int{10, 20, val, 30},
// 		},
// 		{
// 			name:   "Insert into the filled list at the end of the list",
// 			list:   getFilledList(values),
// 			size:   4,
// 			index:  3,
// 			value:  val,
// 			values: []int{10, 20, 30, val},
// 		},
// 	}
// 	for i, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			tc.list.insert(tc.index, tc.value)
// 			assertSize(t, tc.size, tc.list.size)
// 			assertValues(t, tc.values, getValuesFromList(tc.list))
// 			if i != 0 {
// 				assertPointer(t, tc.list.head, tc.list.tail.next)
// 			}
// 		})
// 	}
// }

// func TestDelete(t *testing.T) {
// 	testCase := []struct {
// 		name string
// 		list
// 		size   uint
// 		index  uint
// 		values []int
// 	}{
// 		{
// 			name:   "Delete a value from an empty list with an invalid index",
// 			list:   list{},
// 			size:   0,
// 			index:  1,
// 			values: []int{},
// 		},
// 		{
// 			name:   "Delete a value from an empty list with a valid index",
// 			list:   list{},
// 			size:   0,
// 			index:  1,
// 			values: []int{},
// 		},
// 		{
// 			name:   "Delete a value from a filled list with an invalid index",
// 			list:   getFilledList(),
// 			size:   3,
// 			index:  3,
// 			values: []int{10, 20, 30},
// 		},
// 		{
// 			name:   "Delete a value from a filled list at the zero index",
// 			list:   getFilledList(),
// 			size:   2,
// 			index:  0,
// 			values: []int{20, 30},
// 		},
// 		{
// 			name:   "Delete a value from a filled list at the first index",
// 			list:   getFilledList(),
// 			size:   2,
// 			index:  1,
// 			values: []int{10, 30},
// 		},
// 		{
// 			name:   "Delete a value from a filled list at the second index",
// 			list:   getFilledList(),
// 			size:   2,
// 			index:  2,
// 			values: []int{10, 20},
// 		},
// 	}
// 	for i, tc := range testCase {
// 		t.Run(tc.name, func(t *testing.T) {
// 			tc.list.delete(tc.index)
// 			assertSize(t, tc.size, tc.list.size)
// 			assertValues(t, tc.values, getValuesFromList(tc.list))
// 			if i > 1 {
// 				assertPointer(t, tc.list.head, tc.list.tail.next)
// 			}
// 		})
// 	}
// }

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
		size    uint
		payload []int
		want    []int
	}{
		{size: 1, payload: values[:1], want: []int{10}},
		{size: 2, payload: values[:2], want: []int{10, 20}},
		{size: 3, payload: values, want: values},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case number %d", i+1), func(t *testing.T) {
			l := getFilledList(tc.payload)
			assertSize(t, tc.size, l.size)
			assertValues(t, tc.want, getValuesFromList(l))
			assertPointer(t, l.head, l.tail.next)
		})
	}
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
		t.Errorf("Expected pointer: [ptr: %p value: %[1]v] got: [ptr: %p value: %[2]v]", want, got)
	}
}
