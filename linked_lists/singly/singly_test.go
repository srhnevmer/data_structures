package singly

import (
	"fmt"
	"slices"
	"testing"
)

const template = "Test case number:"

func TestInsert(t *testing.T) {
	nums := []int{10, 20, 30}
	testCases := []struct {
		list  list
		size  int
		index int
		value int
		want  []int
	}{
		{list: list{}, index: -1, value: 10, size: 0, want: []int{}},
		{list: list{}, index: 0, value: 50, size: 1, want: []int{50}},
		{list: fillList(nums...), index: 0, value: 50, size: 4, want: []int{50, 10, 20, 30}},
		{list: fillList(nums...), index: 1, value: 50, size: 4, want: []int{10, 50, 20, 30}},
		{list: fillList(nums...), index: 2, value: 50, size: 4, want: []int{10, 20, 50, 30}},
		{list: fillList(nums...), index: 3, value: 50, size: 4, want: []int{10, 20, 30, 50}},
		{list: fillList(nums...), index: -1, value: 50, size: 3, want: []int{10, 20, 30}},
		{list: fillList(nums...), index: 4, value: 50, size: 3, want: []int{10, 20, 30}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", template, i+1), func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)

			if tc.size != tc.list.size {
				t.Fatalf("Expected size: %d got: %d", tc.size, tc.list.size)
			}

			if v := extractValues(tc.list); slices.Compare(tc.want, v) != 0 {
				t.Errorf("Expected result: %v got %v", tc.want, v)
			}
		})
	}
}

func extractValues(l list) []int {
	r := make([]int, 0, l.size)
	for curr := l.head; curr != nil; curr = curr.next {
		r = append(r, curr.value)
	}
	return r
}

func fillList(values ...int) list {
	n := &node{nil, values[0]}
	l, values := list{}, values[1:]
	l.head, l.size = n, 1

	for i, curr := 0, l.head; i < len(values); i++ {
		curr.next = &node{nil, values[i]}
		curr = curr.next
		l.size++
	}

	return l
}
