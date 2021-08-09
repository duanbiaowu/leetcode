package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"test-1",
		//	args{nil},
		//	0,
		//},
		//{
		//	"test-2",
		//	args{&TreeNode{}},
		//	1,
		//},
		//{
		//	"test-3",
		//	args{&TreeNode{
		//		Val: 1,
		//		Left: &TreeNode{Val: 2},
		//	}},
		//	2,
		//},
		//{
		//	"test-4",
		//	args{&TreeNode{
		//		Val: 1,
		//		Right: &TreeNode{Val: 3},
		//	}},
		//	2,
		//},
		//{
		//	"test-5",
		//	args{&TreeNode{
		//		Val: 1,
		//		Left: &TreeNode{Val: 2},
		//		Right: &TreeNode{Val: 3},
		//	}},
		//	2,
		//},
		//{
		//	"test-6",
		//	args{&TreeNode{
		//		Val: 1,
		//		Left: &TreeNode{
		//			Val: 2,
		//			Left: &TreeNode{Val: 4},
		//		},
		//		Right: &TreeNode{Val: 3},
		//	}},
		//	3,
		//},
		//{
		//	"test-7",
		//	args{&TreeNode{
		//		Val: 1,
		//		Left: &TreeNode{Val: 2},
		//		Right: &TreeNode{
		//			Val: 3,
		//			Right: &TreeNode{Val: 7},
		//		},
		//	}},
		//	3,
		//},
		{
			"test-8",
			args{&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 7},
				},
			}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
