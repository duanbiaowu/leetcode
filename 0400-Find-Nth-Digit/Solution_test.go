package leetcode

import "testing"

func Test_findNthDigit(t *testing.T) {
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
			args{3},
			3,
		},
		{
			"test-2",
			args{11},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
