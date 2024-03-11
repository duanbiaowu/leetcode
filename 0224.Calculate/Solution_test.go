package leetcode

import (
	"testing"
)

func Test_calculate(t *testing.T) {
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
			args{"1 + 1"},
			2,
		},
		{
			"test-3",
			args{"2 - 1 + 2"},
			3,
		},
		{
			"test-4",
			args{"(1+(4+5+2)-3)+(6+8)"},
			23,
		},
		{
			"test-5",
			args{"(1+(4+5"},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
