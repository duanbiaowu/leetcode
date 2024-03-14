package leetcode

import (
	"testing"
)

func Test_getMinimumDifference(t *testing.T) {
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
			args{&TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			}},
			1,
		},
		{
			"test-2",
			args{&TreeNode{
				Val:   10,
				Right: &TreeNode{Val: 20},
			}},
			10,
		},
		{
			"test-3",
			args{&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			1,
		},
		{
			"test-4",
			args{&TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 6},
			}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDifference(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
