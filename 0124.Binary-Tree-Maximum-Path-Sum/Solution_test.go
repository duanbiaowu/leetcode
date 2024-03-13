package leetcode

import (
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{structures.GenerateTreeNodesBySlice([]int{0})},
			0,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			6,
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{-10, 9, 20, structures.NULL, structures.NULL, 15, 7})},
			42,
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{-3})},
			-3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum2(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
