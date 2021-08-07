package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
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
			nil,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{1})},
			[]int{1},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2})},
			[]int{2, 1},
		},
		{
			"test-4",
			args{structures.GenerateTreeNodesBySlice([]int{1, structures.NULL, 3})},
			[]int{1, 3},
		},
		{
			"test-5",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3})},
			[]int{2, 1, 3},
		},
		{
			"test-6",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5, 6, 7})},
			[]int{4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
