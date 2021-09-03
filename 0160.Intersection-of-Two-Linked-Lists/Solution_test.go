package leetcode

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	singleListNode := &ListNode{}

	headA := &ListNode{Val: 4}
	pA := headA
	pA.Next = &ListNode{Val: 1}

	headB := &ListNode{Val: 5}
	pB := headB
	pB.Next = &ListNode{Val: 0}
	pB.Next.Next = &ListNode{Val: 1}

	// 相交节点
	headCommon := &ListNode{Val: 8}
	pCommon := headCommon
	pCommon.Next = &ListNode{Val: 4}
	pCommon.Next.Next = &ListNode{Val: 5}

	pA.Next.Next = headCommon
	pB.Next.Next.Next = headCommon

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{&ListNode{}, nil},
			nil,
		},
		{
			"test-3",
			args{nil, &ListNode{}},
			nil,
		},
		{
			"test-4",
			args{singleListNode, singleListNode},
			singleListNode,
		},
		{
			"test-5",
			args{headA, headB},
			headCommon,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
