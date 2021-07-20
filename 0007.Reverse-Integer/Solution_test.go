package leetcode

import (
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test a positive integer of single digit",
			args{1},
			1,
		},
		{
			"test a negative integer of single digit",
			args{-1},
			-1,
		},
		{
			"test a positive integer of two digits",
			args{12},
			21,
		},
		{
			"test a negative integer of two digits",
			args{-89},
			-98,
		},
		{
			"test a positive integer of multi digits",
			args{12345},
			54321,
		},
		{
			"test a negative integer of multi digits",
			args{-6789},
			-9876,
		},
		{
			"test a positive integer that overflow after reversing",
			args{math.MaxInt32},
			0,
		},
		{
			"test a negative integer that overflow after reversing",
			args{math.MinInt32},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
