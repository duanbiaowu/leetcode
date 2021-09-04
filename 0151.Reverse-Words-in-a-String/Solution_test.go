package leetcode

import "testing"

func Test_reverseWords(t *testing.T) {
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
			args{"a"},
			"a",
		},
		{
			"test-3",
			args{"abc"},
			"abc",
		},
		{
			"test-4",
			args{"the sky is blue"},
			"blue is sky the",
		},
		{
			"test-5",
			args{"  hello world  "},
			"world hello",
		},
		{
			"test-6",
			args{"a good   example"},
			"example good a",
		},
		{
			"test-7",
			args{"  Bob    Loves  Alice   "},
			"Alice Loves Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
