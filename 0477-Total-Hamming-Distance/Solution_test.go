package leetcode

import "testing"

func Test_totalHammingDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil},
			0,
		},
		{
			"test-2",
			args{[]int{0}},
			0,
		},
		{
			"test-3",
			args{[]int{1}},
			0,
		},
		{
			"test-4",
			args{[]int{4, 14, 2}},
			6,
		},
		{
			"test-5",
			args{[]int{4, 14, 4}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalHammingDistanceOpt(tt.args.nums); got != tt.want {
				t.Errorf("totalHammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
