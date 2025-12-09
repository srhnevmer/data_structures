package fixed

import "testing"

func TestPeek(t *testing.T) {
	type expected struct {
		value  int
		result bool
	}
	testCases := []struct {
		name string
		stack
		expected
	}{
		{
			name:     "Attempt to peek a value from an empty stack",
			stack:    initStack(),
			expected: expected{0, false},
		},
		{
			name:     "Peek a value from a stack of size 1",
			stack:    getFilledStack([]int{10}),
			expected: expected{10, true},
		},
		{
			name:     "Peek a value from a stack of size 2",
			stack:    getFilledStack([]int{10, 20}),
			expected: expected{20, true},
		},
		{
			name:     "Peek a value from a stack of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			expected: expected{30, true},
		},
		{
			name:     "Peek a value from a stack of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			expected: expected{40, true},
		},
		{
			name:     "Peek a value from a stack of size 5",
			stack:    getFilledStack([]int{10, 20, 30, 40, 50}),
			expected: expected{50, true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := tc.stack.peek()
			if want, got := tc.expected.value, val; want != got {
				t.Fatalf("Expected value: %d got: %d", want, got)
			}
			if want, got := tc.expected.result, ok; want != got {
				t.Errorf("Expected result: %t got: %t", want, got)
			}
		})
	}
}

func TestPush(t *testing.T) {
	type expected struct {
		size      int
		container [5]int
	}
	testCases := []struct {
		name string
		stack
		value int
		expected
	}{
		{
			name:     "Add a value in an empty stack",
			stack:    initStack(),
			value:    10,
			expected: expected{0, [5]int{10, 0, 0, 0, 0}},
		},
		{
			name:     "Add a value in a stack of size 1",
			stack:    getFilledStack([]int{10}),
			value:    20,
			expected: expected{1, [5]int{10, 20, 0, 0, 0}},
		},
		{
			name:     "Add a value in a stack of size 2",
			stack:    getFilledStack([]int{10, 20}),
			value:    30,
			expected: expected{2, [5]int{10, 20, 30, 0, 0}},
		},
		{
			name:     "Add a value in a stack of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			value:    40,
			expected: expected{3, [5]int{10, 20, 30, 40, 0}},
		},
		{
			name:     "Add a value in a stack of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			value:    50,
			expected: expected{4, [5]int{10, 20, 30, 40, 50}},
		},
		{
			name:     "Attempt to add a value into a filled stack",
			stack:    getFilledStack([]int{10, 20, 30, 40, 50}),
			value:    60,
			expected: expected{4, [5]int{10, 20, 30, 40, 50}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stack.push(tc.value)
			if want, got := tc.expected.size, tc.stack.size; want != got {
				t.Fatalf("Expected size: %d got: %d", want, got)
			}
			if want, got := tc.expected.container, tc.stack.container; want != got {
				t.Errorf("Expected result: %v got: %v", want, got)
			}
		})
	}
}

func initStack() stack {
	return stack{size: -1}
}

func getFilledStack(values []int) stack {
	s := stack{}
	for i, v := range values {
		s.size = i
		s.container[i] = v
	}
	return s
}
