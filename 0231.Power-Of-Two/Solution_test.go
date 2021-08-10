package leetcode

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{0},
			false,
		},
		{
			"test-2",
			args{1},
			true,
		},
		{
			"test-3",
			args{2},
			true,
		},
		{
			"test-4",
			args{5},
			false,
		},
		{
			"test-5",
			args{1024},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
