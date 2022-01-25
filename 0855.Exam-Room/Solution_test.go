package leetcode

import (
	"testing"
)

func TestExamRoom(t *testing.T) {
	exam := Constructor(10)

	if got := exam.Seat(); got != 0 {
		t.Errorf("Seat() = %v, want %v", got, 0)
	}
	if got := exam.Seat(); got != 9 {
		t.Errorf("Seat() = %v, want %v", got, 9)
	}
	if got := exam.Seat(); got != 4 {
		t.Errorf("Seat() = %v, want %v", got, 4)
	}
	if got := exam.Seat(); got != 2 {
		t.Errorf("Seat() = %v, want %v", got, 2)
	}

	exam.Leave(4)

	if got := exam.Seat(); got != 5 {
		t.Errorf("Seat() = %v, want %v", got, 5)
	}
}
