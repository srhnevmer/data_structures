package singly

import (
	"fmt"
	"slices"
	"testing"
)

const template = "Test case number:"

func TestAdd(t *testing.T) {
	testCases := []struct {
		list
		size      int
		lastValue int
	}{
		{list: fillList(), size: 0, lastValue: 0},
		{list: fillList(10), size: 1, lastValue: 10},
		{list: fillList(10, 20, 30), size: 3, lastValue: 30},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", template, i+1), func(t *testing.T) {
			if tc.size != tc.list.size {
				t.Errorf("Expected size: %d got %d", tc.size, tc.list.size)
			}

			if lv := getLastValList(tc.list); tc.lastValue != lv {
				t.Errorf("Expected value: %d got %d", tc.lastValue, lv)
			}
		})
	}
}

func fillList(values ...int) list {
	l := list{}
	for _, v := range values {
		l.add(v)
	}
	return l
}

func getLastValList(l list) int {
	if l.head == nil {
		return 0
	}

	var r int
	for i, curr := 0, l.head; i < l.size; i++ {
		r = curr.value
		curr = curr.next
	}

	return r
}

func TestTraversal(t *testing.T) {
	testCases := []struct {
		list
		want []int
	}{
		{list: fillList(), want: make([]int, 0)},
		{list: fillList(10), want: []int{10}},
		{list: fillList(10, 20, 30), want: []int{10, 20, 30}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s%d", template, i+1), func(t *testing.T) {
			if got := tc.list.traversal(); slices.Compare(tc.want, got) != 0 {
				t.Errorf("Expected result: %v got %v", tc.want, got)
			}
		})
	}
}
