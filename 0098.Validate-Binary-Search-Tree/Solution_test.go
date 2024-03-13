package leetcode

import "testing"

func Test_isValidBST(t *testing.T) {
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
				Val:   2,
				Right: &TreeNode{Val: 1},
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
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			true,
		},
		{
			"test-7",
			args{&TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
