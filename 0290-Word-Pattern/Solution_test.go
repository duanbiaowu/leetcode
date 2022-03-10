package leetcode

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"abba", "dog cat cat fish"},
			false,
		},
		{
			"test-2",
			args{"abba", "dog cat cat dog"},
			true,
		},
		{
			"test-3",
			args{"aaaa", "dog cat cat dog"},
			false,
		},
		{
			"test-4",
			args{"aaaa", "dog dog dog dog"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
