package leetcode

import (
	"reflect"
	"testing"

	"github.com/duanbiaowu/leetcode/structures"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{nil},
			[][]int{},
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			[][]int{
				{1},
			},
		},
		{
			"test-3",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			}},
			[][]int{
				{1},
				{3, 2},
			},
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5, 6, 7})},
			[][]int{
				{1},
				{3, 2},
				{4, 5, 6, 7},
			},
		},
		{
			"test-5",
			args{structures.GenerateTreeNodesBySlice([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})},
			[][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
