package leetcode

import "testing"

func TestCheckPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
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
			args{"abc", "bca"},
			true,
		},
		{
			"test-3",
			args{"abc", "bcd"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
