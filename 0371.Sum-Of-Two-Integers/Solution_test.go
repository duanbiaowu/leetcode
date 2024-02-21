package leetcode

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{0, 0},
			0,
		},
		{
			"test-2",
			args{2, 3},
			5,
		},
		{
			"test-3",
			args{-2, -3},
			-5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
