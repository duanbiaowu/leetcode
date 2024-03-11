package leetcode

import (
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
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
			args{"egg", "add"},
			true,
		},
		{
			"test-3",
			args{"bar", "foo"},
			false,
		},
		{
			"test-4",
			args{"baba", "badc"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
