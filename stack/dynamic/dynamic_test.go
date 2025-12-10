package dynamic

import (
	"slices"
	"testing"
)

type expected struct {
	size   uint
	values []int
}

func TestPush(t *testing.T) {
	testCases := []struct {
		name string
		stack
		value int
		expected
	}{
		{
			name:     "Push a value in an empty stack",
			stack:    stack{},
			value:    10,
			expected: expected{1, []int{10}},
		},
		{
			name:     "Push a value in a stack of size 1",
			stack:    getFilledStack([]int{10}),
			value:    20,
			expected: expected{2, []int{20, 10}},
		},
		{
			name:     "Push a value in a stack of size 2",
			stack:    getFilledStack([]int{10, 20}),
			value:    30,
			expected: expected{3, []int{30, 20, 10}},
		},
		{
			name:     "Push a value in a stack of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			value:    40,
			expected: expected{4, []int{40, 30, 20, 10}},
		},
		{
			name:     "Push a value in a stack of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			value:    50,
			expected: expected{5, []int{50, 40, 30, 20, 10}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stack.push(tc.value)
			assertSize(t, tc.expected.size, tc.stack.size)
			assertValues(t, tc.expected.values, getValues(tc.stack))
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
		t.Errorf("Expected values: %v got: %v", want, got)
	}
}

func getFilledStack(values []int) stack {
	s := stack{size: uint(len(values))}
	for i, v := range values {
		switch i {
		case 0:
			s.head = &node{nil, v}
		default:
			s.head = &node{s.head, v}
		}
	}
	return s
}

func getValues(s stack) []int {
	v := make([]int, 0, int(s.size))
	for curr := s.head; curr != nil; curr = curr.next {
		v = append(v, curr.value)
	}
	return v
}
