package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
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
			args{&TreeNode{}, nil},
			&TreeNode{},
		},
		{
			"test-3",
			args{nil, &TreeNode{}},
			&TreeNode{},
		},
		{
			"test-4",
			args{&TreeNode{}, &TreeNode{}},
			&TreeNode{},
		},
		{
			"test-5",
			args{&TreeNode{Val: 1}, &TreeNode{Val: 2}},
			&TreeNode{Val: 3},
		},
		{
			"test-6",
			args{&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 3},
			}, &TreeNode{Val: 2}},
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 3},
			},
		},
		{
			"test-7",
			args{&TreeNode{Val: 1},
				&TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
			},
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 4},
			},
		},
		{
			"test-8",
			args{&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			}, &TreeNode{Val: 2}},
			&TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 3},
			},
		},
		{
			"test-9",
			args{&TreeNode{Val: 1},
				&TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
			},
			&TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 4},
			},
		},
		{
			"test-10",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
				&TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 6},
				},
			},
			&TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
