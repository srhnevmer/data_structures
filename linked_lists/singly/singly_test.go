package singly

import (
	"fmt"
	"slices"
	"testing"
)

var (
	val              = 50
	nums             = []int{10, 20, 30}
	testNameTemplate = "Test case number:"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		list  list
		index uint
		value int
		size  uint
		want  []int
	}{
		{list: list{}, index: 1, value: val, size: 0, want: []int{}},
		{list: list{}, index: 0, value: val, size: 1, want: []int{50}},
		{list: fillList(nums), index: 4, value: val, size: 3, want: []int{10, 20, 30}},
		{list: fillList(nums), index: 0, value: val, size: 4, want: []int{50, 10, 20, 30}},
		{list: fillList(nums), index: 1, value: val, size: 4, want: []int{10, 50, 20, 30}},
		{list: fillList(nums), index: 2, value: val, size: 4, want: []int{10, 20, 50, 30}},
		{list: fillList(nums), index: 3, value: val, size: 4, want: []int{10, 20, 30, 50}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", testNameTemplate, i+1), func(t *testing.T) {
			tc.list.insert(tc.index, tc.value)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.want, extractValues(tc.list))
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		list  list
		index uint
		size  uint
		want  []int
	}{
		{list: list{}, index: 1, size: 0, want: []int{}},
		{list: list{}, index: 0, size: 0, want: []int{}},
		{list: fillList(nums[:1]), index: 1, size: 1, want: []int{10}},
		{list: fillList(nums[:1]), index: 0, size: 0, want: []int{}},
		{list: fillList(nums), index: 0, size: 2, want: []int{20, 30}},
		{list: fillList(nums), index: 1, size: 2, want: []int{10, 30}},
		{list: fillList(nums), index: 2, size: 2, want: []int{10, 20}},
		{list: fillList(nums), index: 3, size: 3, want: []int{10, 20, 30}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", testNameTemplate, i+1), func(t *testing.T) {
			tc.list.delete(tc.index)
			assertSize(t, tc.size, tc.list.size)
			assertValues(t, tc.want, extractValues(tc.list))
		})
	}
}

func TestTraverse(t *testing.T) {
	testCases := []struct {
		list
		want []int
	}{
		{list: list{}, want: []int{}},
		{list: fillList(nums[:1]), want: []int{10}},
		{list: fillList(nums[:2]), want: []int{10, 20}},
		{list: fillList(nums), want: []int{10, 20, 30}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", testNameTemplate, i+1), func(t *testing.T) {
			assertValues(t, tc.want, tc.list.traverse())
		})
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		list
		value int
		want  bool
	}{
		{list: list{}, value: val, want: false},
		{list: fillList(nums), value: 10, want: true},
		{list: fillList(nums), value: 20, want: true},
		{list: fillList(nums), value: 30, want: true},
		{list: fillList(nums), value: val, want: false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", testNameTemplate, i+1), func(t *testing.T) {
			if got := tc.list.search(tc.value); tc.want != got {
				t.Errorf("Expected result: %t got: %t", tc.want, got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		list
		want []int
	}{
		{list: list{}, want: []int{}},
		{list: fillList(nums[:1]), want: []int{10}},
		{list: fillList(nums[:2]), want: []int{20, 10}},
		{list: fillList(nums), want: []int{30, 20, 10}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", testNameTemplate, i+1), func(t *testing.T) {
			tc.list.reverse()
			assertValues(t, tc.want, extractValues(tc.list))
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

func fillList(values []int) list {
	n := &node{nil, values[0]}
	l, values := list{}, values[1:]
	l.head, l.size = n, 1

	for i, curr := 0, l.head; i < len(values); i, curr = i+1, curr.next {
		curr.next = &node{nil, values[i]}
		l.size++
	}
	return l
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
		t.Errorf("Expected result: %v, got: %v", want, got)
	}
}
