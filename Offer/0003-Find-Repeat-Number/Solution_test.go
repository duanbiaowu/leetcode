package leetcode

import "testing"

func Test_findRepeatNumber(t *testing.T) {
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
			args{[]int{2, 3, 1, 0, 2, 5, 3}},
			2,
		},
		{
			"test-2",
			args{[]int{1, 2, 3, 4, 5, 6, 0}},
			-1,
		},
		{
			"test-3",
			args{[]int{3, 4, 2, 0, 0, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
