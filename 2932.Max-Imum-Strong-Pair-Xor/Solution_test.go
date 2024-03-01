package leetcode

import "testing"

func Test_maximumStrongPairXor(t *testing.T) {
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
			args{[]int{1, 2, 3, 4, 5}},
			7,
		},
		{
			"test-3",
			args{[]int{10, 100}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumStrongPairXor(tt.args.nums); got != tt.want {
				t.Errorf("maximumStrongPairXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
