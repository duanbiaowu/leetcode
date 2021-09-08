package leetcode

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			nil,
		},
		{
			"test-2",
			args{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}},
			[]string{
				"cat sand dog",
				"cats and dog",
			},
		},
		{
			"test-3",
			args{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}},
			[]string{
				"pine apple pen apple",
				"pine applepen apple",
				"pineapple pen apple",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
