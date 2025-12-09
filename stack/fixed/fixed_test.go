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
			name:     "Peek a value from a list of size 1",
			stack:    getFilledStack([]int{10}),
			expected: expected{10, true},
		},
		{
			name:     "Peek a value from a list of size 2",
			stack:    getFilledStack([]int{10, 20}),
			expected: expected{20, true},
		},
		{
			name:     "Peek a value from a list of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			expected: expected{30, true},
		},
		{
			name:     "Peek a value from a list of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			expected: expected{40, true},
		},
		{
			name:     "Peek a value from a list of size 5",
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
