package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
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
			args{&ListNode{Val: 0}},
			&TreeNode{Val: 0},
		},
		{
			"test-3",
			args{&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			}},
			&TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{
				-10, -3, 0, 5, 9,
			})},
			structures.GenerateTreeNodesBySlice([]int{
				0, -3, 9, -10, structures.NULL, 5,
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
