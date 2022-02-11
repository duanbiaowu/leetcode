package leetcode

import (
	"fmt"
	"testing"
)

func Test_exchangeBits(t *testing.T) {
	type args struct {
		num int
	}
	println((2 & 0x55555555) << 1)
	println((2 & 0xaaaaaaaa) >> 1)
	fmt.Printf("%b\n", 21)
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{2},
			1,
		},
		{
			"test-2",
			args{3},
			3,
		},
		{
			"test-3",
			args{21},
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchangeBits(tt.args.num); got != tt.want {
				t.Errorf("exchangeBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
