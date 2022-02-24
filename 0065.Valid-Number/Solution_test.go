package leetcode

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{""},
			false,
		},
		{
			"test-2",
			args{"0"},
			true,
		},
		{
			"test-3",
			args{"e"},
			false,
		},
		{
			"test-4",
			args{"."},
			false,
		},
		{
			"test-5",
			args{".1"},
			true,
		},
		{
			"test-6",
			args{"1 "},
			true,
		},
		{
			"test-7",
			args{" -54.53061"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
