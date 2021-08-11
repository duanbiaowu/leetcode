package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	var (
		treeNode1 = &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}

		treeNodes2 = structures.GenerateTreeNodesBySlice([]int{
			6, 2, 8, 0, 4, 7, 9, structures.NULL, structures.NULL, 3, 5,
		})
	)

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
			args{&TreeNode{}, nil, nil},
			nil,
		},
		{
			"test-3",
			args{&TreeNode{}, &TreeNode{}, nil},
			nil,
		},
		{
			"test-4",
			args{&TreeNode{}, nil, &TreeNode{}},
			nil,
		},

		{
			"test-5",
			args{&TreeNode{}, &TreeNode{}, &TreeNode{}},
			nil,
		},
		{
			"test-6",
			args{treeNode1, treeNode1.Left, treeNode1.Right},
			treeNode1,
		},

		{
			"test-7",
			args{treeNodes2,
				treeNodes2.Left,
				treeNodes2.Right,
			},
			treeNodes2,
		},
		{
			"test-8",
			args{treeNodes2,
				treeNodes2.Left,
				treeNodes2.Left.Right,
			},
			treeNodes2.Left,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestorUsingPath(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
