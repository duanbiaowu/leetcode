package leetcode

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[]int{1, 2, 5}, 11},
			3,
		},
		{
			"test-2",
			args{[]int{2}, 3},
			-1,
		},
		{
			"test-3",
			args{[]int{0}, 0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
