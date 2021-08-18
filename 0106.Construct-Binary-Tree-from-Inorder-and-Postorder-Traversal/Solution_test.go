package leetcode

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{[]int{}, []int{}},
			nil,
		},
		{
			"test-2",
			args{[]int{1}, []int{1}},
			&TreeNode{Val: 1},
		},
		{
			"test-3",
			args{[]int{2, 1}, []int{2, 1}},
			&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
		},
		{
			"test-4",
			args{[]int{1, 2}, []int{2, 1}},
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			"test-5",
			args{[]int{2, 1, 3}, []int{2, 3, 1}},
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			"test-6",
			args{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
