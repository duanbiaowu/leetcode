package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil},
			[][]int{},
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			[][]int{
				{1},
			},
		},
		{
			"test-3",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			[][]int{
				{2, 3},
				{1},
			},
		},
		{
			"test-4",
			args{&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			}},
			[][]int{
				{4, 5},
				{2, 3},
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
