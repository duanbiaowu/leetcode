package leetcode

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"a", ""},
			false,
		},
		{
			"test-2",
			args{"a", "b"},
			false,
		},
		{
			"test-3",
			args{"aa", "ab"},
			false,
		},
		{
			"test-4",
			args{"aa", "aab"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
