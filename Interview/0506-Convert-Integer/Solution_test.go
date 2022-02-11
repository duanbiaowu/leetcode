package leetcode

import "testing"

func Test_convertInteger(t *testing.T) {
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
			args{29, 15},
			2,
		},
		{
			"test-2",
			args{1, 2},
			2,
		},
		{
			"test-3",
			args{826966453, -729934991},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertInteger(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("convertInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
