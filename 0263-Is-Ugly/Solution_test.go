package leetcode

import "testing"

func Test_isUgly(t *testing.T) {
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
			args{6},
			true,
		},
		{
			"test-2",
			args{8},
			true,
		},
		{
			"test-3",
			args{14},
			false,
		},
		{
			"test-4",
			args{0},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
