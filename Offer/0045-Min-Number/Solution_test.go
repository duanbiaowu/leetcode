package leetcode

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{[]int{0}},
			"0",
		},
		{
			"test-2",
			args{[]int{1}},
			"1",
		},
		{
			"test-3",
			args{[]int{10, 2}},
			"102",
		},
		{
			"test-4",
			args{[]int{3, 30, 34, 5, 9}},
			"3033459",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
