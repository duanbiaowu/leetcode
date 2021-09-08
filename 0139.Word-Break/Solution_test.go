package leetcode

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"leetcode", []string{"leet", "code"}},
			true,
		},
		{
			"test-2",
			args{"applepenapple", []string{"apple", "pen"}},
			true,
		},
		{
			"test-3",
			args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
