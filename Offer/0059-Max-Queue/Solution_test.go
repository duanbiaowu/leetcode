package leetcode

import "testing"

func TestMaxQueue_Max_value(t *testing.T) {
	queue := Constructor()

	queue.Push_back(1)
	if got := queue.Max_value(); got != 1 {
		t.Errorf("Max_value() = %v, want %v", got, 1)
	}

	queue.Push_back(0)
	queue.Push_back(2)
	if got := queue.Max_value(); got != 2 {
		t.Errorf("Max_value() = %v, want %v", got, 2)
	}

	queue.Pop_front()
	if got := queue.Max_value(); got != 2 {
		t.Errorf("Max_value() = %v, want %v", got, 2)
	}

	queue.Push_back(100)
	queue.Pop_front()
	queue.Pop_front()
	if got := queue.Max_value(); got != 100 {
		t.Errorf("Max_value() = %v, want %v", got, 100)
	}
}
