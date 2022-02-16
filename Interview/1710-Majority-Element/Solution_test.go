package leetcode

import "testing"

func Test_majorityElement(t *testing.T) {
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
			-1,
		},
		{
			"test-2",
			args{[]int{1, 2, 5, 9, 5, 9, 5, 5, 5}},
			5,
		},
		{
			"test-3",
			args{[]int{2, 2, 1, 1, 1, 2, 2}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
