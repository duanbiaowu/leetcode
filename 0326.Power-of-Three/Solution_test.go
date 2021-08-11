package leetcode

import "testing"

func Test_isPowerOfThree(t *testing.T) {
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
			false,
		},
		{
			"test-4",
			args{3},
			true,
		},
		{
			"test-5",
			args{9},
			true,
		},
		{
			"test-6",
			args{18},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
