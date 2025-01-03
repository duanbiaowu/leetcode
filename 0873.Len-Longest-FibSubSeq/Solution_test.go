package leetcode

import (
	"testing"
)

func Test_lenLongestFibSubseq(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"test-1",
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			5,
		},
		{
			"test-2",
			[]int{1, 2, 5, 6, 7},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenLongestFibSubseq(tt.args); got != tt.want {
				t.Errorf("lenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
