package leetcode

import (
	"reflect"
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	tree1 := structures.GenerateTreeNodesBySlice([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4})
	tree2 := structures.GenerateTreeNodesBySlice([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4})
	tree3 := structures.GenerateTreeNodesBySlice([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4})

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{nil, nil, nil},
			nil,
		},
		{
			"test-2",
			args{tree1, tree1.Left, tree1.Right},
			tree1,
		},
		{
			"test-3",
			args{tree2, tree2.Left, tree2.Left.Right.Right},
			tree2.Left,
		},
		{
			"test-4",
			args{tree3, tree3.Right, tree3.Right.Right},
			tree3.Right,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
