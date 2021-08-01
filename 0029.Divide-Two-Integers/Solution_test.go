package leetcode

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0, 1},
			0,
		},
		{
			"test-2",
			args{1, 1},
			1,
		},
		{
			"test-3",
			args{1, -1},
			-1,
		},
		{
			"test-3",
			args{-1, 1},
			-1,
		},
		{
			"test-4",
			args{-1, -1},
			1,
		},
		{
			"test-5",
			args{100, 10},
			10,
		},
		{
			"test-5.1",
			args{10, 100},
			0,
		},
		{
			"test-6",
			args{10, 3},
			3,
		},
		{
			"test-7",
			args{2147483647, 3},
			715827882,
		},
		{
			"test-8",
			args{math.MinInt32, 1},
			math.MinInt32,
		},
		{
			"test-9",
			args{math.MinInt32, -1},
			math.MaxInt32,
		},
		{
			"test-10",
			args{math.MinInt32 - 1, -1},
			math.MaxInt32,
		},
		{
			"test-11",
			args{math.MinInt32 + 1, -1},
			math.MaxInt32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
