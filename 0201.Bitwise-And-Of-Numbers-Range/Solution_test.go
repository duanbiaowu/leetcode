package leetcode

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{5, 7},
			4,
		},
		{
			"test-2",
			args{0, 0},
			0,
		},
		{
			"test-3",
			args{1, 2147483647},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
