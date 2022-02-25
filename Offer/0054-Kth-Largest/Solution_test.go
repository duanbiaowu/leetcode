package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_kthLargest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, 0},
			-1,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}, 1},
			1,
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{3, 1, 4, structures.NULL, 2}), 1},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
