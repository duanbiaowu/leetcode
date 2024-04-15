package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
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
			args{nil},
			0,
		},
		{
			"test-2",
			args{[]int{1}},
			1,
		},
		{
			"test-3",
			args{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			4,
		},
		{
			"test-4",
			args{[]int{0, 1, 0, 3, 2, 3}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISBruteForce(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
