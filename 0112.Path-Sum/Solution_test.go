package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil, 0},
			false,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}, 0},
			false,
		},
		{
			"test-3",
			args{&TreeNode{Val: 0}, 0},
			true,
		},
		{
			"test-4",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}, 5},
			false,
		},
		{
			"test-4",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}, 1},
			false,
		},
		{
			"test-6",
			args{structures.GenerateTreeNodesBySlice([]int{
				5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, structures.NULL, 1,
			}), 22},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
