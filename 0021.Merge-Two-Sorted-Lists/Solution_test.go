package leetcode

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test two empty listNodes",
			args{nil, nil},
			nil,
		},

		{
			"No1. test an empty listNodes and a non-empty listNodes",
			args{convertArrayToListNodes([]int{1}), nil},
			convertArrayToListNodes([]int{1}),
		},
		{
			"No2. test an empty listNodes and a non-empty listNodes",
			args{nil, convertArrayToListNodes([]int{2})},
			convertArrayToListNodes([]int{2}),
		},

		{
			"No1. test two non-empty listNodes",
			args{convertArrayToListNodes([]int{1}), convertArrayToListNodes([]int{2})},
			convertArrayToListNodes([]int{1, 2}),
		},
		{
			"No2. test two non-empty listNodes",
			args{convertArrayToListNodes([]int{1, 3, 5}), convertArrayToListNodes([]int{2, 4, 6})},
			convertArrayToListNodes([]int{1, 2, 3, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoListsRecursively(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test two empty listNodes",
			args{nil, nil},
			nil,
		},

		{
			"No1. test an empty listNodes and a non-empty listNodes",
			args{convertArrayToListNodes([]int{1}), nil},
			convertArrayToListNodes([]int{1}),
		},
		{
			"No2. test an empty listNodes and a non-empty listNodes",
			args{nil, convertArrayToListNodes([]int{2})},
			convertArrayToListNodes([]int{2}),
		},

		{
			"No1. test two non-empty listNodes",
			args{convertArrayToListNodes([]int{1}), convertArrayToListNodes([]int{2})},
			convertArrayToListNodes([]int{1, 2}),
		},
		{
			"No2. test two non-empty listNodes",
			args{convertArrayToListNodes([]int{1, 3, 5}), convertArrayToListNodes([]int{2, 4, 6})},
			convertArrayToListNodes([]int{1, 2, 3, 4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsRecursively(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoListsRecursively() = %v, want %v", got, tt.want)
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
