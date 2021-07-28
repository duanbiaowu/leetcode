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
		want int
	}{
		{
			"test-1",
			args{[]int{}, 0},
			-1,
		},
		{
			"test-2",
			args{[]int{0}, 0},
			0,
		},
		{
			"test-3",
			args{[]int{-1, 0}, 0},
			1,
		},
		{
			"test-4",
			args{[]int{-1, 0, 1}, 0},
			1,
		},
		{
			"test-5",
			args{[]int{-1, 0, 1}, 2},
			-1,
		},
		{
			"test-6",
			args{[]int{-1, 0, 3, 5, 9, 12}, 9},
			4,
		},
		{
			"test-7",
			args{[]int{-1, 0, 3, 5, 9, 12}, 2},
			-1,
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
