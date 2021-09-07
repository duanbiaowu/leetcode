package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	cycleListNode := &ListNode{Val: 1}
	cycleListNode.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: cycleListNode,
			},
		},
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&ListNode{Val: 1}},
			nil,
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3})},
			nil,
		},
		{
			"test-4",
			args{cycleListNode},
			cycleListNode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
