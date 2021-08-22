package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_minDepth(t *testing.T) {
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
			1,
		},
		{
			"test-3",
			args{&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			}},
			2,
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{
				2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6,
			})},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
