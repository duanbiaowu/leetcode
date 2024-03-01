package leetcode

import (
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, 1},
			-1,
		},
		{
			"test-2",
			args{[]int{3, 1, 5, 4, 2}, 2},
			4,
		},
		{
			"test-3",
			args{[]int{3, 1, 5, 4, 2}, 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
