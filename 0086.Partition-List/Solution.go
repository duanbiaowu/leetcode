package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	dummySmall := &ListNode{}
	dummyLarge := &ListNode{}

	small := dummySmall
	large := dummyLarge

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
	small.Next = dummyLarge.Next

	return dummySmall.Next
}
