package leetcode

import "testing"

func Test_isPossible(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{[]int{1, 2, 3, 3, 4, 5}},
			true,
		},
		{
			"test-2",
			args{[]int{1, 2, 3, 3, 4, 4, 5, 5}},
			true,
		},
		{
			"test-3",
			args{[]int{1, 2, 3, 4, 4, 5}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.nums); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
