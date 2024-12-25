package leetcode

import (
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{""},
			"",
		},
		{
			"test-2",
			args{"3[a]2[bc]"},
			"aaabcbc",
		},
		{
			"test-3",
			args{"2[abc]3[cd]ef"},
			"abcabccdcdcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
