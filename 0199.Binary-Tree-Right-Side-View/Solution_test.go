package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			[]int{},
		},
		{
			"test-2",
			args{&TreeNode{Val: 1}},
			[]int{1},
		},
		{
			"test-3",
			args{&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			}},
			[]int{1, 2},
		},
		{
			"test-4",
			args{&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			}},
			[]int{1, 2},
		},
		{
			"test-5",
			args{structures.GenerateTreeNodesBySlice([]int{
				1, 2, 3, structures.NULL, 5, structures.NULL, 4,
			})},
			[]int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
