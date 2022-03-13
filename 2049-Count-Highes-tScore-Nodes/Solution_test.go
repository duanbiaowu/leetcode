package leetcode

import "testing"

func Test_countHighestScoreNodes(t *testing.T) {
	type args struct {
		parents []int
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
			args{[]int{-1, 2, 0, 2, 0}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHighestScoreNodes(tt.args.parents); got != tt.want {
				t.Errorf("countHighestScoreNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
