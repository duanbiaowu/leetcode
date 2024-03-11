package leetcode

import "testing"

func Test_hasTrailingZeros(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil},
			false,
		},
		{
			"test-2",
			args{[]int{0, 0}},
			true,
		},
		{
			"test-3",
			args{[]int{1}},
			false,
		},
		{
			"test-4",
			args{[]int{1, 2, 3, 4, 5}},
			true,
		},
		{
			"test-5",
			args{[]int{1, 3, 5, 7, 9}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasTrailingZerosOpt(tt.args.nums); got != tt.want {
				t.Errorf("hasTrailingZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
