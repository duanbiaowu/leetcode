package leetcode

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"", ""},
			true,
		},
		{
			"test-2",
			args{"a", ""},
			false,
		},
		{
			"test-3",
			args{"", "a"},
			false,
		},
		{
			"test-4",
			args{"a", "b"},
			false,
		},
		{
			"test-5",
			args{"ab", "b"},
			false,
		},
		{
			"test-6",
			args{"a", "ab"},
			false,
		},
		{
			"test-7",
			args{"rat", "cat"},
			false,
		},
		{
			"test-8",
			args{"anagram", "nagaram"},
			true,
		},
		{
			"test-9",
			args{"aacc", "ccac"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
