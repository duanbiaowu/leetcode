package leetcode

import (
	"testing"
)

func Test_toHex(t *testing.T) {
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
			args{0},
			"0",
		},
		{
			"test-2",
			args{26},
			"1a",
		},
		{
			"test-3",
			args{-1},
			"ffffffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
