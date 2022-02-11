package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}

	tree1 := structures.GenerateTreeNodesBySlice([]int{2, 1, 3})
	tree2 := structures.GenerateTreeNodesBySlice([]int{5, 3, 6, 2, 4, structures.NULL, structures.NULL, 1})

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{tree1, nil},
			nil,
		},
		{
			"test-3",
			args{tree1, tree1.Left},
			tree1,
		},
		{
			"test-4",
			args{tree2, tree2.Right},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
