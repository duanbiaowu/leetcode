package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
			args{[]int{2, 5, 6, 0, 0, 1, 2}, 0},
			true,
		},
		{
			"test-3",
			args{[]int{2, 5, 6, 0, 0, 1, 2}, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
