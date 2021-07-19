package leetcode

import (
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test an invalid numRow",
			args{"hello world", -1},
			"hello world",
		},
		{
			"test a 0 numRow",
			args{"hello world", 0},
			"hello world",
		},
		{
			"test an empty string",
			args{"", 0},
			"",
		},
		{
			"test a single character string",
			args{"a", 0},
			"a",
		},
		{
			"No.1 test a string that contains multiple characters",
			args{"PAYPALISHIRING", 3},
			"PAHNAPLSIIGYIR",
		},
		{
			"No.2 test a string that contains multiple characters",
			args{"PAYPALISHIRING", 4},
			"PINALSIGYAHRPI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}