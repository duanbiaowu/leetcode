package leetcode

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
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
			0,
		},
		{
			"test-4",
			args{[]int{1, 1}},
			0,
		},
		{
			"test-5",
			args{[]int{1, 2}},
			1,
		},
		{
			"test-6",
			args{[]int{2, 5, 1, 2, 7}},
			6,
		},
		{
			"test-7",
			args{[]int{7, 1, 5, 3, 6, 4}},
			5,
		},
		{
			"test-8",
			args{[]int{7, 6, 4, 3, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
