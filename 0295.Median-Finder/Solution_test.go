package leetcode

import (
	"testing"
)

func TestMedianFinder(t *testing.T) {
	mf := Constructor()

	mf.AddNum(1)

	if val := mf.FindMedian(); val != 1 {
		t.Errorf("maximalRectangle() = %v, want %v", val, 1)
	}

	mf.AddNum(3)
	if val := mf.FindMedian(); val != 2 {
		t.Errorf("maximalRectangle() = %v, want %v", val, 2)
	}

	mf.AddNum(2)
	if val := mf.FindMedian(); val != 2 {
		t.Errorf("maximalRectangle() = %v, want %v", val, 2)
	}

	mf.AddNum(99)
	mf.AddNum(99)
	mf.AddNum(99)
	mf.AddNum(100)
	mf.AddNum(100)
	mf.AddNum(100)

	if val := mf.FindMedian(); val != 99 {
		t.Errorf("maximalRectangle() = %v, want %v", val, 99)
	}
}
