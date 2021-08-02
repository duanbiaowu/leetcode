package leetcode

import "testing"

func Test_findSecondMinimumValue(t *testing.T) {
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
			args{nil},
			-1,
		},
		{
			"test-2",
			args{&TreeNode{Val: 0}},
			-1,
		},
		{
			"test-3",
			args{&TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 0},
			}},
			-1,
		},
		{
			"test-4",
			args{&TreeNode{
				Val:   0,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 0},
			}},
			1,
		},
		{
			"test-5",
			args{&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 0},
			}},
			1,
		},
		{
			"test-6",
			args{&TreeNode{
				Val:  0,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			}},
			1,
		},
		{
			"test-7",
			args{&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			}},
			1,
		},
		{
			"test-7",
			args{&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 2},
				},
			}},
			2,
		},
		{
			"test-8",
			args{&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 5},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 10},
						Right: &TreeNode{Val: 7},
					},
				},
			}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
