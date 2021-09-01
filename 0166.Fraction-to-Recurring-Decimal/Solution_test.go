package leetcode

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{1, 0},
			"",
		},
		{
			"test-2",
			args{0, 1},
			"0",
		},
		{
			"test-3",
			args{1, -2},
			"-0.5",
		},
		{
			"test-4",
			args{-1, 2},
			"-0.5",
		},
		{
			"test-5",
			args{1, 2},
			"0.5",
		},
		{
			"test-6",
			args{-2, 3},
			"-0.(6)",
		},
		{
			"test-7",
			args{2, -3},
			"-0.(6)",
		},
		{
			"test-8",
			args{2, 3},
			"0.(6)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
