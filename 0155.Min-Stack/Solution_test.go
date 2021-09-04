package leetcode

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	if got := stack.GetMin(); got != 1 {
		t.Errorf("GetMin() = %v, want %v", got, 1)
	}

	stack.Push(0)
	stack.Push(2)
	if got := stack.GetMin(); got != 0 {
		t.Errorf("GetMin() = %v, want %v", got, 0)
	}

	stack.Pop()
	if got := stack.GetMin(); got != 0 {
		t.Errorf("GetMin() = %v, want %v", got, 0)
	}

	stack.Push(3)
	if got := stack.GetMin(); got != 0 {
		t.Errorf("GetMin() = %v, want %v", got, 0)
	}

	stack.Pop()
	stack.Pop()
	if got := stack.GetMin(); got != 1 {
		t.Errorf("GetMin() = %v, want %v", got, 1)
	}
}
