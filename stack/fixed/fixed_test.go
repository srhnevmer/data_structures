package fixed

import "testing"

type expected struct {
	size      int
	container [max]int
}

type expected2 struct {
	value  int
	result bool
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
			stack:    initStack(),
			value:    10,
			expected: expected{0, [max]int{10}},
		},
		{
			name:     "Push a value in a stack of size 1",
			stack:    getFilledStack([]int{10}),
			value:    20,
			expected: expected{1, [max]int{10, 20}},
		},
		{
			name:     "Push a value in a stack of size 2",
			stack:    getFilledStack([]int{10, 20}),
			value:    30,
			expected: expected{2, [max]int{10, 20, 30}},
		},
		{
			name:     "Push a value in a stack of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			value:    40,
			expected: expected{3, [max]int{10, 20, 30, 40}},
		},
		{
			name:     "Push a value in a stack of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			value:    50,
			expected: expected{4, [max]int{10, 20, 30, 40, 50}},
		},
		{
			name:     "Attempt to push a value in a filled stack",
			stack:    getFilledStack([]int{10, 20, 30, 40, 50}),
			value:    60,
			expected: expected{4, [max]int{10, 20, 30, 40, 50}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stack.push(tc.value)
			assertSize(t, tc.expected.size, tc.stack.size)
			assertContainer(t, tc.expected.container, tc.stack.container)
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		name string
		stack
		expected
	}{
		{
			name:     "Attempt to pop a value from an empty stack",
			stack:    initStack(),
			expected: expected{-1, [max]int{}},
		},
		{
			name:     "Pop a value from a stack of size 1",
			stack:    getFilledStack([]int{10}),
			expected: expected{-1, [max]int{}},
		},
		{
			name:     "Pop a value from a stack of size 2",
			stack:    getFilledStack([]int{10, 20}),
			expected: expected{0, [max]int{10}},
		},
		{
			name:     "Pop a value from a stack of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			expected: expected{1, [max]int{10, 20}},
		},
		{
			name:     "Pop a value from a stack of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			expected: expected{2, [max]int{10, 20, 30}},
		},
		{
			name:     "Pop a value from a stack of size 5",
			stack:    getFilledStack([]int{10, 20, 30, 40, 50}),
			expected: expected{3, [max]int{10, 20, 30, 40}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stack.pop()
			assertSize(t, tc.expected.size, tc.stack.size)
			assertContainer(t, tc.expected.container, tc.stack.container)
		})
	}
}

func TestPeek(t *testing.T) {
	testCases := []struct {
		name string
		stack
		expected2
	}{
		{
			name:      "Attempt to peek a value from an empty stack",
			stack:     initStack(),
			expected2: expected2{0, false},
		},
		{
			name:      "Peek a value from a stack of size 1",
			stack:     getFilledStack([]int{10}),
			expected2: expected2{10, true},
		},
		{
			name:      "Peek a value from a stack of size 2",
			stack:     getFilledStack([]int{10, 20}),
			expected2: expected2{20, true},
		},
		{
			name:      "Peek a value from a stack of size 3",
			stack:     getFilledStack([]int{10, 20, 30}),
			expected2: expected2{30, true},
		},
		{
			name:      "Peek a value from a stack of size 4",
			stack:     getFilledStack([]int{10, 20, 30, 40}),
			expected2: expected2{40, true},
		},
		{
			name:      "Peek a value from a stack of size 5",
			stack:     getFilledStack([]int{10, 20, 30, 40, 50}),
			expected2: expected2{50, true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := tc.stack.peek()
			if want, got := tc.expected2.value, val; want != got {
				t.Fatalf("Expected value: %d got: %d", want, got)
			}
			assertResult(t, tc.expected2.result, ok)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		stack
		expected2
	}{
		{
			name:      "Check if an empty stack is empty",
			stack:     initStack(),
			expected2: expected2{result: true},
		},
		{
			name:      "Check if a stack of size 1 is empty",
			stack:     getFilledStack([]int{10}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 2 is empty",
			stack:     getFilledStack([]int{10, 20}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 3 is empty",
			stack:     getFilledStack([]int{10, 20, 30}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 4 is empty",
			stack:     getFilledStack([]int{10, 20, 30, 40}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 5 is empty",
			stack:     getFilledStack([]int{10, 20, 30, 40, 50}),
			expected2: expected2{result: false},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.stack.isEmpty()
			assertResult(t, tc.expected2.result, result)
		})
	}
}

func TestIsFull(t *testing.T) {
	testCases := []struct {
		name string
		stack
		expected2
	}{
		{
			name:      "Check if an empty stack is full",
			stack:     initStack(),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 1 is full",
			stack:     getFilledStack([]int{10}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 2 is full",
			stack:     getFilledStack([]int{10, 20}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 3 is full",
			stack:     getFilledStack([]int{10, 20, 30}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 4 is full",
			stack:     getFilledStack([]int{10, 20, 30, 40}),
			expected2: expected2{result: false},
		},
		{
			name:      "Check if a stack of size 5 is full",
			stack:     getFilledStack([]int{10, 20, 30, 40, 50}),
			expected2: expected2{result: true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.stack.isFull()
			assertResult(t, tc.expected2.result, result)
		})
	}
}

func TestGetSize(t *testing.T) {
	testCases := []struct {
		name string
		stack
		expected
	}{
		{
			name:     "Get size from an empty stack",
			stack:    initStack(),
			expected: expected{size: -1},
		},
		{
			name:     "Get size from a stack of size 1",
			stack:    getFilledStack([]int{10}),
			expected: expected{size: 1},
		},
		{
			name:     "Get size from a stack of size 2",
			stack:    getFilledStack([]int{10, 20}),
			expected: expected{size: 2},
		},
		{
			name:     "Get size from a stack of size 3",
			stack:    getFilledStack([]int{10, 20, 30}),
			expected: expected{size: 3},
		},
		{
			name:     "Get size from a stack of size 4",
			stack:    getFilledStack([]int{10, 20, 30, 40}),
			expected: expected{size: 4},
		},
		{
			name:     "Get size from a stack of size 5",
			stack:    getFilledStack([]int{10, 20, 30, 40, 50}),
			expected: expected{size: 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := tc.stack.getSize()
			assertSize(t, tc.expected.size, size)
		})
	}
}

func assertSize(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("Expected size: %d got: %d", want, got)
	}
}

func assertContainer(t testing.TB, want, got [max]int) {
	t.Helper()
	if want != got {
		t.Errorf("Expected result: %v got: %v", want, got)
	}
}

func assertResult(t testing.TB, want, got bool) {
	t.Helper()
	if want != got {
		t.Errorf("Expected result: %t got: %t", want, got)
	}
}

func initStack() stack {
	return new()
}

func getFilledStack(values []int) stack {
	s := stack{}
	for i, v := range values {
		s.size = i
		s.container[i] = v
	}
	return s
}
