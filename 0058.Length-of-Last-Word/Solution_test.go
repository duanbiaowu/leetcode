package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{""},
			0,
		},
		{
			"test-2",
			args{"a"},
			1,
		},
		{
			"test-3",
			args{"Hello World"},
			5,
		},
		{
			"test-4",
			args{"   fly me   to   the moon  "},
			4,
		},
		{
			"test-5",
			args{"luffy is still joyboy"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
