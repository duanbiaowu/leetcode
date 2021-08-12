package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil, nil},
			true,
		},
		{
			"test-2",
			args{&TreeNode{}, &TreeNode{}},
			true,
		},
		{
			"test-3",
			args{&TreeNode{}, nil},
			false,
		},
		{
			"test-4",
			args{nil, &TreeNode{}},
			false,
		},
		{
			"test-5",
			args{&TreeNode{Val: 1}, &TreeNode{Val: 2}},
			false,
		},
		{
			"test-6",
			args{
				structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5, 6, 7}),
				structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5, 6, 7}),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
