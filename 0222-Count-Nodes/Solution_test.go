package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_countNodes(t *testing.T) {
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
			0,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			1,
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			3,
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5, 6})},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
