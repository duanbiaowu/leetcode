package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
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
			args{nil},
			0,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			0,
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5})},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
