package singly

import (
	"fmt"
	"testing"
)

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
		t.Run(fmt.Sprintf("Test case number: %d", i+1), func(t *testing.T) {
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
