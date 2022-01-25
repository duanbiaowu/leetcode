package leetcode

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"pale", "ple"},
			true,
		},
		{
			"test-2",
			args{"pal", "pale"},
			true,
		},
		{
			"test-3",
			args{"pales", "pal"},
			false,
		},
		{
			"test-4",
			args{"abcdefg", "abceefg"},
			true,
		},
		{
			"test-5",
			args{"intention", "execution"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
