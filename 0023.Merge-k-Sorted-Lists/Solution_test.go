package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	var tests = []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{[]*ListNode{}},
			nil,
		},
		{
			"test-1.1",
			args{[]*ListNode{
				nil,
			}},
			nil,
		},
		{
			"test-1.2",
			args{[]*ListNode{
				nil,
				nil,
			}},
			nil,
		},
		{
			"test-2",
			args{[]*ListNode{
				{0, nil},
			}},
			&ListNode{0, nil},
		},
		{
			"test-3",
			args{[]*ListNode{
				{0, nil},
				{1, nil},
			}},
			convertArrayToListNodes([]int{0, 1}),
		},
		{
			"test-4",
			args{[]*ListNode{
				{1, nil},
				{0, nil},
			}},
			convertArrayToListNodes([]int{0, 1}),
		},
		{
			"test-5",
			args{[]*ListNode{
				{1, nil},
				{0, nil},
				{2, nil},
			}},
			convertArrayToListNodes([]int{0, 1, 2}),
		},
		{
			"test-6",
			args{[]*ListNode{
				convertArrayToListNodes([]int{1, 3, 5}),
				convertArrayToListNodes([]int{2, 4, 6}),
			}},
			convertArrayToListNodes([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			"test-7",
			args{[]*ListNode{
				convertArrayToListNodes([]int{1, 4, 5}),
				convertArrayToListNodes([]int{1, 3, 4}),
				convertArrayToListNodes([]int{2, 6, 7}),
			}},
			convertArrayToListNodes([]int{1, 1, 2, 3, 4, 4, 5, 6, 7}),
		},
		{
			"test-8",
			args{[]*ListNode{
				convertArrayToListNodes([]int{0, 2, 5}),
			}},
			convertArrayToListNodes([]int{0, 2, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKListsSimple(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func convertArrayToListNodes(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}
