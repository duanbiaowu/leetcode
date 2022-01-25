package leetcode

import "testing"

func Test_canMeasureWater(t *testing.T) {
	type args struct {
		jug1Capacity   int
		jug2Capacity   int
		targetCapacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{3, 5, 4},
			true,
		},
		{
			"test-2",
			args{2, 6, 5},
			false,
		},
		{
			"test-3",
			args{3, 2, 6},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWater(tt.args.jug1Capacity, tt.args.jug2Capacity, tt.args.targetCapacity); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
