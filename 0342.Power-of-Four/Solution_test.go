package leetcode

import "testing"

func Test_isPowerOfFour(t *testing.T) {
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
			false,
		},
		{
			"test-5",
			args{4},
			true,
		},
		{
			"test-6",
			args{16},
			true,
		},
		{
			"test-7",
			args{32},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
