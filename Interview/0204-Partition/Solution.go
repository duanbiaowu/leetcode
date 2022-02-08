package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Val < x {
			slow.Val, fast.Val = fast.Val, slow.Val
			slow = slow.Next
		}
		fast = fast.Next
	}
	return head
}
