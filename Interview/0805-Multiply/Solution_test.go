package leetcode

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0, 0},
			0,
		},
		{
			"test-2",
			args{1, 2},
			2,
		},
		{
			"test-3",
			args{13, 12},
			156,
		},
		{
			"test-4",
			args{2, 3},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
