package leetcode

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
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
			true,
		},
		{
			"test-2",
			args{5},
			true,
		},
		{
			"test-3",
			args{7},
			false,
		},
		{
			"test-4",
			args{3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBitsOpt(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
