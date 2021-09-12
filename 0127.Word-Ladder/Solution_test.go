package leetcode

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{},
			1,
		},
		{
			"test-2",
			args{"hit", "cog", []string{
				"hot", "dot", "dog", "lot", "log", "cog",
			}},
			5,
		},
		{
			"test-3",
			args{"hit", "cog", []string{
				"hot", "dot", "dog", "lot", "log",
			}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
