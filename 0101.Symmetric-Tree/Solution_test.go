package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil},
			true,
		},
		{
			"test-2",
			args{&TreeNode{}},
			true,
		},
		{
			"test-3",
			args{&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			}},
			false,
		},
		{
			"test-4",
			args{&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			}},
			false,
		},
		{
			"test-5",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			false,
		},
		{
			"test-6",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 2},
			}},
			true,
		},
		{
			"test-6",
			args{structures.GenerateTreeNodesBySlice([]int{
				1, 2, 2, 3, 4, 4, 3,
			})},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
