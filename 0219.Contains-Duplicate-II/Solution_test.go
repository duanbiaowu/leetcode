package leetcode

import (
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil, 0},
			false,
		},
		{
			"test-2",
			args{[]int{1, 2, 3, 1}, 3},
			true,
		},
		{
			"test-3",
			args{[]int{1, 2, 3, 1, 2, 3}, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
