package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
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
			0,
		},
		{
			"test-2",
			args{[]int{1}, 0},
			0,
		},
		{
			"test-3",
			args{[]int{1}, 2},
			1,
		},

		{
			"test-4",
			args{[]int{1, 3, 5, 6}, 5},
			2,
		},
		{
			"test-5",
			args{[]int{1, 3, 5, 6}, 2},
			1,
		},
		{
			"test-6",
			args{[]int{1, 3, 5, 6}, 7},
			4,
		},
		{
			"test-7",
			args{[]int{1, 3, 5, 6}, 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
