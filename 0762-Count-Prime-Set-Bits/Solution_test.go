package leetcode

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{6, 10},
			4,
		},
		{
			"test-2",
			args{10, 15},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimeSetBits(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
