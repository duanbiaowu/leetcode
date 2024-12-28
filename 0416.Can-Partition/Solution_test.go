package leetcode

import (
	"testing"
)

func Test_canPartition(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want bool
	}{
		{
			"test-1",
			[]int{1, 1},
			true,
		},
		{
			"test-2",
			[]int{1},
			false,
		},
		{
			"test-3",
			[]int{1, 2},
			false,
		},
		{
			"test-4",
			[]int{1, 5, 11, 5},
			true,
		},
		// {
		// 	"test-5",
		// 	[]int{1, 2, 3, 5},
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
