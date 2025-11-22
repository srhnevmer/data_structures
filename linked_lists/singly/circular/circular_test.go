package circular

import (
	"fmt"
	"slices"
	"testing"
)

var (
	value        = 50
	values       = []int{10, 20, 30}
	invalidIndex = uint(10)
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		list
		size                  uint
		index                 uint
		payload               int
		valuesInExpectedOrder []int
	}{
		{
			name:                  "Attempt to insert into an empty list with an invalid index",
			list:                  list{},
			size:                  0,
			index:                 invalidIndex,
			payload:               value,
			valuesInExpectedOrder: []int{},
		},
		{
			name:                  "Insert into an empty list with a valid index",
			list:                  list{},
			size:                  1,
			index:                 0,
			payload:               value,
			valuesInExpectedOrder: []int{value},
		},
		{
			name:                  "Attempt to insert into a filled list with an invalid index",
			list:                  getFilledList(values),
			size:                  3,
			index:                 invalidIndex,
			payload:               value,
			valuesInExpectedOrder: values,
		},
		{
			name:                  "Insert into a filled list at the zero index",
			list:                  getFilledList(values),
			size:                  4,
			index:                 0,
			payload:               value,
			valuesInExpectedOrder: []int{value, 10, 20, 30},
		},
		{
			name:                  "Insert into a filled list at the first index",
			list:                  getFilledList(values),
			size:                  4,
			index:                 1,
			payload:               value,
			valuesInExpectedOrder: []int{10, value, 20, 30},
		},
		{
			name:                  "Insert into a filled list at the second index",
			list:                  getFilledList(values),
			size:                  4,
			index:                 2,
			payload:               value,
			valuesInExpectedOrder: []int{10, 20, value, 30},
		},
		{
			name:                  "Insert into a filled list at the end of the list",
			list:                  getFilledList(values),
			size:                  4,
			index:                 3,
			payload:               value,
			valuesInExpectedOrder: []int{10, 20, 30, value},
		},
	}
	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.insert(tc.index, tc.payload)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.valuesInExpectedOrder, getValuesFromList(tc.list))
			if i != 0 {
				assertPointer(t, tc.list.head, tc.list.tail.next)
			}
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
