package leetcode

import "testing"

func Test_replaceSpaces(t *testing.T) {
	type args struct {
		S      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{"Mr John Smith    ", 13},
			"Mr%20John%20Smith",
		},
		{
			"test-2",
			args{"               ", 5},
			"%20%20%20%20%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpaces(tt.args.S, tt.args.length); got != tt.want {
				t.Errorf("replaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
