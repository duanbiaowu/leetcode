package leetcode

import "testing"

func Test_romanToInt(t *testing.T) {
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
			args{"III"},
			3,
		},
		{
			"test-2",
			args{"IV"},
			4,
		},
		{
			"test-3",
			args{"IX"},
			9,
		},
		{
			"test-4",
			args{"LVIII"},
			58,
		},
		{
			"test-5",
			args{"MCMXCIV"},
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
