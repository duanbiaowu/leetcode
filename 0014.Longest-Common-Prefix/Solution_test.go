package leetcode

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"No1. test no common prefix",
			args{[]string{" ", ""}},
			"",
		},
		{
			"No2. test no common prefix",
			args{[]string{" ", "hello world"}},
			"",
		},
		{
			"No3. test no common prefix",
			args{[]string{"go", "hello", "world"}},
			"",
		},

		{
			"No4. test common prefix is only one character",
			args{[]string{"bug", "banana", "bus"}},
			"b",
		},
		{
			"No5. test common prefix is only one character",
			args{[]string{"after", "a", "abs"}},
			"a",
		},
		{
			"No6. test common prefix is only one character",
			args{[]string{"c"}},
			"c",
		},

		{
			"No7. test common prefix contains multiple characters",
			args{[]string{"flower", "flow", "flight"}},
			"fl",
		},
		{
			"No8. test common prefix contains multiple characters",
			args{[]string{"go", "golang", "gopher"}},
			"go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
