package leetcode

import "testing"

func Test_reversePairs(t *testing.T) {
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
			0,
		},
		{
			"test-3",
			args{[]int{1, 2, 3}},
			0,
		},
		{
			"test-4",
			args{[]int{7, 5, 6, 4}},
			5,
		},
		{
			"test-5",
			args{[]int{4, 5, 6, 7}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
