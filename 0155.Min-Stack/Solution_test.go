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

func TestMinStack_GetMinOpt(t *testing.T) {
	stack := ConstructorOpt()

	stack.Push(1)
	stack.Push(2)

	if got := stack.Top(); got != 2 {
		t.Errorf("Top() = %v, want %v", got, 2)
	}
	if got := stack.GetMin(); got != 1 {
		t.Errorf("GetMin() = %v, want %v", got, 1)
	}

	stack.Pop()
	if got := stack.GetMin(); got != 1 {
		t.Errorf("GetMin() = %v, want %v", got, 1)
	}
	if got := stack.Top(); got != 1 {
		t.Log(stack.stack)
		t.Log(stack.min)
		t.Errorf("Top() = %v, want %v", got, 1)
	}
}
