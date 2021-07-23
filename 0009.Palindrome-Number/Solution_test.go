package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"No1. test a positive integer of single digit",
			args{0},
			true,
		},
		{
			"No2. test a positive integer of single digit",
			args{1},
			true,
		},
		{
			"test a negative integer of single digit",
			args{-1},
			false,
		},
		{
			"test a positive integer of two digits",
			args{22},
			true,
		},
		{
			"test a negative integer of two digits",
			args{-87},
			false,
		},
		{
			"No1. test a positive integer of multi digits",
			args{100},
			false,
		},
		{
			"No2. test a positive integer of multi digits",
			args{123454321},
			true,
		},
		{
			"test a positive integer of multi digits",
			args{-78987},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
