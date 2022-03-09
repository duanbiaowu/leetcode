package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			&TreeNode{Val: 1},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			structures.GenerateTreeNodesBySlice([]int{1, 3, 2}),
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{4, 2, 7, 1, 3, 6, 9})},
			structures.GenerateTreeNodesBySlice([]int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
