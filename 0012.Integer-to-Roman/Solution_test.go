package leetcode

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{3},
			"III",
		},
		{
			"test-2",
			args{4},
			"IV",
		},
		{
			"test-3",
			args{9},
			"IX",
		},
		{
			"test-4",
			args{58},
			"LVIII",
		},
		{
			"test-5",
			args{1994},
			"MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
