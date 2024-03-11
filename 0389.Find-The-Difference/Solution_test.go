package leetcode

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"test-1",
			args{"", ""},
			byte(0),
		},
		{
			"test-2",
			args{"abcd", "abcde"},
			'e',
		},
		{
			"test-3",
			args{"", "y"},
			'y',
		},
		{
			"test-4",
			args{"a", "aa"},
			'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifferenceOpt2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
