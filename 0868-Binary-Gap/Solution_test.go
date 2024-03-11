package leetcode

import "testing"

func Test_binaryGap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0},
			0,
		},
		{
			"test-2",
			args{1},
			0,
		},
		{
			"test-3",
			args{22},
			2,
		},
		{
			"test-4",
			args{8},
			0,
		},
		{
			"test-5",
			args{5},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.n); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
