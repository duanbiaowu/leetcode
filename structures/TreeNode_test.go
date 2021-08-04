package structures

import (
	"reflect"
	"testing"
)

func TestGenerateTreeNodesByArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test null",
			args{[]int{}},
			nil,
		},
		{
			"test a tree has only the root node",
			args{[]int{1}},
			&TreeNode{Val: 1},
		},
		{
			"test a tree has only the root node and a left node of the root node",
			args{[]int{1, 2}},
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: nil,
			},
		},
		{
			"test a tree has only the root node and a right node of the root node",
			args{[]int{1, NULL, 3}},
			&TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 3},
			},
		},
		{
			"test a tree has only the root node and two nodes left and right of the root node",
			args{[]int{1, 2, 3}},
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			"test a tree has multiple nodes-1",
			args{[]int{1, 2, 3, 4, NULL, 6, NULL}},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
		},
		{
			"test a tree has multiple nodes-2",
			args{[]int{1, 2, 3, 4, 5, NULL, NULL, 8, 9}},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 9},
					},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			"test a tree has multiple nodes-3",
			args{[]int{1, 2, 3, NULL, NULL, 4, 5, NULL, NULL, 6, 7}},
			&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateTreeNodesBySlice(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTreeNodesBySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
