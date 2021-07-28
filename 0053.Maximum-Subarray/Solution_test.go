package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{}},
			0,
		},
		{
			"test-2",
			args{[]int{0}},
			0,
		},
		{
			"test-3",
			args{[]int{1}},
			1,
		},
		{
			"test-4",
			args{[]int{1, -1}},
			1,
		},
		{
			"test-5",
			args{[]int{1, -1, 3}},
			3,
		},
		{
			"test-6",
			args{[]int{2, 1, -2, 3}},
			4,
		},
		{
			"test-7",
			args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
