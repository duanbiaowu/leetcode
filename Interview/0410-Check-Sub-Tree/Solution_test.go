package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_checkSubTree(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}

	tree1 := structures.GenerateTreeNodesBySlice([]int{1, 2, 3})

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{tree1, tree1.Left},
			true,
		},
		{
			"test-2",
			args{
				structures.GenerateTreeNodesBySlice([]int{1, structures.NULL, 2, 4}),
				structures.GenerateTreeNodesBySlice([]int{3, 2}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubTree(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("checkSubTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
