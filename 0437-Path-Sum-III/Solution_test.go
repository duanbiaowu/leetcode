package leetcode

import (
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, 1},
			0,
		},
		{
			"test-2",
			args{&TreeNode{Val: 0}, 1},
			0,
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{10, 5, -3, 3, 2, structures.NULL, 11, 3, -2, structures.NULL, 1}), 8},
			3,
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, structures.NULL, 2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5}), 3},
			2,
		},
		{
			"test-5",
			args{structures.GenerateTreeNodesBySlice([]int{1, -2, -3}), -1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
