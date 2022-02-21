package leetcode

import "testing"

func Test_pileBox(t *testing.T) {
	type args struct {
		box [][]int
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
			args{[][]int{
				{1, 1, 1},
				{2, 2, 2},
				{3, 3, 3},
			}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pileBox(tt.args.box); got != tt.want {
				t.Errorf("pileBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
