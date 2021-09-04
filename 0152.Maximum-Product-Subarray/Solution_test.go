package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
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
			args{[]int{-1}},
			-1,
		},
		{
			"test-4",
			args{[]int{-1, -2}},
			2,
		},
		{
			"test-5",
			args{[]int{3, -1, 4}},
			4,
		},
		{
			"test-6",
			args{[]int{2, 3, -2, 4}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
