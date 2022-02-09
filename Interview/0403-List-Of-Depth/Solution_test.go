package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_listOfDepth(t *testing.T) {
	type args struct {
		tree *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{1})},
			[]*ListNode{{Val: 1}},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{1, 2, 3, 4, 5, structures.NULL, 7, 8})},
			[]*ListNode{
				{Val: 1},
				structures.GenerateListNodesByArray([]int{2, 3}),
				structures.GenerateListNodesByArray([]int{4, 5, 7}),
				structures.GenerateListNodesByArray([]int{8}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listOfDepth(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listOfDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
