package leetcode

import "testing"

func Test_massage(t *testing.T) {
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
			args{[]int{1, 2, 3, 1}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := massage(tt.args.nums); got != tt.want {
				t.Errorf("massage() = %v, want %v", got, tt.want)
			}
		})
	}
}
