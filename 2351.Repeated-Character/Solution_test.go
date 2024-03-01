package leetcode

import "testing"

func Test_repeatedCharacter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"test-1",
			args{"abccbaacz"},
			'c',
		},
		{
			"test-2",
			args{"abcdd"},
			'd',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedCharacter(tt.args.s); got != tt.want {
				t.Errorf("repeatedCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
