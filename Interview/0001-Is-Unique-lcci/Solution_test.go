package leetcode

import "testing"

func Test_isUnique(t *testing.T) {
	type args struct {
		astr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{"leetcode"},
			false,
		},
		{
			"test-2",
			args{"abc"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnique(tt.args.astr); got != tt.want {
				t.Errorf("isUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
