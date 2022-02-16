package leetcode

import "testing"

func Test_getKthMagicNumber(t *testing.T) {
	type args struct {
		k int
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
			1,
		},
		{
			"test-3",
			args{2},
			3,
		},
		{
			"test-4",
			args{3},
			5,
		},
		{
			"test-5",
			args{4},
			7,
		},
		{
			"test-6",
			args{5},
			9,
		},
		{
			"test-7",
			args{9},
			27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthMagicNumber(tt.args.k); got != tt.want {
				t.Errorf("getKthMagicNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
