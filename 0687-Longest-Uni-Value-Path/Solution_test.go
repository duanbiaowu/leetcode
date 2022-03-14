package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"testing"
)

func Test_longestUniValuePath(t *testing.T) {
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
			args{&TreeNode{Val: 0}},
			0,
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			0,
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{5, 4, 5, 1, 1, structures.NULL, 5})},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestUniValuePath(tt.args.root); got != tt.want {
				t.Errorf("longestUniValuePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
