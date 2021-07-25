package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}

	testListNodes := &ListNode{Val: 1}
	testListNodes.Next = &ListNode{Val: 2}
	testListNodes.Next.Next = &ListNode{Val: 3}
	testListNodes.Next.Next.Next = &ListNode{Val: 4}
	testListNodes.Next.Next.Next.Next = &ListNode{Val: 5}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test an empty linkNodes",
			args{&ListNode{Val: 1}, 1},
			nil,
		},

		{
			"No1. test does not deleting listNode",
			args{&ListNode{Val: 1}, 0},
			&ListNode{Val: 1},
		},
		{
			"No1. test does not deleting listNode",
			args{&ListNode{Val: 1}, -1},
			&ListNode{Val: 1},
		},

		{
			"test deleting first listNode from end",
			args{&ListNode{Val: 1}, 1},
			nil,
		},
		{
			"test deleting second listNode from end",
			args{testListNodes, 2},
			testListNodes,
		},

		{
			"test deleting first listNode",
			args{testListNodes, 5},
			testListNodes.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
