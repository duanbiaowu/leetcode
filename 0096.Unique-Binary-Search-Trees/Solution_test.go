package leetcode

import "testing"

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{1},
			1,
		},
		{
			"test-2",
			args{2},
			2,
		},
		{
			"test-3",
			args{3},
			5,
		},
		{
			"test-4",
			args{4},
			14,
		},
		{
			"test-5",
			args{5},
			42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
