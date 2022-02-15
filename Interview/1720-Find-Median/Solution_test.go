package leetcode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	obj.AddNum(2)

	if got := obj.FindMedian(); got != 2 {
		t.Errorf("findLadders() = %v, want %v", got, 2)
	}

	obj.AddNum(3)
	if got := obj.FindMedian(); got != 2.5 {
		t.Errorf("findLadders() = %v, want %v", got, 2.5)
	}

	obj.AddNum(4)
	if got := obj.FindMedian(); got != 3 {
		t.Errorf("findLadders() = %v, want %v", got, 3)
	}

	obj.AddNum(3)
	if got := obj.FindMedian(); got != 3 {
		t.Errorf("findLadders() = %v, want %v", got, 3)
	}

	obj.AddNum(5)
	if got := obj.FindMedian(); got != 3 {
		t.Errorf("findLadders() = %v, want %v", got, 3)
	}

	obj.AddNum(6)
	if got := obj.FindMedian(); got != 3.5 {
		t.Errorf("findLadders() = %v, want %v", got, 3.5)
	}
}
