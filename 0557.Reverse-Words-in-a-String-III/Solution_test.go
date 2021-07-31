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
			args{" a a"},
			" a a",
		},
		{
			"test-4",
			args{"a b"},
			"a b",
		},
		{
			"test-5",
			args{"Let's take LeetCode contest"},
			"s'teL ekat edoCteeL tsetnoc",
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
