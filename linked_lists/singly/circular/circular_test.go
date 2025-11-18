package circular

import (
	"slices"
	"testing"
)

var (
	val    = 50
	values = []int{10, 20, 30}
)

func TestInsert(t *testing.T) {
	t.Run("Insert in empty list with invalid index", func(t *testing.T) {
		l := list{}
		l.insert(1, val)
		assertSize(t, uint(0), l.size)
		assertValues(t, []int{}, getValues(l))
		if got := l.head; got != nil {
			t.Errorf("Expected result: %v, got: %p", nil, got)
		}
	})
	t.Run("Insert in empty list with valid index", func(t *testing.T) {
		l := list{}
		l.insert(0, val)
		assertSize(t, uint(1), l.size)
		assertValues(t, []int{val}, getValues(l))
		assertPointer(t, l.head, getPtrFromLastNode(l))
	})
	t.Run("Insert into the filled list at the first index", func(t *testing.T) {
		l := fillList(values)
		l.insert(0, val)
		assertSize(t, uint(4), l.size)
		assertValues(t, []int{50, 10, 20, 30}, getValues(l))
		assertPointer(t, l.head, getPtrFromLastNode(l))
	})
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
		t.Fatalf("Expected result: %v got: %v", want, got)
	}
}

func assertPointer(t testing.TB, want, got *node) {
	if want != got {
		t.Errorf("Expected result: [ptr: %p val: %[1]v] got: [ptr: %p val: %[2]v]", want, got)
	}
}

func getValues(l list) []int {
	res := make([]int, 0, l.size)
	for i, curr := uint(0), l.head; i < l.size; i, curr = i+1, curr.next {
		res = append(res, curr.value)
	}
	return res
}

func getPtrFromLastNode(l list) *node {
	curr := l.head
	for i := uint(1); i < l.size; i++ {
		curr = curr.next
	}
	return curr.next
}

func fillList(values []int) list {
	l := list{}
	for i := range values {
		l.insert(uint(i), values[i])
	}
	return l
}
