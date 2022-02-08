package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	listCommon := structures.GenerateListNodesByArray([]int{8, 4, 5})
	listA := structures.GenerateListNodesByArray([]int{4, 1})
	listB := structures.GenerateListNodesByArray([]int{5, 0, 1})
	listA.Next.Next = listCommon
	listB.Next.Next.Next = listCommon

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
			args{listA, listB},
			listCommon,
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
