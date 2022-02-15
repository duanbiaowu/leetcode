package leetcode

import (
	"reflect"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{"", "", nil},
			[]string{""},
		},
		{
			"test-2",
			args{"", "a", nil},
			nil,
		},
		{
			"test-3",
			args{"a", "", nil},
			nil,
		},
		{
			"test-4",
			args{"a", "", []string{"c"}},
			nil,
		},
		{
			"test-5",
			args{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			[]string{"hit", "hot", "dot", "dog", "log", "cog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
