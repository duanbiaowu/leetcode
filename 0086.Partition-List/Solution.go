package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	smallDummy := &ListNode{}
	largeDummy := &ListNode{}

	small := smallDummy
	large := largeDummy

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	large.Next = nil
	small.Next = largeDummy.Next

	return smallDummy.Next
}
