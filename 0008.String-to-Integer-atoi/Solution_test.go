package leetcode

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test an empty string",
			args{""},
			0,
		},
		{
			"test a string that contains only spaces",
			args{"    "},
			0,
		},

		{
			"test a non-numeric string",
			args{"hello world"},
			0,
		},

		{
			"test a positive integer without leading spaces",
			args{"42"},
			42,
		},
		{
			"test a negative integer without leading spaces",
			args{"-42"},
			-42,
		},
		{
			"test a positive integer with leading spaces",
			args{"     42"},
			42,
		},
		{
			"test a negative integer with leading spaces",
			args{"     -42"},
			-42,
		},

		{
			"test a string with a leading symbol that a digit",
			args{"4193 with words"},
			4193,
		},
		{
			"test a string with a leading symbol that is not a digit",
			args{"words and 987"},
			0,
		},

		{
			"test the maximum positive integer",
			args{"2147483647"},
			2147483647,
		},
		{
			"test the maximum negative integer",
			args{"-2147483647"},
			-2147483647,
		},

		{
			"No1. test a positive integer that overflow after converting",
			args{"2147483648"},
			2147483647,
		},
		{
			"No2. test a positive integer that overflow after converting",
			args{"21474836460"},
			2147483647,
		},

		{
			"No1. test a negative integer that overflow after converting",
			args{"-2147483648"},
			-2147483648,
		},
		{
			"No2. test a negative integer that overflow after converting",
			args{"-91283472332"},
			-2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
