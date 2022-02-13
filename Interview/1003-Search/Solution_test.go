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
			args{nil, 0},
			-1,
		},
		{
			"test-2",
			args{[]int{0}, 0},
			0,
		},
		{
			"test-3",
			args{[]int{0}, 1},
			-1,
		},
		{
			"test-4",
			args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5},
			8,
		},
		{
			"test-5",
			args{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 11},
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
