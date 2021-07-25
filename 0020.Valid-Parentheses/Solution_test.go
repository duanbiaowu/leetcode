package leetcode

import "testing"

func Test_isValid(t *testing.T) {
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
			true,
		},
		{
			"test-2",
			args{"("},
			false,
		},
		{
			"test-3",
			args{"["},
			false,
		},
		{
			"test-4",
			args{"{"},
			false,
		},
		{
			"test-5",
			args{")"},
			false,
		},
		{
			"test-6",
			args{"}"},
			false,
		},
		{
			"test-7",
			args{"]"},
			false,
		},

		{
			"test-8",
			args{"()"},
			true,
		},
		{
			"test-9",
			args{"()[]{}"},
			true,
		},
		{
			"test-10",
			args{"(]"},
			false,
		},
		{
			"test-11",
			args{"([)]"},
			false,
		},
		{
			"test-12",
			args{"{[]}"},
			true,
		},
		{
			"test-13",
			args{"({[]})"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
